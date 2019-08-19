/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { BrowserBrowseRequest_query_type,  } from './enums';

export interface IBrowserBrowseRequest {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'uri'?: string;
    'query-type'?: BrowserBrowseRequest_query_type;
    'max-depth'?: number;
    'count-only'?: boolean;
}


export class BrowserBrowseRequest extends BaseModel implements IBrowserBrowseRequest {
    'kind': string = null;
    'api-version': string = null;
    'meta': ApiObjectMeta = null;
    /** length of string should be between 2 and 512 */
    'uri': string = null;
    'query-type': BrowserBrowseRequest_query_type = null;
    'max-depth': number = null;
    'count-only': boolean = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'meta': {
            required: false,
            type: 'object'
        },
        'uri': {
            description:  'length of string should be between 2 and 512',
            required: false,
            type: 'string'
        },
        'query-type': {
            enum: BrowserBrowseRequest_query_type,
            default: 'dependencies',
            required: false,
            type: 'string'
        },
        'max-depth': {
            default: parseInt('1'),
            required: false,
            type: 'number'
        },
        'count-only': {
            required: false,
            type: 'boolean'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return BrowserBrowseRequest.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return BrowserBrowseRequest.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (BrowserBrowseRequest.propInfo[prop] != null &&
                        BrowserBrowseRequest.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['meta'] = new ApiObjectMeta();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('kind')) {
            this['kind'] = BrowserBrowseRequest.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('api-version')) {
            this['api-version'] = BrowserBrowseRequest.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['meta'].setValues(values['meta'], fillDefaults);
        } else {
            this['meta'].setValues(null, fillDefaults);
        }
        if (values && values['uri'] != null) {
            this['uri'] = values['uri'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('uri')) {
            this['uri'] = BrowserBrowseRequest.propInfo['uri'].default;
        } else {
            this['uri'] = null
        }
        if (values && values['query-type'] != null) {
            this['query-type'] = values['query-type'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('query-type')) {
            this['query-type'] = <BrowserBrowseRequest_query_type>  BrowserBrowseRequest.propInfo['query-type'].default;
        } else {
            this['query-type'] = null
        }
        if (values && values['max-depth'] != null) {
            this['max-depth'] = values['max-depth'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('max-depth')) {
            this['max-depth'] = BrowserBrowseRequest.propInfo['max-depth'].default;
        } else {
            this['max-depth'] = null
        }
        if (values && values['count-only'] != null) {
            this['count-only'] = values['count-only'];
        } else if (fillDefaults && BrowserBrowseRequest.hasDefaultValue('count-only')) {
            this['count-only'] = BrowserBrowseRequest.propInfo['count-only'].default;
        } else {
            this['count-only'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), BrowserBrowseRequest.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), BrowserBrowseRequest.propInfo['api-version']),
                'meta': CustomFormGroup(this['meta'].$formGroup, BrowserBrowseRequest.propInfo['meta'].required),
                'uri': CustomFormControl(new FormControl(this['uri'], [minLengthValidator(2), maxLengthValidator(512), ]), BrowserBrowseRequest.propInfo['uri']),
                'query-type': CustomFormControl(new FormControl(this['query-type'], [enumValidator(BrowserBrowseRequest_query_type), ]), BrowserBrowseRequest.propInfo['query-type']),
                'max-depth': CustomFormControl(new FormControl(this['max-depth']), BrowserBrowseRequest.propInfo['max-depth']),
                'count-only': CustomFormControl(new FormControl(this['count-only']), BrowserBrowseRequest.propInfo['count-only']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('meta') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('meta').get(field);
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
            this['meta'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['uri'].setValue(this['uri']);
            this._formGroup.controls['query-type'].setValue(this['query-type']);
            this._formGroup.controls['max-depth'].setValue(this['max-depth']);
            this._formGroup.controls['count-only'].setValue(this['count-only']);
        }
    }
}

