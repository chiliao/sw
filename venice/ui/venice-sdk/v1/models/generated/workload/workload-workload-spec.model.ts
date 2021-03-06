/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { WorkloadWorkloadIntfSpec, IWorkloadWorkloadIntfSpec } from './workload-workload-intf-spec.model';

export interface IWorkloadWorkloadSpec {
    'host-name': string;
    'interfaces'?: Array<IWorkloadWorkloadIntfSpec>;
    'migration-timeout'?: string;
    '_ui'?: any;
}


export class WorkloadWorkloadSpec extends BaseModel implements IWorkloadWorkloadSpec {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    /** Hostname of the server where the workload should be running. Should be a valid host address, IP address or hostname. */
    'host-name': string = null;
    /** Spec of all interfaces in the Workload identified by Primary MAC. */
    'interfaces': Array<WorkloadWorkloadIntfSpec> = null;
    /** Should be a valid time duration. */
    'migration-timeout': string = null;
    public static propInfo: { [prop in keyof IWorkloadWorkloadSpec]: PropInfoItem } = {
        'host-name': {
            description:  `Hostname of the server where the workload should be running. Should be a valid host address, IP address or hostname.`,
            hint:  '10.1.1.1, ff02::5, localhost, example.domain.com ',
            required: true,
            type: 'string'
        },
        'interfaces': {
            description:  `Spec of all interfaces in the Workload identified by Primary MAC.`,
            required: false,
            type: 'object'
        },
        'migration-timeout': {
            default: '3m',
            description:  `Should be a valid time duration.`,
            hint:  '2h',
            required: false,
            type: 'string'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return WorkloadWorkloadSpec.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return WorkloadWorkloadSpec.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (WorkloadWorkloadSpec.propInfo[prop] != null &&
                        WorkloadWorkloadSpec.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['interfaces'] = new Array<WorkloadWorkloadIntfSpec>();
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
        if (values && values['host-name'] != null) {
            this['host-name'] = values['host-name'];
        } else if (fillDefaults && WorkloadWorkloadSpec.hasDefaultValue('host-name')) {
            this['host-name'] = WorkloadWorkloadSpec.propInfo['host-name'].default;
        } else {
            this['host-name'] = null
        }
        if (values) {
            this.fillModelArray<WorkloadWorkloadIntfSpec>(this, 'interfaces', values['interfaces'], WorkloadWorkloadIntfSpec);
        } else {
            this['interfaces'] = [];
        }
        if (values && values['migration-timeout'] != null) {
            this['migration-timeout'] = values['migration-timeout'];
        } else if (fillDefaults && WorkloadWorkloadSpec.hasDefaultValue('migration-timeout')) {
            this['migration-timeout'] = WorkloadWorkloadSpec.propInfo['migration-timeout'].default;
        } else {
            this['migration-timeout'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'host-name': CustomFormControl(new FormControl(this['host-name'], [required, ]), WorkloadWorkloadSpec.propInfo['host-name']),
                'interfaces': new FormArray([]),
                'migration-timeout': CustomFormControl(new FormControl(this['migration-timeout']), WorkloadWorkloadSpec.propInfo['migration-timeout']),
            });
            // generate FormArray control elements
            this.fillFormArray<WorkloadWorkloadIntfSpec>('interfaces', this['interfaces'], WorkloadWorkloadIntfSpec);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('interfaces') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('interfaces').get(field);
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
            this._formGroup.controls['host-name'].setValue(this['host-name']);
            this.fillModelArray<WorkloadWorkloadIntfSpec>(this, 'interfaces', this['interfaces'], WorkloadWorkloadIntfSpec);
            this._formGroup.controls['migration-timeout'].setValue(this['migration-timeout']);
        }
    }
}

