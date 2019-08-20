/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent, IMonitoringAutoMsgAlertPolicyWatchHelperWatchEvent } from './monitoring-auto-msg-alert-policy-watch-helper-watch-event.model';

export interface IMonitoringAutoMsgAlertPolicyWatchHelper {
    'events'?: Array<IMonitoringAutoMsgAlertPolicyWatchHelperWatchEvent>;
}


export class MonitoringAutoMsgAlertPolicyWatchHelper extends BaseModel implements IMonitoringAutoMsgAlertPolicyWatchHelper {
    'events': Array<MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent> = null;
    public static propInfo: { [prop in keyof IMonitoringAutoMsgAlertPolicyWatchHelper]: PropInfoItem } = {
        'events': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return MonitoringAutoMsgAlertPolicyWatchHelper.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return MonitoringAutoMsgAlertPolicyWatchHelper.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (MonitoringAutoMsgAlertPolicyWatchHelper.propInfo[prop] != null &&
                        MonitoringAutoMsgAlertPolicyWatchHelper.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['events'] = new Array<MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent>(this, 'events', values['events'], MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent);
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
            this.fillFormArray<MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent>('events', this['events'], MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent);
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
            this.fillModelArray<MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent>(this, 'events', this['events'], MonitoringAutoMsgAlertPolicyWatchHelperWatchEvent);
        }
    }
}

