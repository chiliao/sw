/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IAuditAuditEventRequest {
    'uuid'?: string;
    '_ui'?: any;
}


export class AuditAuditEventRequest extends BaseModel implements IAuditAuditEventRequest {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'uuid': string = null;
    public static propInfo: { [prop in keyof IAuditAuditEventRequest]: PropInfoItem } = {
        'uuid': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuditAuditEventRequest.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuditAuditEventRequest.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuditAuditEventRequest.propInfo[prop] != null &&
                        AuditAuditEventRequest.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
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
        if (values && values['uuid'] != null) {
            this['uuid'] = values['uuid'];
        } else if (fillDefaults && AuditAuditEventRequest.hasDefaultValue('uuid')) {
            this['uuid'] = AuditAuditEventRequest.propInfo['uuid'].default;
        } else {
            this['uuid'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'uuid': CustomFormControl(new FormControl(this['uuid']), AuditAuditEventRequest.propInfo['uuid']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['uuid'].setValue(this['uuid']);
        }
    }
}

