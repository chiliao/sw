/**
 * This file is generated by the SwaggerTSGenerator.
 * Do not edit.
*/
/* tslint:disable */

// generate enum based on strings instead of numbers
// (see https://blog.rsuter.com/how-to-implement-an-enum-with-string-values-in-typescript/)
export enum LabelsRequirement_operator {
    'equals' = "equals",
    'notEquals' = "notEquals",
    'in' = "in",
    'notIn' = "notIn",
}

export enum Metrics_queryQuerySpec_function {
    'NONE' = "NONE",
    'MEAN' = "MEAN",
    'MAX' = "MAX",
}


export enum LabelsRequirement_operator_uihint {
    'notEquals' = "not equals",
    'notIn' = "not in",
}




/**
 * bundle of all enums for databinding to options, radio-buttons etc.
 * usage in component:
 *   import { AllEnums, minValueValidator, maxValueValidator } from '../../models/webapi';
 *
 *   @Component({
 *       ...
 *   })
 *   export class xxxComponent implements OnInit {
 *       allEnums = AllEnums;
 *       ...
 *       ngOnInit() {
 *           this.allEnums = AllEnums.instance;
 *       }
 *   }
*/
export class AllEnums {
    private static _instance: AllEnums = new AllEnums();
    constructor() {
        if (AllEnums._instance) {
            throw new Error("Error: Instantiation failed: Use AllEnums.instance instead of new");
        }
        AllEnums._instance = this;
    }
    static get instance(): AllEnums {
        return AllEnums._instance;
    }

    LabelsRequirement_operator = LabelsRequirement_operator;
    Metrics_queryQuerySpec_function = Metrics_queryQuerySpec_function;

    LabelsRequirement_operator_uihint = LabelsRequirement_operator_uihint;
}
