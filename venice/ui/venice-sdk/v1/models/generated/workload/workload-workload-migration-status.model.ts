/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { WorkloadWorkloadMigrationStatus_stage,  } from './enums';
import { WorkloadWorkloadMigrationStatus_status,  } from './enums';

export interface IWorkloadWorkloadMigrationStatus {
    'stage': WorkloadWorkloadMigrationStatus_stage;
    'started-at'?: Date;
    'status': WorkloadWorkloadMigrationStatus_status;
    'completed-at'?: Date;
    '_ui'?: any;
}


export class WorkloadWorkloadMigrationStatus extends BaseModel implements IWorkloadWorkloadMigrationStatus {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Controller's migration stage. */
    'stage': WorkloadWorkloadMigrationStatus_stage = null;
    /** Migration start time. */
    'started-at': Date = null;
    /** The status from the dataplane performing migration. */
    'status': WorkloadWorkloadMigrationStatus_status = null;
    /** Migration completion time. */
    'completed-at': Date = null;
    public static propInfo: { [prop in keyof IWorkloadWorkloadMigrationStatus]: PropInfoItem } = {
        'stage': {
            enum: WorkloadWorkloadMigrationStatus_stage,
            default: 'migration-none',
            description:  `Controller's migration stage.`,
            required: true,
            type: 'string'
        },
        'started-at': {
            description:  `Migration start time.`,
            required: false,
            type: 'Date'
        },
        'status': {
            enum: WorkloadWorkloadMigrationStatus_status,
            default: 'none',
            description:  `The status from the dataplane performing migration.`,
            required: true,
            type: 'string'
        },
        'completed-at': {
            description:  `Migration completion time.`,
            required: false,
            type: 'Date'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return WorkloadWorkloadMigrationStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return WorkloadWorkloadMigrationStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (WorkloadWorkloadMigrationStatus.propInfo[prop] != null &&
                        WorkloadWorkloadMigrationStatus.propInfo[prop].default != null);
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
        if (values && values['stage'] != null) {
            this['stage'] = values['stage'];
        } else if (fillDefaults && WorkloadWorkloadMigrationStatus.hasDefaultValue('stage')) {
            this['stage'] = <WorkloadWorkloadMigrationStatus_stage>  WorkloadWorkloadMigrationStatus.propInfo['stage'].default;
        } else {
            this['stage'] = null
        }
        if (values && values['started-at'] != null) {
            this['started-at'] = values['started-at'];
        } else if (fillDefaults && WorkloadWorkloadMigrationStatus.hasDefaultValue('started-at')) {
            this['started-at'] = WorkloadWorkloadMigrationStatus.propInfo['started-at'].default;
        } else {
            this['started-at'] = null
        }
        if (values && values['status'] != null) {
            this['status'] = values['status'];
        } else if (fillDefaults && WorkloadWorkloadMigrationStatus.hasDefaultValue('status')) {
            this['status'] = <WorkloadWorkloadMigrationStatus_status>  WorkloadWorkloadMigrationStatus.propInfo['status'].default;
        } else {
            this['status'] = null
        }
        if (values && values['completed-at'] != null) {
            this['completed-at'] = values['completed-at'];
        } else if (fillDefaults && WorkloadWorkloadMigrationStatus.hasDefaultValue('completed-at')) {
            this['completed-at'] = WorkloadWorkloadMigrationStatus.propInfo['completed-at'].default;
        } else {
            this['completed-at'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'stage': CustomFormControl(new FormControl(this['stage'], [required, enumValidator(WorkloadWorkloadMigrationStatus_stage), ]), WorkloadWorkloadMigrationStatus.propInfo['stage']),
                'started-at': CustomFormControl(new FormControl(this['started-at']), WorkloadWorkloadMigrationStatus.propInfo['started-at']),
                'status': CustomFormControl(new FormControl(this['status'], [required, enumValidator(WorkloadWorkloadMigrationStatus_status), ]), WorkloadWorkloadMigrationStatus.propInfo['status']),
                'completed-at': CustomFormControl(new FormControl(this['completed-at']), WorkloadWorkloadMigrationStatus.propInfo['completed-at']),
            });
        }
        return this._formGroup;
    }

    setModelToBeFormGroupValues() {
        this.setValues(this.$formGroup.value, false);
    }

    setFormGroupValuesToBeModelValues() {
        if (this._formGroup) {
            this._formGroup.controls['stage'].setValue(this['stage']);
            this._formGroup.controls['started-at'].setValue(this['started-at']);
            this._formGroup.controls['status'].setValue(this['status']);
            this._formGroup.controls['completed-at'].setValue(this['completed-at']);
        }
    }
}

