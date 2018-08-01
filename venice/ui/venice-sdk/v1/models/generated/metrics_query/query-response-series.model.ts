/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */
import { Validators, FormControl, FormGroup, FormArray, ValidatorFn } from '@angular/forms';
import { minValueValidator, maxValueValidator, enumValidator } from './validators';
import { BaseModel, EnumDef } from './base-model';

import { QueryResponseRow, IQueryResponseRow } from './query-response-row.model';

export interface IQueryResponseSeries {
    'Columns'?: Array<string>;
    'Rows'?: Array<IQueryResponseRow>;
}


export class QueryResponseSeries extends BaseModel implements IQueryResponseSeries {
    'Columns': Array<string> = null;
    'Rows': Array<QueryResponseRow> = null;
    public static enumProperties: { [key: string] : EnumDef } = {
    }

    /**
     * Returns whether or not there is an enum property with a default value
    */
    public static hasDefaultEnumValue(prop) {
        return (QueryResponseSeries.enumProperties[prop] != null &&
                        QueryResponseSeries.enumProperties[prop].default != null &&
                        QueryResponseSeries.enumProperties[prop].default != '');
    }

    /**
     * constructor
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    constructor(values?: any) {
        super();
        this['Columns'] = new Array<string>();
        this['Rows'] = new Array<QueryResponseRow>();
        this.setValues(values);
    }

    /**
     * set the values. If a value isn't provided and we have a default, we use that.
     * @param values Can be used to set a webapi response to this newly constructed model
    */
    setValues(values: any): void {
        if (values) {
            this.fillModelArray<string>(this, 'Columns', values['Columns']);
        }
        if (values) {
            this.fillModelArray<QueryResponseRow>(this, 'Rows', values['Rows'], QueryResponseRow);
        }
    }




    protected getFormGroup(): FormGroup {
        if (!this._formGroup) {
            this._formGroup = new FormGroup({
                'Columns': new FormArray([]),
                'Rows': new FormArray([]),
            });
            // generate FormArray control elements
            this.fillFormArray<string>('Columns', this['Columns']);
            // generate FormArray control elements
            this.fillFormArray<QueryResponseRow>('Rows', this['Rows'], QueryResponseRow);
        }
        return this._formGroup;
    }

    setFormGroupValues() {
        if (this._formGroup) {
            this.fillModelArray<string>(this, 'Columns', this['Columns']);
            this.fillModelArray<QueryResponseRow>(this, 'Rows', this['Rows'], QueryResponseRow);
        }
    }
}

