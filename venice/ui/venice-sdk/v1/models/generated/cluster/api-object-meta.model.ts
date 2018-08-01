/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';


export interface IApiObjectMeta {
    'name'?: string;
    'tenant'?: string;
    'namespace'?: string;
    'resource-version'?: string;
    'uuid'?: string;
    'labels'?: object;
    'creation-time'?: Date;
    'mod-time'?: Date;
    'self-link'?: string;
}


export class ApiObjectMeta extends BaseModel implements IApiObjectMeta {
    /** Name of the object, unique within a Namespace for scoped objects. */
    'name': string = null;
    /** Tenant is global namespace isolation for various objects. This can be automatically
filled in many cases based on the tenant a user, who created the object, belongs go. */
    'tenant': string = null;
    /** Namespace of the object, for scoped objects. */
    'namespace': string = null;
    /** Resource version in the object store. This can only be set by the server. */
    'resource-version': string = null;
    /** UUID is the unique identifier for the object. This can only be set by the server. */
    'uuid': string = null;
    /** Labels are arbitrary (key,value) pairs associated with any object. */
    'labels': object = null;
    'creation-time': Date = null;
    'mod-time': Date = null;
    /** SelfLink is a link to accessing this object. When stored in the KV store this is
 the key in the kvstore and when the object is served from the API-GW it is the
 URI path. Examples
   - "/venice/tenants/tenants/tenant2" in the kvstore
   - "/v1/tenants/tenants/tenant2" when served by API Gateway. */
    'self-link': string = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (ApiObjectMeta.enumProperties[prop] != null &&
                        ApiObjectMeta.enumProperties[prop].default != null &&
                        ApiObjectMeta.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values && values['name'] != null) {
            this['name'] = values['name'];
        }
        if (values && values['tenant'] != null) {
            this['tenant'] = values['tenant'];
        }
        if (values && values['namespace'] != null) {
            this['namespace'] = values['namespace'];
        }
        if (values && values['resource-version'] != null) {
            this['resource-version'] = values['resource-version'];
        }
        if (values && values['uuid'] != null) {
            this['uuid'] = values['uuid'];
        }
        if (values && values['labels'] != null) {
            this['labels'] = values['labels'];
        }
        if (values && values['creation-time'] != null) {
            this['creation-time'] = values['creation-time'];
        }
        if (values && values['mod-time'] != null) {
            this['mod-time'] = values['mod-time'];
        }
        if (values && values['self-link'] != null) {
            this['self-link'] = values['self-link'];
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'name': new FormControl(this['name']),
                'tenant': new FormControl(this['tenant']),
                'namespace': new FormControl(this['namespace']),
                'resource-version': new FormControl(this['resource-version']),
                'uuid': new FormControl(this['uuid']),
                'labels': new FormControl(this['labels']),
                'creation-time': new FormControl(this['creation-time']),
                'mod-time': new FormControl(this['mod-time']),
                'self-link': new FormControl(this['self-link']),
            });
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this._formGroup.controls['name'].setValue(this['name']);
            this._formGroup.controls['tenant'].setValue(this['tenant']);
            this._formGroup.controls['namespace'].setValue(this['namespace']);
            this._formGroup.controls['resource-version'].setValue(this['resource-version']);
            this._formGroup.controls['uuid'].setValue(this['uuid']);
            this._formGroup.controls['labels'].setValue(this['labels']);
            this._formGroup.controls['creation-time'].setValue(this['creation-time']);
            this._formGroup.controls['mod-time'].setValue(this['mod-time']);
            this._formGroup.controls['self-link'].setValue(this['self-link']);
        }
    }
}

