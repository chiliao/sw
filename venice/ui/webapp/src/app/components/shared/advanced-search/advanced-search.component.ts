import {
  Component,
  EventEmitter,
  Input,
  OnInit,
  Output,
  ViewChild,
  ViewEncapsulation
} from '@angular/core';
import { RepeaterComponent, RepeaterData, RepeaterItem, ValueType } from 'web-app-framework';
import { FormArray, FormControl } from '@angular/forms';
import { SearchUtil } from '@components/search/SearchUtil';
import { Animations } from '@app/animations';
import { TableCol } from '@components/shared/tableviewedit';
import * as _ from 'lodash';
import { SearchSearchRequest, SearchSearchRequest_sort_order, FieldsRequirement, IFieldsRequirement } from '@sdk/v1/models/generated/search';
import { ControllerService } from '@app/services/controller.service';
import { Utility } from '@common/Utility';
import * as moment from 'moment';


/**
 * Advanced Search Component
 *
 * Features:
 * 1. Text search mode
 * 2. Advanced panel ui field selector mode
 * 3. Text mode and field selector mode is interchangeable.
 * 4. Integrate search and cancel inside. The search and cancel could also be delegated to the parent component for handling requests.
 * 5. Support general search and field search parsing and compiling
 *
 * Example usage: naplesdetail.component.html
 *
 * Params:
 * 1. repeaterValues: repeaterValues handler.
 * 2. formArray: pass in value to repeater's form data.
 * 3. cols and kind: this is for converting tableColsToRepeaterData dynamically.
 * 4. cancelEmitter and searchEmitter: cancel and search handler.
 *
 * TODO:
 * Handle form validation by using customValueOnBlur.
 */

export interface LocalSearchRequest {
  query: Array<IFieldsRequirement | FieldsRequirement>;
  sortBy: string;
  sortOrder: SearchSearchRequest_sort_order;
}

export interface LocalSearchResult {
  searchRes: Array<string>;
  err: boolean;
}

@Component({
  selector: 'app-advanced-search',
  templateUrl: './advanced-search.component.html',
  styleUrls: ['./advanced-search.component.scss'],
  animations: [Animations],
  encapsulation: ViewEncapsulation.None,
})

export class AdvancedSearchComponent implements OnInit {
  @ViewChild('fieldRepeater') fieldRepeater: RepeaterComponent;
  @Input() formArray = new FormArray([]);
  @Input() keyFormName: string = 'keyFormControl';
  @Input() operatorFormName: string = 'operatorFormControl';
  @Input() valueFormName: string = 'valueFormControl';
  @Input() keytextFormName: string = 'keytextFormName';
  @Input() cols: TableCol[] = [];
  @Input() kind: string;
  @Input() customQueryOptions: RepeaterData[] = [];
  @Input() maxSearchRecords: number = 4000;

  @Output() repeaterValues: EventEmitter<any> = new EventEmitter();
  @Output() searchEmitter: EventEmitter<any> = new EventEmitter();
  @Output() cancelEmitter: EventEmitter<any> = new EventEmitter();

  localSearchFields: { [key: string]: boolean } = {};
  showAdvancedPanel: boolean = false;
  fieldData: RepeaterData[] = [];
  search: string = '';
  generalSearch: string = '';
  valueLabelToValueMap = {};
  buildFieldValuePlaceholder = SearchUtil.buildFieldValuePlaceholder;
  constructor(private controlerService: ControllerService) {
  }

  customValueOnBlur = ($event: any, repeaterItem: RepeaterItem) => {
    const instance = SearchUtil.getModelInfo(Utility.findCategoryByKind(this.kind), this.kind);
    const key = this.valueLabelToValueMap[repeaterItem.formGroup.value[this.keyFormName]];
    const value = repeaterItem.formGroup.value[this.valueFormName];
    const type = (Utility.getNestedPropInfo(instance, key)) ? Utility.getNestedPropInfo(instance, key).type : null;
    switch (type) {
      case 'Date':
        // TODO: validator logic goes here for Date
        break;
      default:
    }

    const classList: DOMTokenList = $event.target.classList;
    if (classList.contains('repeater-input-value')) {
      // TODO: some how add error class if the input is invalid
    }
  }

  customValueOnKeydown = ($event) => {
    if ($event.key === 'Enter') {
      this.searchClicked();
    }
  }

  ngOnInit() {
    this.fieldData = this.generateFieldData(this.customQueryOptions);
    this.genValueLabelToFieldMap();
  }

