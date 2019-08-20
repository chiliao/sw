/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { BrowserBrowseRequestObject_query_type,  } from './enums';

export interface IBrowserBrowseRequestObject {
    'uri': string;
    'query-type': BrowserBrowseRequestObject_query_type;
    'max-depth'?: number;
    'count-only'?: boolean;
}


export class BrowserBrowseRequestObject extends BaseModel implements IBrowserBrowseRequestObject {
    /** length of string should be between 2 and 512 */
    'uri': string = null;
    'query-type': BrowserBrowseRequestObject_query_type = null;
    /** Max-Depth specifies how deep the query should explore. By default depth is set to 1 which means immediate relations
     0 means to maximum depth. */
    'max-depth': number = null;
    /** When CountOnly is set the response only contains counts and not the actual objects. */
    'count-only': boolean = null;
    public static propInfo: { [prop in keyof IBrowserBrowseRequestObject]: PropInfoItem } = {
        'uri': {
            description:  'length of string should be between 2 and 512',
            required: true,
            type: 'string'
        },
        'query-type': {
            enum: BrowserBrowseRequestObject_query_type,
            default: 'dependencies',
            required: true,
            type: 'string'
        },
        'max-depth': {
            default: parseInt('1'),
            description:  'Max-Depth specifies how deep the query should explore. By default depth is set to 1 which means immediate relations  0 means to maximum depth.',
            required: false,
            type: 'number'
        },
        'count-only': {
            description:  'When CountOnly is set the response only contains counts and not the actual objects.',
            required: false,
            type: 'boolean'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return BrowserBrowseRequestObject.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return BrowserBrowseRequestObject.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (BrowserBrowseRequestObject.propInfo[prop] != null &&
                        BrowserBrowseRequestObject.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['uri'] != null) {
            this['uri'] = values['uri'];
        } else if (fillDefaults && BrowserBrowseRequestObject.hasDefaultValue('uri')) {
            this['uri'] = BrowserBrowseRequestObject.propInfo['uri'].default;
        } else {
            this['uri'] = null
        }
        if (values && values['query-type'] != null) {
            this['query-type'] = values['query-type'];
        } else if (fillDefaults && BrowserBrowseRequestObject.hasDefaultValue('query-type')) {
            this['query-type'] = <BrowserBrowseRequestObject_query_type>  BrowserBrowseRequestObject.propInfo['query-type'].default;
        } else {
            this['query-type'] = null
        }
        if (values && values['max-depth'] != null) {
            this['max-depth'] = values['max-depth'];
        } else if (fillDefaults && BrowserBrowseRequestObject.hasDefaultValue('max-depth')) {
            this['max-depth'] = BrowserBrowseRequestObject.propInfo['max-depth'].default;
        } else {
            this['max-depth'] = null
        }
        if (values && values['count-only'] != null) {
            this['count-only'] = values['count-only'];
        } else if (fillDefaults && BrowserBrowseRequestObject.hasDefaultValue('count-only')) {
            this['count-only'] = BrowserBrowseRequestObject.propInfo['count-only'].default;
        } else {
            this['count-only'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'uri': CustomFormControl(new FormControl(this['uri'], [required, minLengthValidator(2), maxLengthValidator(512), ]), BrowserBrowseRequestObject.propInfo['uri']),
                'query-type': CustomFormControl(new FormControl(this['query-type'], [required, enumValidator(BrowserBrowseRequestObject_query_type), ]), BrowserBrowseRequestObject.propInfo['query-type']),
                'max-depth': CustomFormControl(new FormControl(this['max-depth']), BrowserBrowseRequestObject.propInfo['max-depth']),
                'count-only': CustomFormControl(new FormControl(this['count-only']), BrowserBrowseRequestObject.propInfo['count-only']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['uri'].setValue(this['uri']);
            this._formGroup.controls['query-type'].setValue(this['query-type']);
            this._formGroup.controls['max-depth'].setValue(this['max-depth']);
            this._formGroup.controls['count-only'].setValue(this['count-only']);
        }
    }
}

