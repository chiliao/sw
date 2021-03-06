"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var utils_1 = require("./utils");
var fs = require("fs");
var utils = require("./utils");
var custom_metrics_1 = require("../custom-metrics");
var _ = require("lodash");
var rmdir = require("rimraf");
var enums_1 = require("../v1/models/generated/network/enums");
function generatePipelineMetricMetadata(pipeline, inputBaseFolder, outputFolder) {
    var basetypeToJSType = {
        int8: 'number',
        uint8: 'number',
        int16: 'number',
        uint16: 'number',
        int32: 'number',
        uint32: 'number',
        int64: 'number',
        uint64: 'number',
        int: 'number',
        uint: 'number',
        Counter: 'number',
        float32: 'number',
        float64: 'number',
        complex64: 'number',
        complex128: 'number',
        byte: 'number',
        number: 'number',
        Gauge: 'number',
        PrecisionGauge: 'number',
        Bitmap: 'number',
        Celsius: 'number',
        KB: 'number',
        MHz: 'number',
        MilliWatts: 'number',
        string: 'string',
    };
    var scopeObjMapping = {
        PerLIF: "NetworkInterface",
        PerMgmtPort: "NetworkInterface",
        PerEthPort: "NetworkInterface",
    };
    var networkInfMapping = {
        PerLIF: enums_1.NetworkNetworkInterfaceSpec_type["host-pf"],
        PerMgmtPort: enums_1.NetworkNetworkInterfaceSpec_type["uplink-mgmt"],
        PerEthPort: enums_1.NetworkNetworkInterfaceSpec_type["uplink-eth"],
    };
    var files = utils_1.walk(inputBaseFolder, function (x) { return x.endsWith('.json'); });
    var messages = [];
    files.forEach(function (f) {
        var metadata = JSON.parse(fs.readFileSync(f, 'utf8').trim());
        metadata.Messages.forEach(function (m) {
            m = _.transform(m, function (result, val, key) {
                result[_.camelCase(key)] = val;
            });
            // Object kind
            if (scopeObjMapping[m.scope] != null) {
                m.objectKind = scopeObjMapping[m.scope];
            }
            else {
                m.objectKind = "DistributedServiceCard";
            }
            if (m.objectKind == "NetworkInterface") {
                m.interfaceType = networkInfMapping[m.scope];
            }
            if (m.fields == null) {
                return;
            }
            m.fields.push({
                Name: 'reporterID',
                Description: 'Name of reporting object',
                BaseType: 'string',
                JsType: 'string',
                IsTag: true,
            });
            m.fields = m.fields.map(function (field) {
                if (field.DisplayName == null) {
                    field.DisplayName = field.Name;
                }
                if (field.Tags == null) {
                    field.Tags = ['Level4'];
                }
                if (basetypeToJSType[field.BaseType] == null) {
                    throw new Error("base type " + field.BaseType + " not recognized for field " + field.Name);
                }
                field.jsType = basetypeToJSType[field.BaseType];
                if (field.BaseType === 'Bitmap') {
                    // Set as debug metric for now
                    field.tags = ["Level7"];
                }
                return _.transform(field, function (result, val, key) {
                    result[_.camelCase(key)] = val;
                });
            });
            messages.push(m);
        });
    });
    custom_metrics_1.customMetrics.forEach(function (metric) {
        messages.push(metric);
    });
    var data = {
        pipeline: pipeline,
        messages: messages,
    };
    var template = 'generate-pipeline-metrics-ts.hbs';
    var outputFile = outputFolder + pipeline.toLowerCase() + '_metadata.ts';
    utils.writeTemplate(template, data, outputFile);
}
exports.generatePipelineMetricMetadata = generatePipelineMetricMetadata;
function generateMetricMetadata(pipelines, outputFolder) {
    if (fs.existsSync(outputFolder)) {
        // Delete all contents
        rmdir.sync(outputFolder);
    }
    fs.mkdirSync(outputFolder, { recursive: true });
    pipelines.forEach(function (pipeline) {
        generatePipelineMetricMetadata(pipeline.pipeline, pipeline.basePath, outputFolder);
    });
    var template = 'generate-metrics-ts.hbs';
    var outputFile = outputFolder + 'metadata.ts';
    utils.writeTemplate(template, {}, outputFile);
}
exports.generateMetricMetadata = generateMetricMetadata;
