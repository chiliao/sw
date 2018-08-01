/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { MonitoringFlowExportPolicy, IMonitoringFlowExportPolicy } from './monitoring-flow-export-policy.model';

export interface IMonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent {
    'Type'?: string;
    'Object'?: IMonitoringFlowExportPolicy;
}


export class MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent extends BaseModel implements IMonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent {
    'Type': string = null;
    'Object': MonitoringFlowExportPolicy = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent.enumProperties[prop] != null &&
                        MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent.enumProperties[prop].default != null &&
                        MonitoringAutoMsgFlowExportPolicyWatchHelperWatchEvent.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Object'] = new MonitoringFlowExportPolicy();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values && values['Type'] != null) {
            this['Type'] = values['Type'];
        }
        if (values) {
            this['Object'].setValues(values['Object']);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Type': new FormControl(this['Type']),
                'Object': this['Object'].$formGroup,
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['Type'].setValue(this['Type']);
            this['Object'].setFormGroupValues();
        }
    }
}

