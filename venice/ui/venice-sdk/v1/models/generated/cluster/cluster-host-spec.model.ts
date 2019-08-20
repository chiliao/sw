/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterSmartNICID, IClusterSmartNICID } from './cluster-smart-nicid.model';

export interface IClusterHostSpec {
    'smart-nics'?: Array<IClusterSmartNICID>;
}


export class ClusterHostSpec extends BaseModel implements IClusterHostSpec {
    'smart-nics': Array<ClusterSmartNICID> = null;
    public static propInfo: { [prop in keyof IClusterHostSpec]: PropInfoItem } = {
        'smart-nics': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterHostSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterHostSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterHostSpec.propInfo[prop] != null &&
                        ClusterHostSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['smart-nics'] = new Array<ClusterSmartNICID>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<ClusterSmartNICID>(this, 'smart-nics', values['smart-nics'], ClusterSmartNICID);
        } else {
            this['smart-nics'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'smart-nics': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<ClusterSmartNICID>('smart-nics', this['smart-nics'], ClusterSmartNICID);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('smart-nics') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('smart-nics').get(field);
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
            this.fillModelArray<ClusterSmartNICID>(this, 'smart-nics', this['smart-nics'], ClusterSmartNICID);
        }
    }
}

