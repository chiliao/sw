// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * Package cmd is a auto generated package.
 * Input file: asicerrord.proto
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dbwaintdbmetricsShowCmd = &cobra.Command{

	Use:   "dbwaintdb",
	Short: "Show DbwaintdbMetrics from Naples",
	Long:  "\n---------------------------------\n Show DbwaintdbMetrics From Naples \n---------------------------------\n",
	RunE:  dbwaintdbmetricsShowCmdHandler,
}

func dbwaintdbmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dbwaintdbmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dbwaintdb object(s) found")
	}
	return nil
}

var dbwaintlifqstatemapmetricsShowCmd = &cobra.Command{

	Use:   "dbwaintlifqstatemap",
	Short: "Show DbwaintlifqstatemapMetrics from Naples",
	Long:  "\n---------------------------------\n Show DbwaintlifqstatemapMetrics From Naples \n---------------------------------\n",
	RunE:  dbwaintlifqstatemapmetricsShowCmdHandler,
}

func dbwaintlifqstatemapmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dbwaintlifqstatemapmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dbwaintlifqstatemap object(s) found")
	}
	return nil
}

var dppintcreditmetricsShowCmd = &cobra.Command{

	Use:   "dppintcredit",
	Short: "Show DppintcreditMetrics from Naples",
	Long:  "\n---------------------------------\n Show DppintcreditMetrics From Naples \n---------------------------------\n",
	RunE:  dppintcreditmetricsShowCmdHandler,
}

func dppintcreditmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dppintcreditmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dppintcredit object(s) found")
	}
	return nil
}

var dppintfifometricsShowCmd = &cobra.Command{

	Use:   "dppintfifo",
	Short: "Show DppintfifoMetrics from Naples",
	Long:  "\n---------------------------------\n Show DppintfifoMetrics From Naples \n---------------------------------\n",
	RunE:  dppintfifometricsShowCmdHandler,
}

func dppintfifometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dppintfifometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dppintfifo object(s) found")
	}
	return nil
}

var dppintreg1metricsShowCmd = &cobra.Command{

	Use:   "dppintreg1",
	Short: "Show Dppintreg1Metrics from Naples",
	Long:  "\n---------------------------------\n Show Dppintreg1Metrics From Naples \n---------------------------------\n",
	RunE:  dppintreg1metricsShowCmdHandler,
}

func dppintreg1metricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dppintreg1metrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dppintreg1 object(s) found")
	}
	return nil
}

var dppintreg2metricsShowCmd = &cobra.Command{

	Use:   "dppintreg2",
	Short: "Show Dppintreg2Metrics from Naples",
	Long:  "\n---------------------------------\n Show Dppintreg2Metrics From Naples \n---------------------------------\n",
	RunE:  dppintreg2metricsShowCmdHandler,
}

func dppintreg2metricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dppintreg2metrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dppintreg2 object(s) found")
	}
	return nil
}

var dppintsramseccmetricsShowCmd = &cobra.Command{

	Use:   "dppintsramsecc",
	Short: "Show DppintsramseccMetrics from Naples",
	Long:  "\n---------------------------------\n Show DppintsramseccMetrics From Naples \n---------------------------------\n",
	RunE:  dppintsramseccmetricsShowCmdHandler,
}

func dppintsramseccmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dppintsramseccmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dppintsramsecc object(s) found")
	}
	return nil
}

var dprintcreditmetricsShowCmd = &cobra.Command{

	Use:   "dprintcredit",
	Short: "Show DprintcreditMetrics from Naples",
	Long:  "\n---------------------------------\n Show DprintcreditMetrics From Naples \n---------------------------------\n",
	RunE:  dprintcreditmetricsShowCmdHandler,
}

func dprintcreditmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintcreditmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintcredit object(s) found")
	}
	return nil
}

