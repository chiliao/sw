/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { AuthAutoMsgRoleBindingWatchHelperWatchEvent, IAuthAutoMsgRoleBindingWatchHelperWatchEvent } from './auth-auto-msg-role-binding-watch-helper-watch-event.model';

export interface IAuthAutoMsgRoleBindingWatchHelper {
    'Events'?: Array<IAuthAutoMsgRoleBindingWatchHelperWatchEvent>;
}


export class AuthAutoMsgRoleBindingWatchHelper extends BaseModel implements IAuthAutoMsgRoleBindingWatchHelper {
    'Events': Array<AuthAutoMsgRoleBindingWatchHelperWatchEvent> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (AuthAutoMsgRoleBindingWatchHelper.enumProperties[prop] != null &&
                        AuthAutoMsgRoleBindingWatchHelper.enumProperties[prop].default != null &&
                        AuthAutoMsgRoleBindingWatchHelper.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Events'] = new Array<AuthAutoMsgRoleBindingWatchHelperWatchEvent>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<AuthAutoMsgRoleBindingWatchHelperWatchEvent>(this, 'Events', values['Events'], AuthAutoMsgRoleBindingWatchHelperWatchEvent);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Events': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<AuthAutoMsgRoleBindingWatchHelperWatchEvent>('Events', this['Events'], AuthAutoMsgRoleBindingWatchHelperWatchEvent);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<AuthAutoMsgRoleBindingWatchHelperWatchEvent>(this, 'Events', this['Events'], AuthAutoMsgRoleBindingWatchHelperWatchEvent);
        }
    }
}

