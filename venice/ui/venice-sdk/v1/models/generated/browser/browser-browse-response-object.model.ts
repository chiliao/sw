/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface IBrowserBrowseResponseObject {
    'root-uri'?: string;
    'query-type'?: string;
    'max-depth'?: number;
    'total-count'?: number;
    'objects'?: object;
}


export class BrowserBrowseResponseObject extends BaseModel implements IBrowserBrowseResponseObject {
    'root-uri': string = null;
    'query-type': string = null;
    'max-depth': number = null;
    'total-count': number = null;
    /** map of results. Key to the map is the URI of the  Object. */
    'objects': object = null;
    public static propInfo: { [prop: string]: PropInfoItem } = {
        'root-uri': {
            required: false,
            type: 'string'
        },
        'query-type': {
            required: false,
            type: 'string'
        },
        'max-depth': {
            required: false,
            type: 'number'
        },
        'total-count': {
            required: false,
            type: 'number'
        },
        'objects': {
            description:  'map of results. Key to the map is the URI of the  Object.',
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return BrowserBrowseResponseObject.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return BrowserBrowseResponseObject.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (BrowserBrowseResponseObject.propInfo[prop] != null &&
                        BrowserBrowseResponseObject.propInfo[prop].default != null);
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
        if (values && values['root-uri'] != null) {
            this['root-uri'] = values['root-uri'];
        } else if (fillDefaults && BrowserBrowseResponseObject.hasDefaultValue('root-uri')) {
            this['root-uri'] = BrowserBrowseResponseObject.propInfo['root-uri'].default;
        } else {
            this['root-uri'] = null
        }
        if (values && values['query-type'] != null) {
            this['query-type'] = values['query-type'];
        } else if (fillDefaults && BrowserBrowseResponseObject.hasDefaultValue('query-type')) {
            this['query-type'] = BrowserBrowseResponseObject.propInfo['query-type'].default;
        } else {
            this['query-type'] = null
        }
        if (values && values['max-depth'] != null) {
            this['max-depth'] = values['max-depth'];
        } else if (fillDefaults && BrowserBrowseResponseObject.hasDefaultValue('max-depth')) {
            this['max-depth'] = BrowserBrowseResponseObject.propInfo['max-depth'].default;
        } else {
            this['max-depth'] = null
        }
        if (values && values['total-count'] != null) {
            this['total-count'] = values['total-count'];
        } else if (fillDefaults && BrowserBrowseResponseObject.hasDefaultValue('total-count')) {
            this['total-count'] = BrowserBrowseResponseObject.propInfo['total-count'].default;
        } else {
            this['total-count'] = null
        }
        if (values && values['objects'] != null) {
            this['objects'] = values['objects'];
        } else if (fillDefaults && BrowserBrowseResponseObject.hasDefaultValue('objects')) {
            this['objects'] = BrowserBrowseResponseObject.propInfo['objects'].default;
        } else {
            this['objects'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'root-uri': CustomFormControl(new FormControl(this['root-uri']), BrowserBrowseResponseObject.propInfo['root-uri']),
                'query-type': CustomFormControl(new FormControl(this['query-type']), BrowserBrowseResponseObject.propInfo['query-type']),
                'max-depth': CustomFormControl(new FormControl(this['max-depth']), BrowserBrowseResponseObject.propInfo['max-depth']),
                'total-count': CustomFormControl(new FormControl(this['total-count']), BrowserBrowseResponseObject.propInfo['total-count']),
                'objects': CustomFormControl(new FormControl(this['objects']), BrowserBrowseResponseObject.propInfo['objects']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['root-uri'].setValue(this['root-uri']);
            this._formGroup.controls['query-type'].setValue(this['query-type']);
            this._formGroup.controls['max-depth'].setValue(this['max-depth']);
            this._formGroup.controls['total-count'].setValue(this['total-count']);
            this._formGroup.controls['objects'].setValue(this['objects']);
        }
    }
}

