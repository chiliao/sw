/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { SecurityTrafficEncryptionPolicy, ISecurityTrafficEncryptionPolicy } from './security-traffic-encryption-policy.model';

export interface ISecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent {
    'type'?: string;
    'object'?: ISecurityTrafficEncryptionPolicy;
}


export class SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent extends BaseModel implements ISecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent {
    'type': string = null;
    'object': SecurityTrafficEncryptionPolicy = null;
    public static propInfo: { [prop in keyof ISecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent]: PropInfoItem } = {
        'type': {
            required: false,
            type: 'string'
        },
        'object': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo[prop] != null &&
                        SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['object'] = new SecurityTrafficEncryptionPolicy();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['type'] != null) {
            this['type'] = values['type'];
        } else if (fillDefaults && SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.hasDefaultValue('type')) {
            this['type'] = SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo['type'].default;
        } else {
            this['type'] = null
        }
        if (values) {
            this['object'].setValues(values['object'], fillDefaults);
        } else {
            this['object'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'type': CustomFormControl(new FormControl(this['type']), SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo['type']),
                'object': CustomFormGroup(this['object'].$formGroup, SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent.propInfo['object'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('object') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('object').get(field);
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
            this._formGroup.controls['type'].setValue(this['type']);
            this['object'].setFormGroupValuesToBeModelValues();
        }
    }
}

