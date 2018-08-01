/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { NetworkLbPolicySpec, INetworkLbPolicySpec } from './network-lb-policy-spec.model';
import { NetworkLbPolicyStatus, INetworkLbPolicyStatus } from './network-lb-policy-status.model';

export interface INetworkLbPolicy {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'spec'?: INetworkLbPolicySpec;
    'status'?: INetworkLbPolicyStatus;
}


export class NetworkLbPolicy extends BaseModel implements INetworkLbPolicy {
    'kind': string = null;
    'api-version': string = null;
    'meta': ApiObjectMeta = null;
    /** Spec contains the configuration of the LbPolicy. */
    'spec': NetworkLbPolicySpec = null;
    /** Status contains the current state of the LbPolicy. */
    'status': NetworkLbPolicyStatus = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (NetworkLbPolicy.enumProperties[prop] != null &&
                        NetworkLbPolicy.enumProperties[prop].default != null &&
                        NetworkLbPolicy.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['spec'] = new NetworkLbPolicySpec();
        this['status'] = new NetworkLbPolicyStatus();
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

