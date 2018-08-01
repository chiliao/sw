/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { SecuritySGRule, ISecuritySGRule } from './security-sg-rule.model';

export interface ISecuritySGPolicySpec {
    'attach-groups'?: Array<string>;
    'attach-tenant'?: boolean;
    'rules'?: Array<ISecuritySGRule>;
}


export class SecuritySGPolicySpec extends BaseModel implements ISecuritySGPolicySpec {
    'attach-groups': Array<string> = null;
    'attach-tenant': boolean = null;
    'rules': Array<SecuritySGRule> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (SecuritySGPolicySpec.enumProperties[prop] != null &&
                        SecuritySGPolicySpec.enumProperties[prop].default != null &&
                        SecuritySGPolicySpec.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['attach-groups'] = new Array<string>();
        this['rules'] = new Array<SecuritySGRule>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<string>(this, 'attach-groups', values['attach-groups']);
        }
        if (values && values['attach-tenant'] != null) {
            this['attach-tenant'] = values['attach-tenant'];
        }
        if (values) {
            this.fillModelArray<SecuritySGRule>(this, 'rules', values['rules'], SecuritySGRule);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'attach-groups': new FormArray([]),
                'attach-tenant': new FormControl(this['attach-tenant']),
                'rules': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<string>('attach-groups', this['attach-groups']);
            // generate FormArray control elements
            this.fillFormArray<SecuritySGRule>('rules', this['rules'], SecuritySGRule);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<string>(this, 'attach-groups', this['attach-groups']);
            this._formGroup.controls['attach-tenant'].setValue(this['attach-tenant']);
            this.fillModelArray<SecuritySGRule>(this, 'rules', this['rules'], SecuritySGRule);
        }
    }
}

