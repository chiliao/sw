/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ClusterQuorumMemberStatus, IClusterQuorumMemberStatus } from './cluster-quorum-member-status.model';

export interface IClusterQuorumStatus {
    'members'?: Array<IClusterQuorumMemberStatus>;
}


export class ClusterQuorumStatus extends BaseModel implements IClusterQuorumStatus {
    'members': Array<ClusterQuorumMemberStatus> = null;
    public static propInfo: { [prop in keyof IClusterQuorumStatus]: PropInfoItem } = {
        'members': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return ClusterQuorumStatus.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return ClusterQuorumStatus.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (ClusterQuorumStatus.propInfo[prop] != null &&
                        ClusterQuorumStatus.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['members'] = new Array<ClusterQuorumMemberStatus>();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values) {
            this.fillModelArray<ClusterQuorumMemberStatus>(this, 'members', values['members'], ClusterQuorumMemberStatus);
        } else {
            this['members'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'members': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<ClusterQuorumMemberStatus>('members', this['members'], ClusterQuorumMemberStatus);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('members') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('members').get(field);
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
            this.fillModelArray<ClusterQuorumMemberStatus>(this, 'members', this['members'], ClusterQuorumMemberStatus);
        }
    }
}

