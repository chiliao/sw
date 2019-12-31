/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterFeature {
    'feature-key'?: string;
    'licence'?: string;
}


export class ClusterFeature extends BaseModel implements IClusterFeature {
    'feature-key': string = null;
    'licence': string = null;
    public static propInfo: { [prop in keyof IClusterFeature]: PropInfoItem } = {
        'feature-key': {
            required: false,
            type: 'string'
        },
        'licence': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterFeature.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterFeature.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterFeature.propInfo[prop] != null &&
                        ClusterFeature.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['feature-key'] != null) {
            this['feature-key'] = values['feature-key'];
        } else if (fillDefaults && ClusterFeature.hasDefaultValue('feature-key')) {
            this['feature-key'] = ClusterFeature.propInfo['feature-key'].default;
        } else {
            this['feature-key'] = null
        }
        if (values && values['licence'] != null) {
            this['licence'] = values['licence'];
        } else if (fillDefaults && ClusterFeature.hasDefaultValue('licence')) {
            this['licence'] = ClusterFeature.propInfo['licence'].default;
        } else {
            this['licence'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'feature-key': CustomFormControl(new FormControl(this['feature-key']), ClusterFeature.propInfo['feature-key']),
                'licence': CustomFormControl(new FormControl(this['licence']), ClusterFeature.propInfo['licence']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['feature-key'].setValue(this['feature-key']);
            this._formGroup.controls['licence'].setValue(this['licence']);
        }
    }
}