var dprintfifometricsShowCmd = &cobra.Command{

	Use:   "dprintfifo",
	Short: "Show DprintfifoMetrics from Naples",
	Long:  "\n---------------------------------\n Show DprintfifoMetrics From Naples \n---------------------------------\n",
	RunE:  dprintfifometricsShowCmdHandler,
}

func dprintfifometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintfifometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintfifo object(s) found")
	}
	return nil
}

var dprintflopfifometricsShowCmd = &cobra.Command{

	Use:   "dprintflopfifo",
	Short: "Show DprintflopfifoMetrics from Naples",
	Long:  "\n---------------------------------\n Show DprintflopfifoMetrics From Naples \n---------------------------------\n",
	RunE:  dprintflopfifometricsShowCmdHandler,
}

func dprintflopfifometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintflopfifometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintflopfifo object(s) found")
	}
	return nil
}

var dprintreg1metricsShowCmd = &cobra.Command{

	Use:   "dprintreg1",
	Short: "Show Dprintreg1Metrics from Naples",
	Long:  "\n---------------------------------\n Show Dprintreg1Metrics From Naples \n---------------------------------\n",
	RunE:  dprintreg1metricsShowCmdHandler,
}

func dprintreg1metricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintreg1metrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintreg1 object(s) found")
	}
	return nil
}

var dprintreg2metricsShowCmd = &cobra.Command{

	Use:   "dprintreg2",
	Short: "Show Dprintreg2Metrics from Naples",
	Long:  "\n---------------------------------\n Show Dprintreg2Metrics From Naples \n---------------------------------\n",
	RunE:  dprintreg2metricsShowCmdHandler,
}

func dprintreg2metricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintreg2metrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintreg2 object(s) found")
	}
	return nil
}

var dprintsramseccmetricsShowCmd = &cobra.Command{

	Use:   "dprintsramsecc",
	Short: "Show DprintsramseccMetrics from Naples",
	Long:  "\n---------------------------------\n Show DprintsramseccMetrics From Naples \n---------------------------------\n",
	RunE:  dprintsramseccmetricsShowCmdHandler,
}

func dprintsramseccmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/dprintsramseccmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No dprintsramsecc object(s) found")
	}
	return nil
}

var inteccdescmetricsShowCmd = &cobra.Command{

	Use:   "inteccdesc",
	Short: "Show InteccdescMetrics from Naples",
	Long:  "\n---------------------------------\n Show InteccdescMetrics From Naples \n---------------------------------\n",
	RunE:  inteccdescmetricsShowCmdHandler,
}

func inteccdescmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/inteccdescmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No inteccdesc object(s) found")
	}
	return nil
}

var intsparemetricsShowCmd = &cobra.Command{

	Use:   "intspare",
	Short: "Show IntspareMetrics from Naples",
	Long:  "\n---------------------------------\n Show IntspareMetrics From Naples \n---------------------------------\n",
	RunE:  intsparemetricsShowCmdHandler,
}

func intsparemetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/intsparemetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No intspare object(s) found")
	}
	return nil
}

var mcmchintmcmetricsShowCmd = &cobra.Command{

	Use:   "mcmchintmc",
	Short: "Show McmchintmcMetrics from Naples",
	Long:  "\n---------------------------------\n Show McmchintmcMetrics From Naples \n---------------------------------\n",
	RunE:  mcmchintmcmetricsShowCmdHandler,
}

func mcmchintmcmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/mcmchintmcmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No mcmchintmc object(s) found")
	}
	return nil
}

var mdhensintaxierrmetricsShowCmd = &cobra.Command{

	Use:   "mdhensintaxierr",
	Short: "Show MdhensintaxierrMetrics from Naples",
	Long:  "\n---------------------------------\n Show MdhensintaxierrMetrics From Naples \n---------------------------------\n",
	RunE:  mdhensintaxierrmetricsShowCmdHandler,
}

func mdhensintaxierrmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/mdhensintaxierrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No mdhensintaxierr object(s) found")
	}
	return nil
}

