/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { MonitoringFlowExportTarget, IMonitoringFlowExportTarget } from './monitoring-flow-export-target.model';

export interface IMonitoringFlowExportSpec {
    'targets'?: Array<IMonitoringFlowExportTarget>;
}


export class MonitoringFlowExportSpec extends BaseModel implements IMonitoringFlowExportSpec {
    'targets': Array<MonitoringFlowExportTarget> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (MonitoringFlowExportSpec.enumProperties[prop] != null &&
                        MonitoringFlowExportSpec.enumProperties[prop].default != null &&
                        MonitoringFlowExportSpec.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['targets'] = new Array<MonitoringFlowExportTarget>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<MonitoringFlowExportTarget>(this, 'targets', values['targets'], MonitoringFlowExportTarget);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'targets': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<MonitoringFlowExportTarget>('targets', this['targets'], MonitoringFlowExportTarget);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<MonitoringFlowExportTarget>(this, 'targets', this['targets'], MonitoringFlowExportTarget);
        }
    }
}

