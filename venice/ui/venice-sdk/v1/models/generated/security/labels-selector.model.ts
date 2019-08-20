/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { LabelsRequirement, ILabelsRequirement } from './labels-requirement.model';

export interface ILabelsSelector {
    'requirements'?: Array<ILabelsRequirement>;
}


export class LabelsSelector extends BaseModel implements ILabelsSelector {
    /** Requirements are ANDed. */
    'requirements': Array<LabelsRequirement> = null;
    public static propInfo: { [prop in keyof ILabelsSelector]: PropInfoItem } = {
        'requirements': {
            description:  'Requirements are ANDed.',
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return LabelsSelector.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return LabelsSelector.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (LabelsSelector.propInfo[prop] != null &&
                        LabelsSelector.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['requirements'] = new Array<LabelsRequirement>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<LabelsRequirement>(this, 'requirements', values['requirements'], LabelsRequirement);
        } else {
            this['requirements'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'requirements': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<LabelsRequirement>('requirements', this['requirements'], LabelsRequirement);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('requirements') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('requirements').get(field);
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
            this.fillModelArray<LabelsRequirement>(this, 'requirements', this['requirements'], LabelsRequirement);
        }
    }
}

