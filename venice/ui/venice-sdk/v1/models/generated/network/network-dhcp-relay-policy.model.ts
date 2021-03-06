/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { NetworkDHCPServer, INetworkDHCPServer } from './network-dhcp-server.model';

export interface INetworkDHCPRelayPolicy {
    'relay-servers'?: Array<INetworkDHCPServer>;
    '_ui'?: any;
}


export class NetworkDHCPRelayPolicy extends BaseModel implements INetworkDHCPRelayPolicy {
    /** Field for holding arbitrary ui state */
    '_ui': any = {};
    'relay-servers': Array<NetworkDHCPServer> = null;
    public static propInfo: { [prop in keyof INetworkDHCPRelayPolicy]: PropInfoItem } = {
        'relay-servers': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return NetworkDHCPRelayPolicy.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return NetworkDHCPRelayPolicy.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (NetworkDHCPRelayPolicy.propInfo[prop] != null &&
                        NetworkDHCPRelayPolicy.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['relay-servers'] = new Array<NetworkDHCPServer>();
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
        if (values) {
            this.fillModelArray<NetworkDHCPServer>(this, 'relay-servers', values['relay-servers'], NetworkDHCPServer);
        } else {
            this['relay-servers'] = [];
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'relay-servers': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<NetworkDHCPServer>('relay-servers', this['relay-servers'], NetworkDHCPServer);
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('relay-servers') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('relay-servers').get(field);
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
            this.fillModelArray<NetworkDHCPServer>(this, 'relay-servers', this['relay-servers'], NetworkDHCPServer);
        }
    }
}