  /**
   * This function generates field Data
   * It applies any custom options given for this field
   * @param queryArray
   */
  generateFieldData(queryArray: RepeaterData[]): RepeaterData[] {
    const { repeaterData, localFields } = SearchUtil.tableColsToRepeaterData(this.cols, this.kind);
    if (queryArray) {
      queryArray.forEach(obj => {
        for (let i = 0; i < repeaterData.length; i++) {
          if (repeaterData[i].key.value === obj.key.value) {
            repeaterData[i] = obj;
            if (obj.valueType === ValueType.singleSelect && (obj.key.value in localFields)) {
              localFields[obj.key.value].singleSelect = true;
            }
            return;
          }
        }
      });
      Object.values(localFields).forEach(f => {
        this.localSearchFields[f.field] = f.singleSelect;
      });
    }
    return repeaterData;
  }

  /**
   * Table col Header to field mapping helper
   * This is for mapping user readable label to backend friendly label.
   * Example:
   * 'who' => 'user.name'
   */
  genValueLabelToFieldMap() {
    this.cols.forEach(ele => {
      this.valueLabelToValueMap[ele.header] = ele.field;
    });
  }

  setDefaultData() {
    this.fieldData = [
      {
        key: { label: 'name', value: 'name' },
        operators: SearchUtil.stringOperators,
        valueType: ValueType.inputField
      },
      {
        key: { label: 'tenant', value: 'tenant' },
        operators: SearchUtil.stringOperators,
        valueType: ValueType.inputField
      },
      {
        key: { label: 'namespace', value: 'namespace' },
        operators: SearchUtil.stringOperators,
        valueType: ValueType.inputField
      },
      {
        key: { label: 'creation-time', value: 'creation-time' },
        operators: SearchUtil.numberOperators,
        valueType: ValueType.inputField
      },
      {
        key: { label: 'modified-time', value: 'mod-time' },
        valueType: ValueType.inputField,
        operators: SearchUtil.numberOperators
      }
    ];
  }

  handleRepeaterValues(values) {
    // pass modal in panel to search input as string
    this.repeaterValues.emit(this.getValues());
  }

  /**
   * Expose raw modal data to the parent component from here.
   */
  getValues() {
    if (!this.showAdvancedPanel) {
      // in text search mode
      return this.getValueFromSearchTextHelper();
    }
    return this.fieldRepeater ? this.getValueHelper() : [];
  }

  /**
   * Human readable text values from getValues in fieldRepeater.
   * The result is user readable.
   * Only used in this component.
   */
  getUserFriendlyValues() {
    return this.fieldRepeater ? Utility.formatRepeaterData(this.fieldRepeater.getValues(), this.valueFormName) : [];
  }

  /**
   * A helper for getValue with the backend friendly keyFormControl from valueLabelToValueMap.
   * Example: exposing 'meta.name' rather than 'Who'
   */
  getValueHelper() {
    const temp = _.cloneDeep(Utility.formatRepeaterData(this.fieldRepeater.getValues(), this.valueFormName));
    temp.forEach((ele, i) => {
      if (ele) {
        temp[i].keyFormControl = this.valueLabelToValueMap.hasOwnProperty(ele.keyFormControl) ? this.valueLabelToValueMap[ele.keyFormControl] : '';
      }
    });
    return temp;
  }

  /**
   * A helper for getValue with the backend friendly keyFormControl from valueLabelToValueMap.
   * This also handle parsing logic from text input in search bar to actual modal.
   */
  getValueFromSearchTextHelper() {
    this.searchTextToFormArrayModal();
    const temp = [];
    this.formArray.controls.forEach(ele => temp.push(ele.value));
    temp.forEach((ele, i) => {
      if (ele) {
        temp[i].keyFormControl = this.valueLabelToValueMap.hasOwnProperty(ele.keyFormControl) ? this.valueLabelToValueMap[ele.keyFormControl] : '';
      }
    });
    return Utility.formatRepeaterData(temp, this.valueFormName);
  }

  /**
   * Toggle panel action handler. Handling the show/hide panel.
   */
  togglePanel() {
    if (this.showAdvancedPanel) {
      // pass modal in panel to search input as string
      this.search = SearchUtil.advancedSearchCompiler(this.getUserFriendlyValues(), this.generalSearch);
    } else {
      // pass search input string to modal.
      this.searchTextToFormArrayModal();
    }
    this.showAdvancedPanel = !this.showAdvancedPanel;
  }

