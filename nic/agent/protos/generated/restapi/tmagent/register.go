// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.
/*
 * This file is a auto generated.
 * Input file: *_apigen.go
 */
package restapi

func (s *RestServer) RegisterAPIRoutes() {
	s.PrefixRoutes = map[string]routeAddFunc{
		"/telemetry/v1/metrics/accelhwringmetrics/":                 s.AddAccelHwRingMetricsAPIRoutes,
		"/telemetry/v1/metrics/accelseqqueueinfometrics/":           s.AddAccelSeqQueueInfoMetricsAPIRoutes,
		"/telemetry/v1/metrics/accelseqqueuemetrics/":               s.AddAccelSeqQueueMetricsAPIRoutes,
		"/telemetry/v1/metrics/asicpowermetrics/":                   s.AddAsicPowerMetricsAPIRoutes,
		"/telemetry/v1/metrics/asictemperaturemetrics/":             s.AddAsicTemperatureMetricsAPIRoutes,
		"/telemetry/v1/metrics/dbwaintdbmetrics/":                   s.AddDbwaintdbMetricsAPIRoutes,
		"/telemetry/v1/metrics/dbwaintlifqstatemapmetrics/":         s.AddDbwaintlifqstatemapMetricsAPIRoutes,
		"/telemetry/v1/metrics/dppintcreditmetrics/":                s.AddDppintcreditMetricsAPIRoutes,
		"/telemetry/v1/metrics/dppintfifometrics/":                  s.AddDppintfifoMetricsAPIRoutes,
		"/telemetry/v1/metrics/dppintreg1metrics/":                  s.AddDppintreg1MetricsAPIRoutes,
		"/telemetry/v1/metrics/dppintreg2metrics/":                  s.AddDppintreg2MetricsAPIRoutes,
		"/telemetry/v1/metrics/dppintsramseccmetrics/":              s.AddDppintsramseccMetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintcreditmetrics/":                s.AddDprintcreditMetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintfifometrics/":                  s.AddDprintfifoMetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintflopfifometrics/":              s.AddDprintflopfifoMetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintreg1metrics/":                  s.AddDprintreg1MetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintreg2metrics/":                  s.AddDprintreg2MetricsAPIRoutes,
		"/telemetry/v1/metrics/dprintsramseccmetrics/":              s.AddDprintsramseccMetricsAPIRoutes,
		"/telemetry/v1/metrics/ftecpsmetrics/":                      s.AddFteCPSMetricsAPIRoutes,
		"/telemetry/v1/metrics/ftelifqmetrics/":                     s.AddFteLifQMetricsAPIRoutes,
		"/telemetry/v1/metrics/inteccdescmetrics/":                  s.AddInteccdescMetricsAPIRoutes,
		"/telemetry/v1/metrics/intsparemetrics/":                    s.AddIntspareMetricsAPIRoutes,
		"/telemetry/v1/metrics/lifmetrics/":                         s.AddLifMetricsAPIRoutes,
		"/telemetry/v1/metrics/mcmchintmcmetrics/":                  s.AddMcmchintmcMetricsAPIRoutes,
		"/telemetry/v1/metrics/mdhensintaxierrmetrics/":             s.AddMdhensintaxierrMetricsAPIRoutes,
		"/telemetry/v1/metrics/mdhensinteccmetrics/":                s.AddMdhensinteccMetricsAPIRoutes,
		"/telemetry/v1/metrics/mdhensintipcoremetrics/":             s.AddMdhensintipcoreMetricsAPIRoutes,
		"/telemetry/v1/metrics/mpmpnsintcryptometrics/":             s.AddMpmpnsintcryptoMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbchbmintecchbmrbmetrics/":         s.AddPbpbchbmintecchbmrbMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbchbminthbmaxierrrspmetrics/":     s.AddPbpbchbminthbmaxierrrspMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbchbminthbmdropmetrics/":          s.AddPbpbchbminthbmdropMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbchbminthbmpbusviolationmetrics/": s.AddPbpbchbminthbmpbusviolationMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbchbminthbmxoffmetrics/":          s.AddPbpbchbminthbmxoffMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbcintcreditunderflowmetrics/":     s.AddPbpbcintcreditunderflowMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbcintpbusviolationmetrics/":       s.AddPbpbcintpbusviolationMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbcintrplmetrics/":                 s.AddPbpbcintrplMetricsAPIRoutes,
		"/telemetry/v1/metrics/pbpbcintwritemetrics/":               s.AddPbpbcintwriteMetricsAPIRoutes,
		"/telemetry/v1/metrics/pciemgrmetrics/":                     s.AddPcieMgrMetricsAPIRoutes,
		"/telemetry/v1/metrics/pcieportmetrics/":                    s.AddPciePortMetricsAPIRoutes,
		"/telemetry/v1/metrics/sessionsummarymetrics/":              s.AddSessionSummaryMetricsAPIRoutes,
		"/telemetry/v1/metrics/sgempuinterrmetrics/":                s.AddSgempuinterrMetricsAPIRoutes,
		"/telemetry/v1/metrics/sgempuintinfometrics/":               s.AddSgempuintinfoMetricsAPIRoutes,
		"/telemetry/v1/metrics/sgeteinterrmetrics/":                 s.AddSgeteinterrMetricsAPIRoutes,
		"/telemetry/v1/metrics/sgeteintinfometrics/":                s.AddSgeteintinfoMetricsAPIRoutes,
		"/telemetry/v1/metrics/ssepicsintbadaddrmetrics/":           s.AddSsepicsintbadaddrMetricsAPIRoutes,
		"/telemetry/v1/metrics/ssepicsintbgmetrics/":                s.AddSsepicsintbgMetricsAPIRoutes,
		"/telemetry/v1/metrics/ssepicsintpicsmetrics/":              s.AddSsepicsintpicsMetricsAPIRoutes,
		"/telemetry/v1/metrics/upgrademetrics/":                     s.AddUpgradeMetricsAPIRoutes,
	}
}
func (s *RestServer) RegisterListMetrics() {
	s.GetPointsFuncList = map[string]getPointsFunc{
		"AccelHwRingMetrics":                 s.getAccelHwRingMetricsPoints,
		"AccelSeqQueueInfoMetrics":           s.getAccelSeqQueueInfoMetricsPoints,
		"AccelSeqQueueMetrics":               s.getAccelSeqQueueMetricsPoints,
		"AsicPowerMetrics":                   s.getAsicPowerMetricsPoints,
		"AsicTemperatureMetrics":             s.getAsicTemperatureMetricsPoints,
		"DbwaintdbMetrics":                   s.getDbwaintdbMetricsPoints,
		"DbwaintlifqstatemapMetrics":         s.getDbwaintlifqstatemapMetricsPoints,
		"DppintcreditMetrics":                s.getDppintcreditMetricsPoints,
		"DppintfifoMetrics":                  s.getDppintfifoMetricsPoints,
		"Dppintreg1Metrics":                  s.getDppintreg1MetricsPoints,
		"Dppintreg2Metrics":                  s.getDppintreg2MetricsPoints,
		"DppintsramseccMetrics":              s.getDppintsramseccMetricsPoints,
		"DprintcreditMetrics":                s.getDprintcreditMetricsPoints,
		"DprintfifoMetrics":                  s.getDprintfifoMetricsPoints,
		"DprintflopfifoMetrics":              s.getDprintflopfifoMetricsPoints,
		"Dprintreg1Metrics":                  s.getDprintreg1MetricsPoints,
		"Dprintreg2Metrics":                  s.getDprintreg2MetricsPoints,
		"DprintsramseccMetrics":              s.getDprintsramseccMetricsPoints,
		"FteCPSMetrics":                      s.getFteCPSMetricsPoints,
		"FteLifQMetrics":                     s.getFteLifQMetricsPoints,
		"InteccdescMetrics":                  s.getInteccdescMetricsPoints,
		"IntspareMetrics":                    s.getIntspareMetricsPoints,
		"LifMetrics":                         s.getLifMetricsPoints,
		"McmchintmcMetrics":                  s.getMcmchintmcMetricsPoints,
		"MdhensintaxierrMetrics":             s.getMdhensintaxierrMetricsPoints,
		"MdhensinteccMetrics":                s.getMdhensinteccMetricsPoints,
		"MdhensintipcoreMetrics":             s.getMdhensintipcoreMetricsPoints,
		"MpmpnsintcryptoMetrics":             s.getMpmpnsintcryptoMetricsPoints,
		"PbpbchbmintecchbmrbMetrics":         s.getPbpbchbmintecchbmrbMetricsPoints,
		"PbpbchbminthbmaxierrrspMetrics":     s.getPbpbchbminthbmaxierrrspMetricsPoints,
		"PbpbchbminthbmdropMetrics":          s.getPbpbchbminthbmdropMetricsPoints,
		"PbpbchbminthbmpbusviolationMetrics": s.getPbpbchbminthbmpbusviolationMetricsPoints,
		"PbpbchbminthbmxoffMetrics":          s.getPbpbchbminthbmxoffMetricsPoints,
		"PbpbcintcreditunderflowMetrics":     s.getPbpbcintcreditunderflowMetricsPoints,
		"PbpbcintpbusviolationMetrics":       s.getPbpbcintpbusviolationMetricsPoints,
		"PbpbcintrplMetrics":                 s.getPbpbcintrplMetricsPoints,
		"PbpbcintwriteMetrics":               s.getPbpbcintwriteMetricsPoints,
		"PcieMgrMetrics":                     s.getPcieMgrMetricsPoints,
		"PciePortMetrics":                    s.getPciePortMetricsPoints,
		"SessionSummaryMetrics":              s.getSessionSummaryMetricsPoints,
		"SgempuinterrMetrics":                s.getSgempuinterrMetricsPoints,
		"SgempuintinfoMetrics":               s.getSgempuintinfoMetricsPoints,
		"SgeteinterrMetrics":                 s.getSgeteinterrMetricsPoints,
		"SgeteintinfoMetrics":                s.getSgeteintinfoMetricsPoints,
		"SsepicsintbadaddrMetrics":           s.getSsepicsintbadaddrMetricsPoints,
		"SsepicsintbgMetrics":                s.getSsepicsintbgMetricsPoints,
		"SsepicsintpicsMetrics":              s.getSsepicsintpicsMetricsPoints,
		"UpgradeMetrics":                     s.getUpgradeMetricsPoints,
	}
}
