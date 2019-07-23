/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from './base-model';

import { AuthLdapAttributeMapping, IAuthLdapAttributeMapping } from './auth-ldap-attribute-mapping.model';
import { AuthLdapServer, IAuthLdapServer } from './auth-ldap-server.model';

export interface IAuthLdap {
    'enabled'?: boolean;
    'base-dn'?: string;
    'bind-dn'?: string;
    'bind-password'?: string;
    'attribute-mapping'?: IAuthLdapAttributeMapping;
    'servers'?: Array<IAuthLdapServer>;
}


export class AuthLdap extends BaseModel implements IAuthLdap {
    'enabled': boolean = null;
    /** The LDAP base DN to be used in a user search. */
    'base-dn': string = null;
    /** The bind DN is the string that Venice uses to log in to the LDAP server. Venice uses this account to validate the remote user attempting to log in. The base DN is the container name and path in the LDAPserver where Venice searches for the remote user account. This is where the password is validated. This contains the user authorization and assigned RBAC roles for use on Venice. Venice requests the attribute from theLDAP server. */
    'bind-dn': string = null;
    /** The password for the LDAP database account specified in the Root DN field. */
    'bind-password': string = null;
    'attribute-mapping': AuthLdapAttributeMapping = null;
    'servers': Array<AuthLdapServer> = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'enabled': {
            required: false,
            type: 'boolean'
        },
        'base-dn': {
            description:  'The LDAP base DN to be used in a user search.',
            required: false,
            type: 'string'
        },
        'bind-dn': {
            description:  'The bind DN is the string that Venice uses to log in to the LDAP server. Venice uses this account to validate the remote user attempting to log in. The base DN is the container name and path in the LDAPserver where Venice searches for the remote user account. This is where the password is validated. This contains the user authorization and assigned RBAC roles for use on Venice. Venice requests the attribute from theLDAP server.',
            required: false,
            type: 'string'
        },
        'bind-password': {
            description:  'The password for the LDAP database account specified in the Root DN field.',
            required: false,
            type: 'string'
        },
        'attribute-mapping': {
            required: false,
            type: 'object'
        },
        'servers': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthLdap.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthLdap.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthLdap.propInfo[prop] != null &&
                        AuthLdap.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['attribute-mapping'] = new AuthLdapAttributeMapping();
        this['servers'] = new Array<AuthLdapServer>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['enabled'] != null) {
            this['enabled'] = values['enabled'];
        } else if (fillDefaults && AuthLdap.hasDefaultValue('enabled')) {
            this['enabled'] = AuthLdap.propInfo['enabled'].default;
        } else {
            this['enabled'] = null
        }
        if (values && values['base-dn'] != null) {
            this['base-dn'] = values['base-dn'];
        } else if (fillDefaults && AuthLdap.hasDefaultValue('base-dn')) {
            this['base-dn'] = AuthLdap.propInfo['base-dn'].default;
        } else {
            this['base-dn'] = null
        }
        if (values && values['bind-dn'] != null) {
            this['bind-dn'] = values['bind-dn'];
        } else if (fillDefaults && AuthLdap.hasDefaultValue('bind-dn')) {
            this['bind-dn'] = AuthLdap.propInfo['bind-dn'].default;
        } else {
            this['bind-dn'] = null
        }
        if (values && values['bind-password'] != null) {
            this['bind-password'] = values['bind-password'];
        } else if (fillDefaults && AuthLdap.hasDefaultValue('bind-password')) {
            this['bind-password'] = AuthLdap.propInfo['bind-password'].default;
        } else {
            this['bind-password'] = null
        }
        if (values) {
            this['attribute-mapping'].setValues(values['attribute-mapping'], fillDefaults);
        } else {
            this['attribute-mapping'].setValues(null, fillDefaults);
        }
        if (values) {
            this.fillModelArray<AuthLdapServer>(this, 'servers', values['servers'], AuthLdapServer);
        } else {
            this['servers'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'enabled': CustomFormControl(new FormControl(this['enabled']), AuthLdap.propInfo['enabled']),
                'base-dn': CustomFormControl(new FormControl(this['base-dn']), AuthLdap.propInfo['base-dn']),
                'bind-dn': CustomFormControl(new FormControl(this['bind-dn']), AuthLdap.propInfo['bind-dn']),
                'bind-password': CustomFormControl(new FormControl(this['bind-password']), AuthLdap.propInfo['bind-password']),
                'attribute-mapping': CustomFormGroup(this['attribute-mapping'].$formGroup, AuthLdap.propInfo['attribute-mapping'].required),
                'servers': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<AuthLdapServer>('servers', this['servers'], AuthLdapServer);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('attribute-mapping') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('attribute-mapping').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('servers') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('servers').get(field);
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
            this._formGroup.controls['enabled'].setValue(this['enabled']);
            this._formGroup.controls['base-dn'].setValue(this['base-dn']);
            this._formGroup.controls['bind-dn'].setValue(this['bind-dn']);
            this._formGroup.controls['bind-password'].setValue(this['bind-password']);
            this['attribute-mapping'].setFormGroupValuesToBeModelValues();
            this.fillModelArray<AuthLdapServer>(this, 'servers', this['servers'], AuthLdapServer);
        }
    }
}

