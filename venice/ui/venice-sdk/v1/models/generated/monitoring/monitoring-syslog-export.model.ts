/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringSyslogExport_format,  MonitoringSyslogExport_format_uihint  } from './enums';
import { MonitoringExportConfig, IMonitoringExportConfig } from './monitoring-export-config.model';
import { MonitoringSyslogExportConfig, IMonitoringSyslogExportConfig } from './monitoring-syslog-export-config.model';

export interface IMonitoringSyslogExport {
    'format': MonitoringSyslogExport_format;
    'targets'?: Array<IMonitoringExportConfig>;
    'config'?: IMonitoringSyslogExportConfig;
    '_ui'?: any;
}


export class MonitoringSyslogExport extends BaseModel implements IMonitoringSyslogExport {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Event export format, SYSLOG_BSD default. */
    'format': MonitoringSyslogExport_format = null;
    /** Export target ip/port/protocol. */
    'targets': Array<MonitoringExportConfig> = null;
    /** Syslog specific configuration; one of the supported configs. */
    'config': MonitoringSyslogExportConfig = null;
    public static propInfo: { [prop in keyof IMonitoringSyslogExport]: PropInfoItem } = {
        'format': {
            enum: MonitoringSyslogExport_format_uihint,
            default: 'syslog-bsd',
            description:  `Event export format, SYSLOG_BSD default.`,
            required: true,
            type: 'string'
        },
        'targets': {
            description:  `Export target ip/port/protocol.`,
            required: false,
            type: 'object'
        },
        'config': {
            description:  `Syslog specific configuration; one of the supported configs.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringSyslogExport.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringSyslogExport.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringSyslogExport.propInfo[prop] != null &&
                        MonitoringSyslogExport.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['targets'] = new Array<MonitoringExportConfig>();
        this['config'] = new MonitoringSyslogExportConfig();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['_ui']) {
            this['_ui'] = values['_ui']
        }
        if (values && values['format'] != null) {
            this['format'] = values['format'];
        } else if (fillDefaults && MonitoringSyslogExport.hasDefaultValue('format')) {
            this['format'] = <MonitoringSyslogExport_format>  MonitoringSyslogExport.propInfo['format'].default;
        } else {
            this['format'] = null
        }
        if (values) {
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', values['targets'], MonitoringExportConfig);
        } else {
            this['targets'] = [];
        }
        if (values) {
            this['config'].setValues(values['config'], fillDefaults);
        } else {
            this['config'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'format': CustomFormControl(new FormControl(this['format'], [required, enumValidator(MonitoringSyslogExport_format), ]), MonitoringSyslogExport.propInfo['format']),
                'targets': new FormArray([]),
                'config': CustomFormGroup(this['config'].$formGroup, MonitoringSyslogExport.propInfo['config'].required),
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringExportConfig>('targets', this['targets'], MonitoringExportConfig);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('targets') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('targets').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('config') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('config').get(field);
                control.updateValueAndValidity();
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['format'].setValue(this['format']);
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', this['targets'], MonitoringExportConfig);
            this['config'].setFormGroupValuesToBeModelValues();
        }
    }
}

