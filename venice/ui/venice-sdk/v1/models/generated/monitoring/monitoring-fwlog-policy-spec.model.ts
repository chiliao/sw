/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringExportConfig, IMonitoringExportConfig } from './monitoring-export-config.model';
import { MonitoringFwlogPolicySpec_format,  MonitoringFwlogPolicySpec_format_uihint  } from './enums';
import { MonitoringFwlogPolicySpec_filter,  MonitoringFwlogPolicySpec_filter_uihint  } from './enums';
import { MonitoringSyslogExportConfig, IMonitoringSyslogExportConfig } from './monitoring-syslog-export-config.model';

export interface IMonitoringFwlogPolicySpec {
    'vrf-name'?: string;
    'targets'?: Array<IMonitoringExportConfig>;
    'format': MonitoringFwlogPolicySpec_format;
    'filter': Array<MonitoringFwlogPolicySpec_filter>;
    'config'?: IMonitoringSyslogExportConfig;
}


export class MonitoringFwlogPolicySpec extends BaseModel implements IMonitoringFwlogPolicySpec {
    'vrf-name': string = null;
    'targets': Array<MonitoringExportConfig> = null;
    'format': MonitoringFwlogPolicySpec_format = null;
    'filter': Array<MonitoringFwlogPolicySpec_filter> = null;
    'config': MonitoringSyslogExportConfig = null;
    public static propInfo: { [prop in keyof IMonitoringFwlogPolicySpec]: PropInfoItem } = {
        'vrf-name': {
            required: false,
            type: 'string'
        },
        'targets': {
            required: false,
            type: 'object'
        },
        'format': {
            enum: MonitoringFwlogPolicySpec_format_uihint,
            default: 'syslog-bsd',
            required: true,
            type: 'string'
        },
        'filter': {
            enum: MonitoringFwlogPolicySpec_filter_uihint,
            default: 'all',
            required: true,
            type: 'Array<string>'
        },
        'config': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringFwlogPolicySpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringFwlogPolicySpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringFwlogPolicySpec.propInfo[prop] != null &&
                        MonitoringFwlogPolicySpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['targets'] = new Array<MonitoringExportConfig>();
        this['filter'] = new Array<MonitoringFwlogPolicySpec_filter>();
        this['config'] = new MonitoringSyslogExportConfig();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['vrf-name'] != null) {
            this['vrf-name'] = values['vrf-name'];
        } else if (fillDefaults && MonitoringFwlogPolicySpec.hasDefaultValue('vrf-name')) {
            this['vrf-name'] = MonitoringFwlogPolicySpec.propInfo['vrf-name'].default;
        } else {
            this['vrf-name'] = null
        }
        if (values) {
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', values['targets'], MonitoringExportConfig);
        } else {
            this['targets'] = [];
        }
        if (values && values['format'] != null) {
            this['format'] = values['format'];
        } else if (fillDefaults && MonitoringFwlogPolicySpec.hasDefaultValue('format')) {
            this['format'] = <MonitoringFwlogPolicySpec_format>  MonitoringFwlogPolicySpec.propInfo['format'].default;
        } else {
            this['format'] = null
        }
        if (values && values['filter'] != null) {
            this['filter'] = values['filter'];
        } else if (fillDefaults && MonitoringFwlogPolicySpec.hasDefaultValue('filter')) {
            this['filter'] = [ MonitoringFwlogPolicySpec.propInfo['filter'].default];
        } else {
            this['filter'] = [];
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
                'vrf-name': CustomFormControl(new FormControl(this['vrf-name']), MonitoringFwlogPolicySpec.propInfo['vrf-name']),
                'targets': new FormArray([]),
                'format': CustomFormControl(new FormControl(this['format'], [required, enumValidator(MonitoringFwlogPolicySpec_format), ]), MonitoringFwlogPolicySpec.propInfo['format']),
                'filter': CustomFormControl(new FormControl(this['filter']), MonitoringFwlogPolicySpec.propInfo['filter']),
                'config': CustomFormGroup(this['config'].$formGroup, MonitoringFwlogPolicySpec.propInfo['config'].required),
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
            this._formGroup.controls['vrf-name'].setValue(this['vrf-name']);
            this.fillModelArray<MonitoringExportConfig>(this, 'targets', this['targets'], MonitoringExportConfig);
            this._formGroup.controls['format'].setValue(this['format']);
            this._formGroup.controls['filter'].setValue(this['filter']);
            this['config'].setFormGroupValuesToBeModelValues();
        }
    }
}

