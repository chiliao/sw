/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiStatusResult, IApiStatusResult } from './api-status-result.model';
import { ApiObjectRef, IApiObjectRef } from './api-object-ref.model';

export interface IApiStatus {
    'kind'?: string;
    'api-version'?: string;
    'result'?: IApiStatusResult;
    'message'?: Array<string>;
    'code'?: number;
    'object-ref'?: IApiObjectRef;
}


export class ApiStatus extends BaseModel implements IApiStatus {
    'kind': string = null;
    'api-version': string = null;
    /** Result contains the status of the operation, success or failure. */
    'result': ApiStatusResult = null;
    /** Message contains human readable form of the error. */
    'message': Array<string> = null;
    /** Code is the HTTP status code. */
    'code': number = null;
    /** Reference to the object (optional) for which this status is being sent. */
    'object-ref': ApiObjectRef = null;
    public static propInfo: { [prop in keyof IApiStatus]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'result': {
            description:  'Result contains the status of the operation, success or failure.',
            required: false,
            type: 'object'
        },
        'message': {
            description:  'Message contains human readable form of the error.',
            required: false,
            type: 'Array<string>'
        },
        'code': {
            description:  'Code is the HTTP status code.',
            required: false,
            type: 'number'
        },
        'object-ref': {
            description:  'Reference to the object (optional) for which this status is being sent.',
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ApiStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ApiStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ApiStatus.propInfo[prop] != null &&
                        ApiStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['result'] = new ApiStatusResult();
        this['message'] = new Array<string>();
        this['object-ref'] = new ApiObjectRef();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && ApiStatus.hasDefaultValue('kind')) {
            this['kind'] = ApiStatus.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && ApiStatus.hasDefaultValue('api-version')) {
            this['api-version'] = ApiStatus.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['result'].setValues(values['result'], fillDefaults);
        } else {
            this['result'].setValues(null, fillDefaults);
        }
        if (values && values['message'] != null) {
            this['message'] = values['message'];
        } else if (fillDefaults && ApiStatus.hasDefaultValue('message')) {
            this['message'] = [ ApiStatus.propInfo['message'].default];
        } else {
            this['message'] = [];
        }
        if (values && values['code'] != null) {
            this['code'] = values['code'];
        } else if (fillDefaults && ApiStatus.hasDefaultValue('code')) {
            this['code'] = ApiStatus.propInfo['code'].default;
        } else {
            this['code'] = null
        }
        if (values) {
            this['object-ref'].setValues(values['object-ref'], fillDefaults);
        } else {
            this['object-ref'].setValues(null, fillDefaults);
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), ApiStatus.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), ApiStatus.propInfo['api-version']),
                'result': CustomFormGroup(this['result'].$formGroup, ApiStatus.propInfo['result'].required),
                'message': CustomFormControl(new FormControl(this['message']), ApiStatus.propInfo['message']),
                'code': CustomFormControl(new FormControl(this['code']), ApiStatus.propInfo['code']),
                'object-ref': CustomFormGroup(this['object-ref'].$formGroup, ApiStatus.propInfo['object-ref'].required),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('result') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('result').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('object-ref') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('object-ref').get(field);
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
            this._formGroup.controls['kind'].setValue(this['kind']);
            this._formGroup.controls['api-version'].setValue(this['api-version']);
            this['result'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['message'].setValue(this['message']);
            this._formGroup.controls['code'].setValue(this['code']);
            this['object-ref'].setFormGroupValuesToBeModelValues();
        }
    }
}

