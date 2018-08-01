/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { AuthRadiusServer_auth_method,  } from './enums';

export interface IAuthRadiusServer {
    'url'?: string;
    'secret'?: string;
    'auth-method'?: AuthRadiusServer_auth_method;
    'trusted-certs'?: string;
}


export class AuthRadiusServer extends BaseModel implements IAuthRadiusServer {
    'url': string = null;
    'secret': string = null;
    'auth-method': AuthRadiusServer_auth_method = null;
    'trusted-certs': string = null;
    public static enumProperties: { [key: string] : EnumDef } = {
        'auth-method': {
            enum: AuthRadiusServer_auth_method,
            default: 'PAP',
        },
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (AuthRadiusServer.enumProperties[prop] != null &&
                        AuthRadiusServer.enumProperties[prop].default != null &&
                        AuthRadiusServer.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values && values['url'] != null) {
            this['url'] = values['url'];
        }
        if (values && values['secret'] != null) {
            this['secret'] = values['secret'];
        }
        if (values && values['auth-method'] != null) {
            this['auth-method'] = values['auth-method'];
        } else if (AuthRadiusServer.hasDefaultEnumValue('auth-method')) {
            this['auth-method'] = <AuthRadiusServer_auth_method> AuthRadiusServer.enumProperties['auth-method'].default;
        }
        if (values && values['trusted-certs'] != null) {
            this['trusted-certs'] = values['trusted-certs'];
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'url': new FormControl(this['url']),
                'secret': new FormControl(this['secret']),
                'auth-method': new FormControl(this['auth-method'], [enumValidator(AuthRadiusServer_auth_method), ]),
                'trusted-certs': new FormControl(this['trusted-certs']),
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['url'].setValue(this['url']);
            this._formGroup.controls['secret'].setValue(this['secret']);
            this._formGroup.controls['auth-method'].setValue(this['auth-method']);
            this._formGroup.controls['trusted-certs'].setValue(this['trusted-certs']);
        }
    }
}

