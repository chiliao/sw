/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { FieldsRequirement, IFieldsRequirement } from './fields-requirement.model';

export interface IFieldsSelector {
    'requirements'?: Array<IFieldsRequirement>;
}


export class FieldsSelector extends BaseModel implements IFieldsSelector {
    /** Requirements are ANDed. */
    'requirements': Array<FieldsRequirement> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (FieldsSelector.enumProperties[prop] != null &&
                        FieldsSelector.enumProperties[prop].default != null &&
                        FieldsSelector.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['requirements'] = new Array<FieldsRequirement>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<FieldsRequirement>(this, 'requirements', values['requirements'], FieldsRequirement);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'requirements': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<FieldsRequirement>('requirements', this['requirements'], FieldsRequirement);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<FieldsRequirement>(this, 'requirements', this['requirements'], FieldsRequirement);
        }
    }
}

