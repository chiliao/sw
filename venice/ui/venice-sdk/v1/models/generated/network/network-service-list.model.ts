/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiListMeta, IApiListMeta } from './api-list-meta.model';
import { NetworkService, INetworkService } from './network-service.model';

export interface INetworkServiceList {
    'kind'?: string;
    'api-version'?: string;
    'list-meta'?: IApiListMeta;
    'items'?: Array<INetworkService>;
    '_ui'?: any;
}


export class NetworkServiceList extends BaseModel implements INetworkServiceList {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'kind': string = null;
    'api-version': string = null;
    'list-meta': ApiListMeta = null;
    /** List of Service objects. */
    'items': Array<NetworkService> = null;
    public static propInfo: { [prop in keyof INetworkServiceList]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'list-meta': {
            required: false,
            type: 'object'
        },
        'items': {
            description:  `List of Service objects.`,
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkServiceList.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkServiceList.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkServiceList.propInfo[prop] != null &&
                        NetworkServiceList.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['list-meta'] = new ApiListMeta();
        this['items'] = new Array<NetworkService>();
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
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && NetworkServiceList.hasDefaultValue('kind')) {
            this['kind'] = NetworkServiceList.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && NetworkServiceList.hasDefaultValue('api-version')) {
            this['api-version'] = NetworkServiceList.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['list-meta'].setValues(values['list-meta'], fillDefaults);
        } else {
            this['list-meta'].setValues(null, fillDefaults);
        }
        if (values) {
            this.fillModelArray<NetworkService>(this, 'items', values['items'], NetworkService);
        } else {
            this['items'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), NetworkServiceList.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), NetworkServiceList.propInfo['api-version']),
                'list-meta': CustomFormGroup(this['list-meta'].$formGroup, NetworkServiceList.propInfo['list-meta'].required),
                'items': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<NetworkService>('items', this['items'], NetworkService);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('list-meta') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('list-meta').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('items') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('items').get(field);
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
            this['list-meta'].setFormGroupValuesToBeModelValues();
            this.fillModelArray<NetworkService>(this, 'items', this['items'], NetworkService);
        }
    }
}

