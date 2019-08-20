/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { AuthTLSOptions, IAuthTLSOptions } from './auth-tls-options.model';

export interface IAuthLdapServer {
    'url'?: string;
    'tls-options'?: IAuthTLSOptions;
}


export class AuthLdapServer extends BaseModel implements IAuthLdapServer {
    'url': string = null;
    'tls-options': AuthTLSOptions = null;
    public static propInfo: { [prop in keyof IAuthLdapServer]: PropInfoItem } = {
        'url': {
            required: false,
            type: 'string'
        },
        'tls-options': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthLdapServer.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthLdapServer.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthLdapServer.propInfo[prop] != null &&
                        AuthLdapServer.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['tls-options'] = new AuthTLSOptions();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['url'] != null) {
            this['url'] = values['url'];
        } else if (fillDefaults && AuthLdapServer.hasDefaultValue('url')) {
            this['url'] = AuthLdapServer.propInfo['url'].default;
        } else {
            this['url'] = null
        }
        if (values) {
            this['tls-options'].setValues(values['tls-options'], fillDefaults);
        } else {
            this['tls-options'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'url': CustomFormControl(new FormControl(this['url']), AuthLdapServer.propInfo['url']),
                'tls-options': CustomFormGroup(this['tls-options'].$formGroup, AuthLdapServer.propInfo['tls-options'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('tls-options') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('tls-options').get(field);
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
            this._formGroup.controls['url'].setValue(this['url']);
            this['tls-options'].setFormGroupValuesToBeModelValues();
        }
    }
}

