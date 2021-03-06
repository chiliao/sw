/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { SecurityAutoMsgSecurityGroupWatchHelperWatchEvent, ISecurityAutoMsgSecurityGroupWatchHelperWatchEvent } from './security-auto-msg-security-group-watch-helper-watch-event.model';

export interface ISecurityAutoMsgSecurityGroupWatchHelper {
    'events'?: Array<ISecurityAutoMsgSecurityGroupWatchHelperWatchEvent>;
    '_ui'?: any;
}


export class SecurityAutoMsgSecurityGroupWatchHelper extends BaseModel implements ISecurityAutoMsgSecurityGroupWatchHelper {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'events': Array<SecurityAutoMsgSecurityGroupWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof ISecurityAutoMsgSecurityGroupWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SecurityAutoMsgSecurityGroupWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SecurityAutoMsgSecurityGroupWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SecurityAutoMsgSecurityGroupWatchHelper.propInfo[prop] != null &&
                        SecurityAutoMsgSecurityGroupWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<SecurityAutoMsgSecurityGroupWatchHelperWatchEvent>();
        this._inputValue = values;
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['_ui']) {
            this['_ui'] = values['_ui']
        }
        if (values) {
            this.fillModelArray<SecurityAutoMsgSecurityGroupWatchHelperWatchEvent>(this, 'events', values['events'], SecurityAutoMsgSecurityGroupWatchHelperWatchEvent);
        } else {
            this['events'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'events': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<SecurityAutoMsgSecurityGroupWatchHelperWatchEvent>('events', this['events'], SecurityAutoMsgSecurityGroupWatchHelperWatchEvent);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('events') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('events').get(field);
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
            this.fillModelArray<SecurityAutoMsgSecurityGroupWatchHelperWatchEvent>(this, 'events', this['events'], SecurityAutoMsgSecurityGroupWatchHelperWatchEvent);
        }
    }
}