  /**
   * Convert this.search string to FromArray modal which will be used in repeater.
   */
  searchTextToFormArrayModal() {
    if (this.search === null || this.search === '') {
      this.generalSearch = '';
      this.formArray = new FormArray([]);  // clear formArray
      return;
    }
    let searchExpressions = [], generalSearch = [];

    try {
      const a = SearchUtil.advancedSearchParser(this.search);
      searchExpressions = a.searchExpressions;
      generalSearch = a.generalSearch;

      if (searchExpressions === null || searchExpressions.length === 0) {
        this.formArray = new FormArray([]);  // clear formArray
      }
      if (searchExpressions !== null && searchExpressions.length !== 0) {
        this.formArray = new FormArray([]);  // clear formArray
        searchExpressions.forEach(ele => {
          this.formArray.push(new FormControl({
            keyFormControl: ele.key,
            operatorFormControl: ele.operator,
            valueFormControl: ele.values
          }));
        });
      }

    } catch (error) {
      generalSearch = [this.search];
    }

    if (generalSearch) {
      this.generalSearch = generalSearch.join(' ');
    }
  }

  /**
   * Format any valid string date input to yyyy-mm-dd
   * if the date is invalid, it will return the original input.
   * @param date
   */
  formatDate(date: string): string {
    const timestamp = moment(date, 'yyyy-mm-dd');

    if (timestamp.isValid() !== false) {
      const d = new Date(date),
        year = d.getFullYear();
      let month = '' + (d.getMonth() + 1),
        day = '' + d.getDate();

      if (month.length < 2) {
        month = '0' + month;
      }
      if (day.length < 2) {
        day = '0' + day;
      }

      return [year, month, day].join('-');
    } else {
      const d = Date.parse(date);
      const myDate = new Date(d);
      return myDate.toISOString();
    }

  }

  /**
   * This is the request builder function for helping parent component sending remote request
   * @param field
   * @param order
   * @param kind
   * @param maxRecords
   * @returns {SearchSearchRequest}
   */
  getSearchRequest(field, order, kind, aggregate = true, maxRecords = this.maxSearchRecords): SearchSearchRequest {
    let sortOrder = SearchSearchRequest_sort_order.ascending;
    if (order === -1) {
      sortOrder = SearchSearchRequest_sort_order.descending;
    }

    let searchSearchRequest: SearchSearchRequest = new SearchSearchRequest(null, false);
    const model = this.getValues();
    const texts = SearchUtil.splitString(this.generalSearch);
    if (model !== null && model.length !== 0) {
      const fields = [];

      const instance = SearchUtil.getModelInfo(Utility.findCategoryByKind(this.kind), this.kind);
      model.forEach(ele => {
        // all value post process logic goes here
        let processedValue;
        if (field !== '') {
          const type = Utility.getNestedPropInfo(instance, ele.keyFormControl) ? Utility.getNestedPropInfo(instance, ele.keyFormControl).type : '';
          if (type === 'Date') {
            processedValue = ele.valueFormControl.map(e => this.formatDate(e));
          } else {
            processedValue = ele.valueFormControl.map(e => e);
          }
        } else {
          processedValue = [];
        }
        if (!(ele.keyFormControl in this.localSearchFields)) {
          fields.push({
            key: ele.keyFormControl, // this.buildSearchKindFieldKey(ele) ,  // VS-774 ele.keyFormControl,
            operator: ele.operatorFormControl,
            values: processedValue
          });
        }
      });
      const payload = {
        'query': {
          'fields': { 'requirements': fields },
          'labels': { 'requirements': [] },
          'texts': [
            {
              'text': texts
            }
          ]
        }
      };
      searchSearchRequest = new SearchSearchRequest(payload, false);  // we don't to fill default values. So set the second parameter as false;
    } else {
      // model is empty, we only have general texts
      const payload = {
        'query': {
          'texts': [
            {
              'text': texts
            }
          ]
        }
      };
      searchSearchRequest = new SearchSearchRequest(payload, false);  // we don't to fill default values. So set the second parameter as false;
    }
    searchSearchRequest.query.kinds = [kind];
    searchSearchRequest['sort-by'] = field;
    searchSearchRequest['sort-order'] = sortOrder;
    searchSearchRequest.from = 0;
    searchSearchRequest['max-results'] = maxRecords;
    searchSearchRequest['aggregate'] = aggregate;
    return searchSearchRequest;
  }

