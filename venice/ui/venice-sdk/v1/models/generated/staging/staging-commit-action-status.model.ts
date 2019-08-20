/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { StagingCommitActionStatus_status,  } from './enums';

export interface IStagingCommitActionStatus {
    'status': StagingCommitActionStatus_status;
    'reason'?: string;
}


export class StagingCommitActionStatus extends BaseModel implements IStagingCommitActionStatus {
    'status': StagingCommitActionStatus_status = null;
    'reason': string = null;
    public static propInfo: { [prop in keyof IStagingCommitActionStatus]: PropInfoItem } = {
        'status': {
            enum: StagingCommitActionStatus_status,
            default: 'success',
            required: true,
            type: 'string'
        },
        'reason': {
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return StagingCommitActionStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return StagingCommitActionStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (StagingCommitActionStatus.propInfo[prop] != null &&
                        StagingCommitActionStatus.propInfo[prop].default != null);
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
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && StagingCommitActionStatus.hasDefaultValue('status')) {
            this['status'] = <StagingCommitActionStatus_status>  StagingCommitActionStatus.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['reason'] != null) {
            this['reason'] = values['reason'];
        } else if (fillDefaults && StagingCommitActionStatus.hasDefaultValue('reason')) {
            this['reason'] = StagingCommitActionStatus.propInfo['reason'].default;
        } else {
            this['reason'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'status': CustomFormControl(new FormControl(this['status'], [required, enumValidator(StagingCommitActionStatus_status), ]), StagingCommitActionStatus.propInfo['status']),
                'reason': CustomFormControl(new FormControl(this['reason']), StagingCommitActionStatus.propInfo['reason']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['reason'].setValue(this['reason']);
        }
    }
}

