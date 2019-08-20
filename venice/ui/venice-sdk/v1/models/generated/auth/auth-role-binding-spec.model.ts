/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IAuthRoleBindingSpec {
    'users'?: Array<string>;
    'user-groups'?: Array<string>;
    'role'?: string;
}


export class AuthRoleBindingSpec extends BaseModel implements IAuthRoleBindingSpec {
    'users': Array<string> = null;
    'user-groups': Array<string> = null;
    'role': string = null;
    public static propInfo: { [prop in keyof IAuthRoleBindingSpec]: PropInfoItem } = {
        'users': {
            required: false,
            type: 'Array<string>'
        },
        'user-groups': {
            required: false,
            type: 'Array<string>'
        },
        'role': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuthRoleBindingSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuthRoleBindingSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuthRoleBindingSpec.propInfo[prop] != null &&
                        AuthRoleBindingSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['users'] = new Array<string>();
        this['user-groups'] = new Array<string>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['users'] != null) {
            this['users'] = values['users'];
        } else if (fillDefaults && AuthRoleBindingSpec.hasDefaultValue('users')) {
            this['users'] = [ AuthRoleBindingSpec.propInfo['users'].default];
        } else {
            this['users'] = [];
        }
        if (values && values['user-groups'] != null) {
            this['user-groups'] = values['user-groups'];
        } else if (fillDefaults && AuthRoleBindingSpec.hasDefaultValue('user-groups')) {
            this['user-groups'] = [ AuthRoleBindingSpec.propInfo['user-groups'].default];
        } else {
            this['user-groups'] = [];
        }
        if (values && values['role'] != null) {
            this['role'] = values['role'];
        } else if (fillDefaults && AuthRoleBindingSpec.hasDefaultValue('role')) {
            this['role'] = AuthRoleBindingSpec.propInfo['role'].default;
        } else {
            this['role'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'users': CustomFormControl(new FormControl(this['users']), AuthRoleBindingSpec.propInfo['users']),
                'user-groups': CustomFormControl(new FormControl(this['user-groups']), AuthRoleBindingSpec.propInfo['user-groups']),
                'role': CustomFormControl(new FormControl(this['role']), AuthRoleBindingSpec.propInfo['role']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['users'].setValue(this['users']);
            this._formGroup.controls['user-groups'].setValue(this['user-groups']);
            this._formGroup.controls['role'].setValue(this['role']);
        }
    }
}

