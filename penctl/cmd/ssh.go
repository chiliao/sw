//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

package cmd

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	nmd "github.com/pensando/sw/nic/agent/protos/nmd"
)

var setSSHConfigCmd = &cobra.Command{
	Use:   "ssh-pub-key",
	Short: "Configure ssh pub-key on Naples",
	Long:  "\n---------------------------------\n Configure ssh pub-key on Naples \n---------------------------------\n",
	RunE:  setSSHConfigCmdHandler,
}

var delSSHConfigCmd = &cobra.Command{
	Use:   "ssh-pub-key",
	Short: "Delete ssh pub-key on Naples",
	Long:  "\n------------------------------\n Delete ssh pub-key on Naples \n------------------------------\n",
	RunE:  delSSHConfigCmdHandler,
}

var enableSSHConfigCmd = &cobra.Command{
	Use:   "enable-sshd",
	Short: "Enable sshd on Naples",
	Long:  "\n------------------------------\n Enable sshd on Naples \n------------------------------\n",
	RunE:  enableSSHConfigCmdHandler,
}

var disableSSHConfigCmd = &cobra.Command{
	Use:   "disable-sshd",
	Short: "Disable sshd on Naples",
	Long:  "\n------------------------------\n Disable sshd on Naples \n------------------------------\n",
	RunE:  disableSSHConfigCmdHandler,
}

var uploadPubKeyFile string

func init() {
	updateCmd.AddCommand(setSSHConfigCmd)
	deleteCmd.AddCommand(delSSHConfigCmd)
	sysCmd.AddCommand(enableSSHConfigCmd)
	sysCmd.AddCommand(disableSSHConfigCmd)

	setSSHConfigCmd.Flags().StringVarP(&uploadPubKeyFile, "file", "f", "", "Public Key file location/name")
	setSSHConfigCmd.MarkFlagRequired("file")
}

func delSSHConfigCmdHandler(cmd *cobra.Command, args []string) error {
	v := &nmd.NaplesCmdExecute{
		Executable: "/bin/rm",
		Opts:       strings.Join([]string{"/root/.ssh/authorized_keys"}, ""),
	}
	return naplesExecCmd(v)
}

func setSSHConfigCmdHandler(cmd *cobra.Command, args []string) error {
	v := &nmd.NaplesCmdExecute{
		Executable: "/bin/mkdir",
		Opts:       strings.Join([]string{"/root/.ssh/"}, ""),
	}
	naplesExecCmd(v)

	v = &nmd.NaplesCmdExecute{
		Executable: "/bin/touch",
		Opts:       strings.Join([]string{"/root/.ssh/authorized_keys"}, ""),
	}
	if err := naplesExecCmd(v); err != nil {
		return err
	}

	if errF := canOpen(uploadPubKeyFile); errF != nil {
		return errF
	}
	//prepare the reader instances to encode
	values := map[string]io.Reader{
		"uploadFile": mustOpen(uploadPubKeyFile),
		"uploadPath": strings.NewReader("/update/"),
	}
	_, err := restPostForm("update/", values)
	if err != nil {
		fmt.Println(err)
		return err
	}

	pubKeyFile := filepath.Base(uploadPubKeyFile)

	v = &nmd.NaplesCmdExecute{
		Executable: "setsshauthkey",
		Opts:       strings.Join([]string{"/update/" + pubKeyFile}, ""),
	}

	if err = naplesExecCmd(v); err != nil {
		fmt.Println(err)
		v = &nmd.NaplesCmdExecute{
			Executable: "rm",
			Opts:       strings.Join([]string{"-rf ", "/update/" + pubKeyFile}, ""),
		}
		if err := naplesExecCmd(v); err != nil {
			return err
		}
		return errors.New("Unable to install pubKeyFile " + pubKeyFile)
	}

	v = &nmd.NaplesCmdExecute{
		Executable: "rm",
		Opts:       strings.Join([]string{"-rf ", "/update/" + pubKeyFile}, ""),
	}
	return naplesExecCmd(v)
}

func disableSSHConfigCmdHandler(cmd *cobra.Command, args []string) error {
	v := &nmd.NaplesCmdExecute{
		Executable: "killall",
		Opts:       strings.Join([]string{"sshd"}, ""),
	}
	if err := naplesExecCmd(v); err != nil {
		fmt.Println(err)
		return errors.New("Unable to killall sshd")
	}
	v = &nmd.NaplesCmdExecute{
		Executable: "rm",
		Opts:       strings.Join([]string{"-f ", "/var/lock/sshd /root/.ssh/authorized_keys"}, ""),
	}
	return naplesExecCmd(v)
}

func enableSSHConfigCmdHandler(cmd *cobra.Command, args []string) error {
	v := &nmd.NaplesCmdExecute{
		Executable: "/usr/bin/ssh-keygen",
		Opts:       strings.Join([]string{"-A"}, ""),
	}
	if err := naplesExecCmd(v); err != nil {
		fmt.Println(err)
		return errors.New("Unable to create missing keys")
	}
	v = &nmd.NaplesCmdExecute{
		Executable: "/usr/sbin/sshd",
		Opts:       strings.Join([]string{""}, ""),
	}
	if err := naplesExecCmd(v); err != nil {
		fmt.Println(err)
		return errors.New("Unable to start sshd")
	}
	v = &nmd.NaplesCmdExecute{
		Executable: "touch",
		Opts:       strings.Join([]string{"/var/lock/sshd"}, ""),
	}
	return naplesExecCmd(v)
}