var mdhensinteccmetricsShowCmd = &cobra.Command{

	Use:   "mdhensintecc",
	Short: "Show MdhensinteccMetrics from Naples",
	Long:  "\n---------------------------------\n Show MdhensinteccMetrics From Naples \n---------------------------------\n",
	RunE:  mdhensinteccmetricsShowCmdHandler,
}

func mdhensinteccmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/mdhensinteccmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No mdhensintecc object(s) found")
	}
	return nil
}

var mdhensintipcoremetricsShowCmd = &cobra.Command{

	Use:   "mdhensintipcore",
	Short: "Show MdhensintipcoreMetrics from Naples",
	Long:  "\n---------------------------------\n Show MdhensintipcoreMetrics From Naples \n---------------------------------\n",
	RunE:  mdhensintipcoremetricsShowCmdHandler,
}

func mdhensintipcoremetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/mdhensintipcoremetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No mdhensintipcore object(s) found")
	}
	return nil
}

var mpmpnsintcryptometricsShowCmd = &cobra.Command{

	Use:   "mpmpnsintcrypto",
	Short: "Show MpmpnsintcryptoMetrics from Naples",
	Long:  "\n---------------------------------\n Show MpmpnsintcryptoMetrics From Naples \n---------------------------------\n",
	RunE:  mpmpnsintcryptometricsShowCmdHandler,
}

func mpmpnsintcryptometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/mpmpnsintcryptometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No mpmpnsintcrypto object(s) found")
	}
	return nil
}

var pbpbchbmintecchbmrbmetricsShowCmd = &cobra.Command{

	Use:   "pbpbchbmintecchbmrb",
	Short: "Show PbpbchbmintecchbmrbMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbchbmintecchbmrbMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbchbmintecchbmrbmetricsShowCmdHandler,
}

func pbpbchbmintecchbmrbmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbchbmintecchbmrbmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbchbmintecchbmrb object(s) found")
	}
	return nil
}

var pbpbchbminthbmaxierrrspmetricsShowCmd = &cobra.Command{

	Use:   "pbpbchbminthbmaxierrrsp",
	Short: "Show PbpbchbminthbmaxierrrspMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbchbminthbmaxierrrspMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbchbminthbmaxierrrspmetricsShowCmdHandler,
}

func pbpbchbminthbmaxierrrspmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbchbminthbmaxierrrspmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbchbminthbmaxierrrsp object(s) found")
	}
	return nil
}

var pbpbchbminthbmdropmetricsShowCmd = &cobra.Command{

	Use:   "pbpbchbminthbmdrop",
	Short: "Show PbpbchbminthbmdropMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbchbminthbmdropMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbchbminthbmdropmetricsShowCmdHandler,
}

func pbpbchbminthbmdropmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbchbminthbmdropmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbchbminthbmdrop object(s) found")
	}
	return nil
}

var pbpbchbminthbmpbusviolationmetricsShowCmd = &cobra.Command{

	Use:   "pbpbchbminthbmpbusviolation",
	Short: "Show PbpbchbminthbmpbusviolationMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbchbminthbmpbusviolationMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbchbminthbmpbusviolationmetricsShowCmdHandler,
}

func pbpbchbminthbmpbusviolationmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbchbminthbmpbusviolationmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbchbminthbmpbusviolation object(s) found")
	}
	return nil
}

var pbpbchbminthbmxoffmetricsShowCmd = &cobra.Command{

	Use:   "pbpbchbminthbmxoff",
	Short: "Show PbpbchbminthbmxoffMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbchbminthbmxoffMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbchbminthbmxoffmetricsShowCmdHandler,
}

func pbpbchbminthbmxoffmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbchbminthbmxoffmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbchbminthbmxoff object(s) found")
	}
	return nil
}

var pbpbcintcreditunderflowmetricsShowCmd = &cobra.Command{

	Use:   "pbpbcintcreditunderflow",
	Short: "Show PbpbcintcreditunderflowMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbcintcreditunderflowMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbcintcreditunderflowmetricsShowCmdHandler,
}

func pbpbcintcreditunderflowmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbcintcreditunderflowmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbcintcreditunderflow object(s) found")
	}
	return nil
}

var pbpbcintpbusviolationmetricsShowCmd = &cobra.Command{

	Use:   "pbpbcintpbusviolation",
	Short: "Show PbpbcintpbusviolationMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbcintpbusviolationMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbcintpbusviolationmetricsShowCmdHandler,
}

func pbpbcintpbusviolationmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbcintpbusviolationmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbcintpbusviolation object(s) found")
	}
	return nil
}

var pbpbcintrplmetricsShowCmd = &cobra.Command{

	Use:   "pbpbcintrpl",
	Short: "Show PbpbcintrplMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbcintrplMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbcintrplmetricsShowCmdHandler,
}

func pbpbcintrplmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbcintrplmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbcintrpl object(s) found")
	}
	return nil
}

var pbpbcintwritemetricsShowCmd = &cobra.Command{

	Use:   "pbpbcintwrite",
	Short: "Show PbpbcintwriteMetrics from Naples",
	Long:  "\n---------------------------------\n Show PbpbcintwriteMetrics From Naples \n---------------------------------\n",
	RunE:  pbpbcintwritemetricsShowCmdHandler,
}

func pbpbcintwritemetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/pbpbcintwritemetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No pbpbcintwrite object(s) found")
	}
	return nil
}

var sgempuinterrmetricsShowCmd = &cobra.Command{

	Use:   "sgempuinterr",
	Short: "Show SgempuinterrMetrics from Naples",
	Long:  "\n---------------------------------\n Show SgempuinterrMetrics From Naples \n---------------------------------\n",
	RunE:  sgempuinterrmetricsShowCmdHandler,
}

func sgempuinterrmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/sgempuinterrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No sgempuinterr object(s) found")
	}
	return nil
}

var sgempuintinfometricsShowCmd = &cobra.Command{

	Use:   "sgempuintinfo",
	Short: "Show SgempuintinfoMetrics from Naples",
	Long:  "\n---------------------------------\n Show SgempuintinfoMetrics From Naples \n---------------------------------\n",
	RunE:  sgempuintinfometricsShowCmdHandler,
}

func sgempuintinfometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/sgempuintinfometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No sgempuintinfo object(s) found")
	}
	return nil
}

var sgeteinterrmetricsShowCmd = &cobra.Command{

	Use:   "sgeteinterr",
	Short: "Show SgeteinterrMetrics from Naples",
	Long:  "\n---------------------------------\n Show SgeteinterrMetrics From Naples \n---------------------------------\n",
	RunE:  sgeteinterrmetricsShowCmdHandler,
}

func sgeteinterrmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/sgeteinterrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No sgeteinterr object(s) found")
	}
	return nil
}

var sgeteintinfometricsShowCmd = &cobra.Command{

	Use:   "sgeteintinfo",
	Short: "Show SgeteintinfoMetrics from Naples",
	Long:  "\n---------------------------------\n Show SgeteintinfoMetrics From Naples \n---------------------------------\n",
	RunE:  sgeteintinfometricsShowCmdHandler,
}

func sgeteintinfometricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/sgeteintinfometrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No sgeteintinfo object(s) found")
	}
	return nil
}

var ssepicsintbadaddrmetricsShowCmd = &cobra.Command{

	Use:   "ssepicsintbadaddr",
	Short: "Show SsepicsintbadaddrMetrics from Naples",
	Long:  "\n---------------------------------\n Show SsepicsintbadaddrMetrics From Naples \n---------------------------------\n",
	RunE:  ssepicsintbadaddrmetricsShowCmdHandler,
}

func ssepicsintbadaddrmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/ssepicsintbadaddrmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No ssepicsintbadaddr object(s) found")
	}
	return nil
}

