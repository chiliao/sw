/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent, ISecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent } from './security-auto-msg-traffic-encryption-policy-watch-helper-watch-event.model';

export interface ISecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
    'Events'?: Array<ISecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent>;
}


export class SecurityAutoMsgTrafficEncryptionPolicyWatchHelper extends BaseModel implements ISecurityAutoMsgTrafficEncryptionPolicyWatchHelper {
    'Events': Array<SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (SecurityAutoMsgTrafficEncryptionPolicyWatchHelper.enumProperties[prop] != null &&
                        SecurityAutoMsgTrafficEncryptionPolicyWatchHelper.enumProperties[prop].default != null &&
                        SecurityAutoMsgTrafficEncryptionPolicyWatchHelper.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Events'] = new Array<SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent>(this, 'Events', values['Events'], SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Events': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent>('Events', this['Events'], SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent>(this, 'Events', this['Events'], SecurityAutoMsgTrafficEncryptionPolicyWatchHelperWatchEvent);
        }
    }
}

