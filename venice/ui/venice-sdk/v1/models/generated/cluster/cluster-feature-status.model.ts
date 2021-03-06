/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IClusterFeatureStatus {
    'feature-key'?: string;
    'value'?: string;
    'expiry'?: string;
    '_ui'?: any;
}


export class ClusterFeatureStatus extends BaseModel implements IClusterFeatureStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'feature-key': string = null;
    'value': string = null;
    'expiry': string = null;
    public static propInfo: { [prop in keyof IClusterFeatureStatus]: PropInfoItem } = {
        'feature-key': {
            required: false,
            type: 'string'
        },
        'value': {
            required: false,
            type: 'string'
        },
        'expiry': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterFeatureStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterFeatureStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterFeatureStatus.propInfo[prop] != null &&
                        ClusterFeatureStatus.propInfo[prop].default != null);
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
        if (values && values['_ui']) {
            this['_ui'] = values['_ui']
        }
        if (values && values['feature-key'] != null) {
            this['feature-key'] = values['feature-key'];
        } else if (fillDefaults && ClusterFeatureStatus.hasDefaultValue('feature-key')) {
            this['feature-key'] = ClusterFeatureStatus.propInfo['feature-key'].default;
        } else {
            this['feature-key'] = null
        }
        if (values && values['value'] != null) {
            this['value'] = values['value'];
        } else if (fillDefaults && ClusterFeatureStatus.hasDefaultValue('value')) {
            this['value'] = ClusterFeatureStatus.propInfo['value'].default;
        } else {
            this['value'] = null
        }
        if (values && values['expiry'] != null) {
            this['expiry'] = values['expiry'];
        } else if (fillDefaults && ClusterFeatureStatus.hasDefaultValue('expiry')) {
            this['expiry'] = ClusterFeatureStatus.propInfo['expiry'].default;
        } else {
            this['expiry'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'feature-key': CustomFormControl(new FormControl(this['feature-key']), ClusterFeatureStatus.propInfo['feature-key']),
                'value': CustomFormControl(new FormControl(this['value']), ClusterFeatureStatus.propInfo['value']),
                'expiry': CustomFormControl(new FormControl(this['expiry']), ClusterFeatureStatus.propInfo['expiry']),
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
            this._formGroup.controls['value'].setValue(this['value']);
            this._formGroup.controls['expiry'].setValue(this['expiry']);
        }
    }
}

