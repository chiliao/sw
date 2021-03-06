package vcli

// process.go: contains command processing code

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/api/login"
	"github.com/pensando/sw/venice/cli/gen"
	"github.com/pensando/sw/venice/utils/runtime"
)

// createCmd creates an object, it describes failure if any occured during post
func createCmd(c *cli.Context) {
	createCmdInternal(c, "create")
}

// updateCmd  updates an object, it describes failure if any occured during put
func updateCmd(c *cli.Context) {
	createCmdInternal(c, "update")
}

// createCmdInternal handles create and update
func createCmdInternal(c *cli.Context, cmd string) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, cmd); err != nil {
		log.Fatalf("error processing global flags: %s", err)
		return
	}

	if ctx.subcmd == "from-file" {
		return
	}

	if c.Bool("file") {
		filename := ""
		if len(c.Args()) > 1 {
			log.Fatalf("One file is permitted to be edited at a time")
		} else if len(c.Args()) == 1 {
			filename = c.Args()[0]
		}
		processFiles(ctx, filename)
		return
	}

	if len(ctx.names) != 1 {
		log.Fatalf("%s - invalid object name(s) %s", ctx.subcmd, ctx.names)
	}

	obj := getNewObj(ctx)
	if cmd == "update" {
		objm, _ := runtime.GetObjectMeta(obj)
		objm.Tenant = ctx.tenant
		objm.Name = ctx.names[0]
		if err := ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, obj); err != nil {
			fmt.Printf("Error getting %s '%s': %v", ctx.subcmd, ctx.names, err)
			return
		}
	}

	specKvs, err := getKvsFromCliCtx(ctx)
	if err != nil {
		fmt.Printf("error parsing arguments: %s, obj %+v", err, obj)
		return
	}

	if err := writeObj(ctx, obj, specKvs); err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	if cmd == "update" {
		if err = ctx.restPutFunc(ctx.server, ctx.token, obj); err != nil {
			fmt.Printf("%s\n", err)
		}
	} else {
		if err = ctx.restPostFunc(ctx.server, ctx.token, obj); err != nil {
			fmt.Printf("%s\n", err)
		}
	}
}

// patchCmd patches updates one or more fields in an object
func patchCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "patch"); err != nil {
		log.Fatalf("error processing global flags: %s", err)
		return
	}

	if len(ctx.names) != 1 {
		log.Fatalf("%s object name not provided", ctx.subcmd)
	}

	obj := getNewObj(ctx)
	objm, _ := runtime.GetObjectMeta(obj)
	objm.Tenant = ctx.tenant
	objm.Name = ctx.names[0]

	if err := ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, obj); err != nil {
		fmt.Printf("Error getting %s '%s': %v", ctx.subcmd, ctx.names, err)
		return
	}

	specKvs, err := getKvsFromCliCtx(ctx)
	if err != nil {
		fmt.Printf("error parsing arguments: %s, obj %+v", err, obj)
		return
	}

	if err := writeObj(ctx, obj, specKvs); err != nil {
		fmt.Printf("%s", err)
		return
	}

	if err = ctx.restPutFunc(ctx.server, ctx.token, obj); err != nil {
		fmt.Printf("%s", err)
	}
}

