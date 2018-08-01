/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { SecurityTrafficEncryptionPolicySpec, ISecurityTrafficEncryptionPolicySpec } from './security-traffic-encryption-policy-spec.model';
import { SecurityTrafficEncryptionPolicyStatus, ISecurityTrafficEncryptionPolicyStatus } from './security-traffic-encryption-policy-status.model';

export interface ISecurityTrafficEncryptionPolicy {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'spec'?: ISecurityTrafficEncryptionPolicySpec;
    'status'?: ISecurityTrafficEncryptionPolicyStatus;
}


export class SecurityTrafficEncryptionPolicy extends BaseModel implements ISecurityTrafficEncryptionPolicy {
    'kind': string = null;
    'api-version': string = null;
    'meta': ApiObjectMeta = null;
    /** Spec contains the configuration of the encryption policy. */
    'spec': SecurityTrafficEncryptionPolicySpec = null;
    /** Status contains the current state of the encryption policy. */
    'status': SecurityTrafficEncryptionPolicyStatus = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (SecurityTrafficEncryptionPolicy.enumProperties[prop] != null &&
                        SecurityTrafficEncryptionPolicy.enumProperties[prop].default != null &&
                        SecurityTrafficEncryptionPolicy.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['spec'] = new SecurityTrafficEncryptionPolicySpec();
        this['status'] = new SecurityTrafficEncryptionPolicyStatus();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        }
        if (values) {
            this['meta'].setValues(values['meta']);
        }
        if (values) {
            this['spec'].setValues(values['spec']);
        }
        if (values) {
            this['status'].setValues(values['status']);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': new FormControl(this['kind']),
                'api-version': new FormControl(this['api-version']),
                'meta': this['meta'].$formGroup,
                'spec': this['spec'].$formGroup,
                'status': this['status'].$formGroup,
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this['meta'].setFormGroupValues();
            this['spec'].setFormGroupValues();
            this['status'].setFormGroupValues();
        }
    }
}

