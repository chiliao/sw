// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: metrics.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lifmetricsShowCmd = &cobra.Command{

	Use:   "lif",
	Short: "Show LifMetrics from Naples",
	Long:  "\n---------------------------------\n Show LifMetrics From Naples \n---------------------------------\n",
	RunE:  lifmetricsShowCmdHandler,
}

func lifmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/lifmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No lif object(s) found")
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(lifmetricsShowCmd)

}