var ssepicsintbgmetricsShowCmd = &cobra.Command{

	Use:   "ssepicsintbg",
	Short: "Show SsepicsintbgMetrics from Naples",
	Long:  "\n---------------------------------\n Show SsepicsintbgMetrics From Naples \n---------------------------------\n",
	RunE:  ssepicsintbgmetricsShowCmdHandler,
}

func ssepicsintbgmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/ssepicsintbgmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No ssepicsintbg object(s) found")
	}
	return nil
}

var ssepicsintpicsmetricsShowCmd = &cobra.Command{

	Use:   "ssepicsintpics",
	Short: "Show SsepicsintpicsMetrics from Naples",
	Long:  "\n---------------------------------\n Show SsepicsintpicsMetrics From Naples \n---------------------------------\n",
	RunE:  ssepicsintpicsmetricsShowCmdHandler,
}

func ssepicsintpicsmetricsShowCmdHandler(cmd *cobra.Command, args []string) error {
	tabularFormat = false
	jsonFormat = true
	bodyBytes, err := restGet("telemetry/v1/metrics/ssepicsintpicsmetrics/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	if bodyBytes == nil {
		fmt.Println("No ssepicsintpics object(s) found")
	}
	return nil
}

func init() {

	metricsShowCmd.AddCommand(dbwaintdbmetricsShowCmd)

	metricsShowCmd.AddCommand(dbwaintlifqstatemapmetricsShowCmd)

	metricsShowCmd.AddCommand(dppintcreditmetricsShowCmd)

	metricsShowCmd.AddCommand(dppintfifometricsShowCmd)

	metricsShowCmd.AddCommand(dppintreg1metricsShowCmd)

	metricsShowCmd.AddCommand(dppintreg2metricsShowCmd)

	metricsShowCmd.AddCommand(dppintsramseccmetricsShowCmd)

	metricsShowCmd.AddCommand(dprintcreditmetricsShowCmd)

	metricsShowCmd.AddCommand(dprintfifometricsShowCmd)

	metricsShowCmd.AddCommand(dprintflopfifometricsShowCmd)

	metricsShowCmd.AddCommand(dprintreg1metricsShowCmd)

	metricsShowCmd.AddCommand(dprintreg2metricsShowCmd)

	metricsShowCmd.AddCommand(dprintsramseccmetricsShowCmd)

	metricsShowCmd.AddCommand(inteccdescmetricsShowCmd)

	metricsShowCmd.AddCommand(intsparemetricsShowCmd)

	metricsShowCmd.AddCommand(mcmchintmcmetricsShowCmd)

	metricsShowCmd.AddCommand(mdhensintaxierrmetricsShowCmd)

	metricsShowCmd.AddCommand(mdhensinteccmetricsShowCmd)

	metricsShowCmd.AddCommand(mdhensintipcoremetricsShowCmd)

	metricsShowCmd.AddCommand(mpmpnsintcryptometricsShowCmd)

	metricsShowCmd.AddCommand(pbpbchbmintecchbmrbmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbchbminthbmaxierrrspmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbchbminthbmdropmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbchbminthbmpbusviolationmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbchbminthbmxoffmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbcintcreditunderflowmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbcintpbusviolationmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbcintrplmetricsShowCmd)

	metricsShowCmd.AddCommand(pbpbcintwritemetricsShowCmd)

	metricsShowCmd.AddCommand(sgempuinterrmetricsShowCmd)

	metricsShowCmd.AddCommand(sgempuintinfometricsShowCmd)

	metricsShowCmd.AddCommand(sgeteinterrmetricsShowCmd)

	metricsShowCmd.AddCommand(sgeteintinfometricsShowCmd)

	metricsShowCmd.AddCommand(ssepicsintbadaddrmetricsShowCmd)

	metricsShowCmd.AddCommand(ssepicsintbgmetricsShowCmd)

	metricsShowCmd.AddCommand(ssepicsintpicsmetricsShowCmd)

}