  /**
   * This is the request builder function for generating local request query
   * @param field
   * @param order
   * @returns {LocalSearchResult}
   */
  getLocalSearchResult(field, order, searchObject): LocalSearchResult {
    const localSearchRequest: LocalSearchRequest = this.getLocalSearchRequest(order, field);
    let localSearchResult: LocalSearchResult = {
      searchRes: null,
      err: false
    };
    if (localSearchRequest.query != null && localSearchRequest.query.length > 0) {
      localSearchResult = this.localSearch(localSearchRequest, searchObject);
    }
    return localSearchResult;
  }

  getLocalSearchRequest(field: any, order: any) {
    let sortOrder = SearchSearchRequest_sort_order.ascending;
    const localQueryFields: Array<IFieldsRequirement> = [];
    if (order === -1) {
      sortOrder = SearchSearchRequest_sort_order.descending;
    }
    let localSearchRequest: LocalSearchRequest;
    const model = this.getValues();
    if (model !== null && model.length !== 0) {
      const instance = SearchUtil.getModelInfo(Utility.findCategoryByKind(this.kind), this.kind);
      model.forEach(ele => {
        // all value post process logic goes here
        let processedValue;
        if (field !== '') {
          const type = Utility.getNestedPropInfo(instance, ele.keyFormControl) ? Utility.getNestedPropInfo(instance, ele.keyFormControl).type : '';
          if (type === 'Date') {
            processedValue = ele.valueFormControl.map(e => this.formatDate(e));
          } else {
            processedValue = ele.valueFormControl.map(e => e);
          }
        } else {
          processedValue = [];
        }
        if (ele.keyFormControl in this.localSearchFields) {
          localQueryFields.push({
            key: ele.keyFormControl,
            operator: ele.operatorFormControl,
            values: processedValue
          });
        }
      });
      // We may need to add maxRecords and Aggregate later on
    }
    localSearchRequest = {
      query: localQueryFields,
      sortBy: field,
      sortOrder: sortOrder
    };
    return localSearchRequest;
  }

  /**
   * Function to perform local search
   * @param searchReq Local Search Query
   * @param searchObj Object which is to be queried
   */
  localSearch(searchReq: LocalSearchRequest, searchObj: { [key: string]: any }): LocalSearchResult {
    const res: LocalSearchResult = {
      searchRes: [],
      err: false
    };
    searchReq.query.forEach(q => {
      this.localSearchQueryRequirement(res, q, searchObj);
    });
    return res;
  }

  localSearchQueryRequirement(res: LocalSearchResult, q: FieldsRequirement | IFieldsRequirement, searchObj: { [key: string]: any; }) {
    if (!res.err && q.key in searchObj) {
      if (this.localSearchFields[q.key] && q.values.length !== 1) {
        res.err = true;
      } else {
        let conditionRes: Array<string> = [];
        q.values.forEach(value => {
          const val = value.toLowerCase();
          const valRes = searchObj[q.key][val];
          conditionRes = (conditionRes && conditionRes.length) > 0 ? _.union(conditionRes, valRes) : valRes;
        });
        res.searchRes = (res && res.searchRes && res.searchRes.length) > 0 ? _.intersection(res.searchRes, conditionRes) : conditionRes;
      }
    }
  }

  /**
   * Generates aggregate of remote and local search results
   * @param remoteSearchResult
   * @param localSearchResult
   */
  generateAggregateResults(remoteSearchResult, localSearchResult): Array<string> {
    let searchResult: Array<string> = [];
    if (localSearchResult != null) {
      searchResult = _.intersection(remoteSearchResult, localSearchResult);
    } else {
      searchResult = remoteSearchResult;
    }
    return searchResult;
  }

  onKeydown(event) {
    if (event.key === 'Enter') {
      this.searchClicked();
    }
  }

  /**
   * Search Clicked logic
   */
  searchClicked() {
    this.searchEmitter.emit(null);
  }

  /**
   * Cancel Clicked logic
   */
  cancelClicked() {
    this.search = '';
    this.generalSearch = '';
    this.formArray = new FormArray([]);
    if (this.fieldRepeater) {
      this.fieldRepeater.formArray = new FormArray([]);
      this.fieldRepeater.initData();
    }
    this.cancelEmitter.emit(null);
  }

  /**
   * Determine show/hide for search and cancel button
   */
  showSearchAndCancel(): boolean {
    return (this.search !== '' && !this.showAdvancedPanel) || this.showAdvancedPanel;
  }
}
