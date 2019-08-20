/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, minLengthValidator, maxLengthValidator, required, enumValidator, patternValidator, CustomFormControl, CustomFormGroup } from '../../../utils/validators';
import { BaseModel, PropInfoItem } from '../basemodel/base-model';

import { ApiObjectMeta, IApiObjectMeta } from './api-object-meta.model';
import { AuditEvent_stage,  } from './enums';
import { AuditEvent_level,  } from './enums';
import { ApiObjectRef, IApiObjectRef } from './api-object-ref.model';
import { AuditEvent_outcome,  } from './enums';

export interface IAuditEvent {
    'kind'?: string;
    'api-version'?: string;
    'meta'?: IApiObjectMeta;
    'stage'?: AuditEvent_stage;
    'level'?: AuditEvent_level;
    'user'?: IApiObjectRef;
    'client-ips'?: Array<string>;
    'resource'?: IApiObjectRef;
    'action'?: string;
    'outcome'?: AuditEvent_outcome;
    'request-uri'?: string;
    'request-object'?: string;
    'response-object'?: string;
    'gateway-node'?: string;
    'gateway-ip'?: string;
    'service-name'?: string;
    'data'?: object;
}


export class AuditEvent extends BaseModel implements IAuditEvent {
    'kind': string = null;
    'api-version': string = null;
    /** ObjectMeta.Name will be the UUID for an audit log object. */
    'meta': ApiObjectMeta = null;
    'stage': AuditEvent_stage = null;
    'level': AuditEvent_level = null;
    'user': ApiObjectRef = null;
    'client-ips': Array<string> = null;
    'resource': ApiObjectRef = null;
    'action': string = null;
    'outcome': AuditEvent_outcome = null;
    /** should be a valid URI */
    'request-uri': string = null;
    'request-object': string = null;
    'response-object': string = null;
    'gateway-node': string = null;
    'gateway-ip': string = null;
    'service-name': string = null;
    'data': object = null;
    public static propInfo: { [prop in keyof IAuditEvent]: PropInfoItem } = {
        'kind': {
            required: false,
            type: 'string'
        },
        'api-version': {
            required: false,
            type: 'string'
        },
        'meta': {
            description:  'ObjectMeta.Name will be the UUID for an audit log object.',
            required: false,
            type: 'object'
        },
        'stage': {
            enum: AuditEvent_stage,
            default: 'requestauthorization',
            required: false,
            type: 'string'
        },
        'level': {
            enum: AuditEvent_level,
            default: 'basic',
            required: false,
            type: 'string'
        },
        'user': {
            required: false,
            type: 'object'
        },
        'client-ips': {
            required: false,
            type: 'Array<string>'
        },
        'resource': {
            required: false,
            type: 'object'
        },
        'action': {
            required: false,
            type: 'string'
        },
        'outcome': {
            enum: AuditEvent_outcome,
            default: 'success',
            required: false,
            type: 'string'
        },
        'request-uri': {
            description:  'should be a valid URI',
            hint:  'https://10.1.1.1, ldap://10.1.1.1:800, /path/to/x',
            required: false,
            type: 'string'
        },
        'request-object': {
            required: false,
            type: 'string'
        },
        'response-object': {
            required: false,
            type: 'string'
        },
        'gateway-node': {
            required: false,
            type: 'string'
        },
        'gateway-ip': {
            required: false,
            type: 'string'
        },
        'service-name': {
            required: false,
            type: 'string'
        },
        'data': {
            required: false,
            type: 'object'
        },
    }

    public getPropInfo(propName: string): PropInfoItem {
        return AuditEvent.propInfo[propName];
    }