// readCmd is the function that gets called for any read operation, it gets an object or a list of
// objects; if there are multiple objects, it would show them in a tabulated form
func readCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "read"); err != nil {
		log.Debugf("Context: %+v\n", ctx)
		fmt.Println(err)
		return
	}

	objIf := getNewObjList(ctx)
	if len(ctx.names) == 1 {
		objIf = getNewObj(ctx)
		objm, _ := runtime.GetObjectMeta(objIf)
		objm.Tenant = ctx.tenant
		objm.Name = ctx.names[0]
	}

	err := ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, objIf)
	if err != nil {
		log.Errorf("Error getting %s(s): %s", ctx.subcmd, err)
		return
	}
	log.Debugf("\nGot obj %+v\nContext: %+v\n", objIf, ctx)

	if ctx.dumpStruct {
		dumpStructStdout(ctx.dumpYml, objIf)
		return
	}

	numItems := getNumItems(objIf)
	if ctx.quiet {
		names := []string{}
		for idx := 0; idx < numItems; idx++ {
			obj := getObjOrList(objIf, idx)
			objm, _ := runtime.GetObjectMeta(obj)
			names = append(names, objm.Name)
		}
		os.Stdout.WriteString(strings.Join(names, "\n"))
		return
	}

	header := false
	writer := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
	defer writer.Flush()

	objmKvs, specKvs, objmValidKvs, specValidKvs := getAllKvs(ctx, numItems, objIf)
	for idx := 0; idx < numItems; idx++ {
		obj := getObjOrList(objIf, idx)
		objm, _ := runtime.GetObjectMeta(obj)

		if skipObj(ctx, objm) {
			continue
		}

		if !header {
			oneLine := getLineHeader(ctx, objmValidKvs, specValidKvs)
			writer.Write(oneLine)
			header = true
		}

		recIdx := 0
		for {
			oneLine, more := getLineData(ctx, &recIdx, objmKvs[idx], specKvs[idx], objmValidKvs, specValidKvs)
			writer.Write(oneLine)
			if !more {
				break
			}
		}
	}
}

// deleteCmd is called for all verbs that perform delete on an object; it invokes REST delete APIs in turn
func deleteCmd(c *cli.Context) {
	deleteCmdInternal(c, "delete")
}

func deleteCmdInternal(c *cli.Context, cmd string) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}

	if err := processGlobalFlags(ctx, cmd); err != nil {
		return
	}
	if ctx.subcmd == "from-file" {
		return
	}

	if ctx.re != nil || ctx.labels != nil && len(ctx.labels) > 0 {
		nc := *c
		readCmd(&nc)
		ctx.names = getFilteredNames(ctx)
	}

	for _, name := range ctx.names {
		ctx.names[0] = name
		err := restDelete(ctx)
		if err != nil {
			fmt.Printf("Error deleting %s '%s': %v\n", c.Command.Name, name, err)
			return
		}
	}
}

// clearCmd deletes all objects or objects of a specific type
func clearCmd(c *cli.Context) {
	for _, objKind := range []string{"app", "sgpolicy", "workload", "host", "network"} {
		c.Command.Name = objKind
		deleteCmdInternal(c, "clear")
	}
}

// editCmd allows using an editor to create or modify an object
// if object name is not specified then a sample object is picked up from example template
func editCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "edit"); err != nil {
		return
	}

	filename := ""
	if len(ctx.names) > 1 {
		log.Fatalf("One record update permitted at a time")
	}

	obj := getNewObj(ctx)
	objm, _ := runtime.GetObjectMeta(obj)
	objm.Tenant = ctx.tenant
	objm.Name = ctx.names[0]
	err := ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, obj)
	if err != nil {
		fmt.Printf("Error getting %s '%s': %v", ctx.subcmd, ctx.names, err)
		return
	}
	log.Debugf("Fetched following object %v", obj)

	tf, err := ioutil.TempFile("", ctx.subcmd)
	if err != nil {
		log.Fatalf("Error creating temp file %s '%s': %v", ctx.subcmd, ctx.names, err)
	}
	defer os.Remove(tf.Name())

	if err := removeObjOper(ctx, obj); err != nil {
		log.Fatalf("Error removing oper info from object %+v\n", obj)
	}

	dumpBytes := dumpStruct(ctx.dumpYml, obj)
	tf.Write(dumpBytes)
	// flush the file contents because editFromFile would read from the filename
	if err := tf.Sync(); err != nil {
		fmt.Printf("file sync error - %s", err)
		return
	}
	filename = tf.Name()

	if err := editFromFile(ctx, filename); err != nil {
		fmt.Printf("Edit error: '%s'", err)
	}
}

// updateMetaLabel function is a utility function that adds/deletes a label within object's meta
func updateMetaLabel(objm *api.ObjectMeta, newLabels map[string]string) error {
	for key, value := range newLabels {
		if value == "-" {
			delete(objm.Labels, key)
		} else {
			if objm.Labels == nil {
				objm.Labels = make(map[string]string)
			}
			objm.Labels[key] = value
		}
	}
	return nil
}

