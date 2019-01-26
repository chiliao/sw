// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: upgrade.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var upgrademetricsShowCmd = &cobra.Command{

	Use:   "upgrade",
	Short: "Show UpgradeMetrics from Naples",
	Long:  "\n---------------------------------\n Show UpgradeMetrics From Naples \n---------------------------------\n",
	RunE:  upgrademetricsShowCmdHandler,
}

func upgrademetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/upgrademetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No upgrade object(s) found")
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(upgrademetricsShowCmd)

}