    public getPropInfoConfig(): { [key:string]:PropInfoItem } {
        return AuditEvent.propInfo;
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultValue(prop) {
        return (AuditEvent.propInfo[prop] != null &&
                        AuditEvent.propInfo[prop].default != null);
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any, setDefaults:boolean = true) {
        super();
        this['meta'] = new ApiObjectMeta();
        this['user'] = new ApiObjectRef();
        this['client-ips'] = new Array<string>();
        this['resource'] = new ApiObjectRef();
        this.setValues(values, setDefaults);
    }

    /**
     * set the values for both the Model and the Form Group. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any, fillDefaults = true): void {
        if (values && values['kind'] != null) {
            this['kind'] = values['kind'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('kind')) {
            this['kind'] = AuditEvent.propInfo['kind'].default;
        } else {
            this['kind'] = null
        }
        if (values && values['api-version'] != null) {
            this['api-version'] = values['api-version'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('api-version')) {
            this['api-version'] = AuditEvent.propInfo['api-version'].default;
        } else {
            this['api-version'] = null
        }
        if (values) {
            this['meta'].setValues(values['meta'], fillDefaults);
        } else {
            this['meta'].setValues(null, fillDefaults);
        }
        if (values && values['stage'] != null) {
            this['stage'] = values['stage'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('stage')) {
            this['stage'] = <AuditEvent_stage>  AuditEvent.propInfo['stage'].default;
        } else {
            this['stage'] = null
        }
        if (values && values['level'] != null) {
            this['level'] = values['level'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('level')) {
            this['level'] = <AuditEvent_level>  AuditEvent.propInfo['level'].default;
        } else {
            this['level'] = null
        }
        if (values) {
            this['user'].setValues(values['user'], fillDefaults);
        } else {
            this['user'].setValues(null, fillDefaults);
        }
        if (values && values['client-ips'] != null) {
            this['client-ips'] = values['client-ips'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('client-ips')) {
            this['client-ips'] = [ AuditEvent.propInfo['client-ips'].default];
        } else {
            this['client-ips'] = [];
        }
        if (values) {
            this['resource'].setValues(values['resource'], fillDefaults);
        } else {
            this['resource'].setValues(null, fillDefaults);
        }
        if (values && values['action'] != null) {
            this['action'] = values['action'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('action')) {
            this['action'] = AuditEvent.propInfo['action'].default;
        } else {
            this['action'] = null
        }
        if (values && values['outcome'] != null) {
            this['outcome'] = values['outcome'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('outcome')) {
            this['outcome'] = <AuditEvent_outcome>  AuditEvent.propInfo['outcome'].default;
        } else {
            this['outcome'] = null
        }
        if (values && values['request-uri'] != null) {
            this['request-uri'] = values['request-uri'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('request-uri')) {
            this['request-uri'] = AuditEvent.propInfo['request-uri'].default;
        } else {
            this['request-uri'] = null
        }
        if (values && values['request-object'] != null) {
            this['request-object'] = values['request-object'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('request-object')) {
            this['request-object'] = AuditEvent.propInfo['request-object'].default;
        } else {
            this['request-object'] = null
        }
        if (values && values['response-object'] != null) {
            this['response-object'] = values['response-object'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('response-object')) {
            this['response-object'] = AuditEvent.propInfo['response-object'].default;
        } else {
            this['response-object'] = null
        }
        if (values && values['gateway-node'] != null) {
            this['gateway-node'] = values['gateway-node'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('gateway-node')) {
            this['gateway-node'] = AuditEvent.propInfo['gateway-node'].default;
        } else {
            this['gateway-node'] = null
        }
        if (values && values['gateway-ip'] != null) {
            this['gateway-ip'] = values['gateway-ip'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('gateway-ip')) {
            this['gateway-ip'] = AuditEvent.propInfo['gateway-ip'].default;
        } else {
            this['gateway-ip'] = null
        }
        if (values && values['service-name'] != null) {
            this['service-name'] = values['service-name'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('service-name')) {
            this['service-name'] = AuditEvent.propInfo['service-name'].default;
        } else {
            this['service-name'] = null
        }
        if (values && values['data'] != null) {
            this['data'] = values['data'];
        } else if (fillDefaults && AuditEvent.hasDefaultValue('data')) {
            this['data'] = AuditEvent.propInfo['data'].default;
        } else {
            this['data'] = null
        }
        this.setFormGroupValuesToBeModelValues();
    }


    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'kind': CustomFormControl(new FormControl(this['kind']), AuditEvent.propInfo['kind']),
                'api-version': CustomFormControl(new FormControl(this['api-version']), AuditEvent.propInfo['api-version']),
                'meta': CustomFormGroup(this['meta'].$formGroup, AuditEvent.propInfo['meta'].required),
                'stage': CustomFormControl(new FormControl(this['stage'], [enumValidator(AuditEvent_stage), ]), AuditEvent.propInfo['stage']),
                'level': CustomFormControl(new FormControl(this['level'], [enumValidator(AuditEvent_level), ]), AuditEvent.propInfo['level']),
                'user': CustomFormGroup(this['user'].$formGroup, AuditEvent.propInfo['user'].required),
                'client-ips': CustomFormControl(new FormControl(this['client-ips']), AuditEvent.propInfo['client-ips']),
                'resource': CustomFormGroup(this['resource'].$formGroup, AuditEvent.propInfo['resource'].required),
                'action': CustomFormControl(new FormControl(this['action']), AuditEvent.propInfo['action']),
                'outcome': CustomFormControl(new FormControl(this['outcome'], [enumValidator(AuditEvent_outcome), ]), AuditEvent.propInfo['outcome']),
                'request-uri': CustomFormControl(new FormControl(this['request-uri']), AuditEvent.propInfo['request-uri']),
                'request-object': CustomFormControl(new FormControl(this['request-object']), AuditEvent.propInfo['request-object']),
                'response-object': CustomFormControl(new FormControl(this['response-object']), AuditEvent.propInfo['response-object']),
                'gateway-node': CustomFormControl(new FormControl(this['gateway-node']), AuditEvent.propInfo['gateway-node']),
                'gateway-ip': CustomFormControl(new FormControl(this['gateway-ip']), AuditEvent.propInfo['gateway-ip']),
                'service-name': CustomFormControl(new FormControl(this['service-name']), AuditEvent.propInfo['service-name']),
                'data': CustomFormControl(new FormControl(this['data']), AuditEvent.propInfo['data']),
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('meta') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('meta').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('user') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('user').get(field);
                control.updateValueAndValidity();
            });
            // We force recalculation of controls under a form group
            Object.keys((this._formGroup.get('resource') as FormGroup).controls).forEach(field => {
                const control = this._formGroup.get('resource').get(field);
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
            this['meta'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['stage'].setValue(this['stage']);
            this._formGroup.controls['level'].setValue(this['level']);
            this['user'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['client-ips'].setValue(this['client-ips']);
            this['resource'].setFormGroupValuesToBeModelValues();
            this._formGroup.controls['action'].setValue(this['action']);
            this._formGroup.controls['outcome'].setValue(this['outcome']);
            this._formGroup.controls['request-uri'].setValue(this['request-uri']);
            this._formGroup.controls['request-object'].setValue(this['request-object']);
            this._formGroup.controls['response-object'].setValue(this['response-object']);
            this._formGroup.controls['gateway-node'].setValue(this['gateway-node']);
            this._formGroup.controls['gateway-ip'].setValue(this['gateway-ip']);
            this._formGroup.controls['service-name'].setValue(this['service-name']);
            this._formGroup.controls['data'].setValue(this['data']);
        }
    }
}