// labelcmd performs label update on any object; this is invoked when label verb is invoked for any object
func labelCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "read"); err != nil {
		return
	}

	newLabels, err := sliceToMap(c.StringSlice("update-label"))
	if len(newLabels) == 0 || err != nil {
		if err != nil {
			fmt.Printf("%s", err)
		}
		return
	}

	objList := getNewObjList(ctx)
	err = ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, objList)
	if err != nil {
		fmt.Printf("Error getting %s: %s", ctx.subcmd, err)
		return
	}

	numItems := getNumItems(objList)
	for idx := 0; idx < numItems; idx++ {
		obj := getObjOrList(objList, idx)
		objm, _ := runtime.GetObjectMeta(obj)

		if skipObj(ctx, objm) {
			continue
		}

		updateLabel(obj, newLabels)
		if err := ctx.restPutFunc(ctx.server, ctx.token, obj); err != nil {
			fmt.Printf("%s", err)
		}
	}
}

// treeCmd shows the relationship tree of all (or a kind of) instantiated objects
func treeCmd(c *cli.Context) {
	fmt.Printf("not implemented\n")
}

// exampleCmd is called when example verb is specified for any object
// this command shows an example of a specific operation on an object
func exampleCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "example"); err != nil {
		return
	}

	fmt.Printf("%s\n", walkStruct(ctx.structInfo, 0))
}

// definitionCmd shows object definition
func definitionCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "example"); err != nil {
		return
	}

	fmt.Printf("%s\n", walkStruct(ctx.structInfo, 0))
}

// logoutCmd is called when user executes user logout
// TODO: tie this to RBAC
func logoutCmd(c *cli.Context) {
	clearToken()
	fmt.Println("Logout successful")
}

// loginCmd is called when a user logs in; JWT returned by login is stored at
// $HOME/.pensando/token
func loginCmd(c *cli.Context) {
	clearToken()
	ctx := &cliContext{cli: c, tenant: defaultTenant}
	if err := processGlobalFlags(ctx, "login"); err != nil {
		log.Fatalf("error processing flags: %s", err)
		return
	}

	if len(c.Args()) != 0 {
		log.Fatal("No arguments are allowed for login")
	}

	loginUser := c.String("user")
	password := c.String("password")

	if loginUser == "" || password == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Username: ")
		loginUser, _ = reader.ReadString('\n')
		loginUser = strings.TrimSpace(loginUser)

		fmt.Print("Enter Password: ")
		pBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatalf("Error reading user password: %v", err)
		}
		password = string(pBytes)
	}

	_, token, err := login.UserLogin(context.TODO(), ctx.server, &auth.PasswordCredential{
		Username: loginUser,
		Password: string(password),
		Tenant:   defaultTenant,
	})
	if err != nil {
		log.Fatalf("Failed to login with error: %v", err)
	}
	saveToken(token, ctx.server)
	fmt.Println("\nLogin successful")
}

// snapshotCmd takes the snapshot (inventory) of all objects, removes their
// operational/changeable state, which can then be played back into the system sometime later
func snapshotCmd(c *cli.Context) {
	ctx := &cliContext{cli: c, token: getLoginToken(), server: c.GlobalString("server"), tenant: defaultTenant}

	subCmds := c.Args()
	if len(subCmds) == 0 {
		for _, objKind := range gen.GetInfo().GetAllKeys() {
			subCmds = append(subCmds, objKind)
		}
	}

	dirname := "./"
	if c.String("id") == "" {
		dirname += time.Now().String()
	} else {
		dirname += c.String("id")
	}
	dirname += "/"

	if c.Bool("restore") {
		if err := processGlobalFlags(ctx, "create"); err != nil {
			log.Fatalf("unable to create new cliContext")
		}
		processFiles(ctx, dirname)
		return
	}

	if err := os.Mkdir(dirname, 0777); err != nil {
		log.Fatalf("error creating directory: %s", dirname)
	}

	sort.Strings(subCmds)
	for _, subcmd := range subCmds {
		ctx.subcmd = subcmd
		if err := populateGenCtx(ctx); err != nil {
			fmt.Printf("error populating generated cliContext: %s", err)
			continue
		}

		if ctx.restGetFunc == nil {
			// REST GET function is not available on some objects
			continue
		}

		objList := getNewObjList(ctx)
		err := ctx.restGetFunc(ctx.server, ctx.tenant, ctx.token, objList)
		if err != nil {
			fmt.Printf("Error getting %ss: %v", ctx.subcmd, err)
			continue
		}

		subdir := dirname + "/" + subcmd + "s/"
		if err := os.Mkdir(subdir, 0777); err != nil {
			log.Fatalf("error creating subdir '%s'", subdir)
		}

		numItems := getNumItems(objList)
		for idx := 0; idx < numItems; idx++ {
			obj := getObjOrList(objList, idx)
			objm, _ := runtime.GetObjectMeta(obj)

			if err := removeObjOper(ctx, obj); err != nil {
				log.Fatalf("Error removing oper info from object %+v\n", obj)
			}
			dumpBytes := dumpStruct(ctx.dumpYml, obj)

			filename := subdir + objm.Name
			if ctx.dumpYml {
				filename += ".yml"
			} else {
				filename += ".json"
			}
			tf, err := os.Create(filename)
			if err != nil {
				log.Fatalf("Error creating file %s '%s'", filename, err)
			}
			tf.Write(dumpBytes)
			tf.Close()
		}
	}

	fmt.Printf("Successful - stored snapshot in '%s' directory\n", dirname)
}

