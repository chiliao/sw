/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';


export interface ISearchKindPreview {
    'kinds'?: object;
    '_ui'?: any;
}


export class SearchKindPreview extends BaseModel implements ISearchKindPreview {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'kinds': object = null;
    public static propInfo: { [prop in keyof ISearchKindPreview]: PropInfoItem } = {
        'kinds': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return SearchKindPreview.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return SearchKindPreview.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (SearchKindPreview.propInfo[prop] != null &&
                        SearchKindPreview.propInfo[prop].default != null);
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
        if (values && values['kinds'] != null) {
            this['kinds'] = values['kinds'];
        } else if (fillDefaults && SearchKindPreview.hasDefaultValue('kinds')) {
            this['kinds'] = SearchKindPreview.propInfo['kinds'].default;
        } else {
            this['kinds'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kinds': CustomFormControl(new FormControl(this['kinds']), SearchKindPreview.propInfo['kinds']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['kinds'].setValue(this['kinds']);
        }
    }
}