// metricsCmd displays metrics corresponding to a measurement (aka table)
// it can be specified with a database name optionally to qualify the table within a tenant
func metricsCmd(c *cli.Context) {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	token := getLoginToken()
	server := c.GlobalString("server")
	tenant := defaultTenant
	maxrows := c.Int("maxrows")
	cols := []string{}
	if c.String("cols") != "" {
		cols = strings.Split(c.String("cols"), ",")
	}

	tableName := "node"
	if len(c.Args()) > 0 {
		tableName = c.Args()[0]
	}
	svcName, _, err := gen.GetInfo().FindSvcName(tableName)
	if err != nil {
		log.Debugf("Error finding svc name - %s", err)
		fmt.Printf("unrecognized table: %s", tableName)
		return
	}
	kind := strings.Split(svcName, ".")[1]

	log.Debugf("server '%s', tenant '%s', cols = %s, maxrows = %v kind = %s \n", server, tenant, cols, maxrows, kind)

	resp, err := metricsQuery(server, kind, token)
	if err != nil {
		fmt.Printf("Internal query error: %s", err)
		return
	}

	for ridx, result := range resp.Results {
		log.Debugf("Result %d", ridx)
		for sidx, series := range result.Series {
			log.Debugf("[%d] columns = %+v value %+v\n", sidx, series.Columns, series.Values)
			printSeries(os.Stdout, maxrows, cols, series)
		}
	}
}

// DefaultVersion is the backend in image version for the CLI frontend
// TODO: obtain the server version from the backend
const DefaultVersion = "v0.1-alpha"

func showVersion(c *cli.Context) {
	if len(c.Args()) != 0 {
		log.Warnf("More arguments than required")
	}
	fmt.Printf("Client Version: %s\n", DefaultVersion)
	fmt.Printf("Server Version: %s\n", DefaultVersion)
}

// configCommandCompletion configures a linux bash command completion
func configCommandCompletion(c *cli.Context) {
	var autoCompletionScript = `#! /bin/bash

: ${PROG:=$(basename ${BASH_SOURCE})}

_cli_bash_autocomplete() {
    local cur opts base
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --gbc )
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
}

complete -F _cli_bash_autocomplete $PROG

unset PROG
`

	fi, err := os.Lstat("/etc/bash_completion.d")
	if err != nil {
		fmt.Printf("Unable to stat /etc/bash_completion.d ... quitting")
		return
	}

	if fi.Mode().IsDir() {
		err = ioutil.WriteFile("/etc/bash_completion.d/venice", []byte(autoCompletionScript), 0755)
	} else {
		err = ioutil.WriteFile("/usr/local/etc/bash_completion.d/venice", []byte(autoCompletionScript), 0755)
	}
	if err != nil {
		fmt.Printf("error creating a file: %s\n", err)
		return
	}
	fmt.Printf("Almost there... please also execute following in your bash: \"source /etc/bash_completion.d/venice\"\n")
}
