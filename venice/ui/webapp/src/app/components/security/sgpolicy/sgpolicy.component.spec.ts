import { HttpClientTestingModule } from '@angular/common/http/testing';
import { DebugElement } from '@angular/core';
/**-----
 Angular imports
 ------------------*/
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatIconRegistry } from '@angular/material';
import { By } from '@angular/platform-browser';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { TestingUtility } from '@app/common/TestingUtility';
import { Utility } from '@app/common/Utility';
import { PrettyDatePipe } from '@app/components/shared/Pipes/PrettyDate.pipe';
import { SharedModule } from '@app/components/shared/shared.module';
/**-----
 Venice web-app imports
 ------------------*/
import { ControllerService } from '@app/services/controller.service';
import { SearchService } from '@app/services/generated/search.service';
import { LogPublishersService } from '@app/services/logging/log-publishers.service';
import { LogService } from '@app/services/logging/log.service';
import { MaterialdesignModule } from '@lib/materialdesign.module';
import { PrimengModule } from '@lib/primeng.module';
import { SearchPolicySearchRequest } from '@sdk/v1/models/generated/search';
import { SecurityService } from 'app/services/generated/security.service';
import { BehaviorSubject } from 'rxjs/BehaviorSubject';
import { SgpolicyComponent } from './sgpolicy.component';

/**
 * We don't use the verify table in TestingUtility as we need to skip the
 * first field index
 */
function verifyTable(data: any[], columns: any[], tableElem: DebugElement) {
  const rows = tableElem.queryAll(By.css('tr'));
  expect(rows.length).toBe(data.length, 'Data did not match number of entries in the table');
  rows.forEach((row, rowIndex) => {
    const rowData = data[rowIndex];
    row.children.forEach((field, fieldIndex) => {
      if (fieldIndex === 0) {
        // SG Policy doesn't have a header for the first column
        expect(field.nativeElement.textContent).toContain(rowIndex);
        return;
      }
      const colData = columns[fieldIndex - 1];
      switch (colData.field) {
        case 'sourceIPs':
          expect(field.nativeElement.textContent)
            .toContain(rowData['from-ip-addresses'].join(', '),
              'source IPs time did not match for row ' + rowIndex);
          break;

        case 'destIPs':
          expect(field.nativeElement.textContent)
            .toContain(rowData['to-ip-addresses'].join(', '),
              'dest IPs time did not match for row ' + rowIndex);
          break;

        case 'action':
          expect(field.nativeElement.textContent).toContain(rowData.action);
          break;
        case 'protocolPort':
          expect(field.nativeElement.textContent).toContain(rowData.apps.join(', '));
          break;
        default:
          const fieldData = Utility.getObjectValueByPropertyPath(data[rowIndex], colData.field.split('.'));
          expect(field.nativeElement.textContent).toContain(fieldData, colData.header + ' did not match');
      }
    });
  });
}

function setNewRuleData(component, fixture, numRules: number = 40) {
  let tableRules = [];
  for (let index = 0; index < numRules; index++) {
    tableRules.push({
      order: index,
      rule: {
        "apps": [
          "tcp/8080",
        ],
        "action": "PERMIT",
        "from-ip-addresses": [
          '10.1.1.0'
        ],
        "to-ip-addresses": [
          '10.1.1.8'
        ]
      }
    })
  }
  component.sgPolicyRules = tableRules;
  fixture.detectChanges();
}

describe('SgpolicyComponent', () => {
  let component: SgpolicyComponent;
  let fixture: ComponentFixture<SgpolicyComponent>;
  let testingUtility: TestingUtility

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [SgpolicyComponent],
      imports: [
        RouterTestingModule,
        FormsModule,
        ReactiveFormsModule,
        NoopAnimationsModule,
        SharedModule,
        HttpClientTestingModule,
        PrimengModule,
        MaterialdesignModule
      ],
      providers: [
        ControllerService,
        LogService,
        LogPublishersService,
        MatIconRegistry,
        SearchService,
        SecurityService
      ]
    })
      .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SgpolicyComponent);
    component = fixture.componentInstance;
    testingUtility = new TestingUtility(fixture);
    component.cols = [
      { field: 'meta', header: 'Meta' },
      { field: 'spec', header: 'Spec' },
      { field: 'status', header: 'Status' }
    ];
  });

  afterEach(() => {
    TestBed.resetTestingModule();
  });

  it('should use source/dest IP and port fields when making query', () => {
    fixture.detectChanges();

    const _ = Utility.getLodash();
    const service = TestBed.get(SearchService);
    const postQuerySpy = spyOn(service, 'PostPolicyQuery').and.returnValue(
      new BehaviorSubject<any>(
        {
          body: {
            status: 'MISS'
          }
        }
      ));

    const inputs = fixture.debugElement.queryAll(By.css('mat-form-field'));
    const sourceIPInput = inputs[0].query(By.css('input'));
    const destIPInput = inputs[1].query(By.css('input'));
    const portInput = inputs[2].query(By.css('input'));

    // Putting in text into source IP
    testingUtility.setText(sourceIPInput, '10.1.1.1');
    testingUtility.sendEnterKeyup(sourceIPInput);
    let expectedReq = new SearchPolicySearchRequest();
    expectedReq["to-ip-address"] = "any";
    expectedReq["from-ip-address"] = '10.1.1.1';
    expectedReq["tenant"] = 'default';
    let calledObj = postQuerySpy.calls.mostRecent().args[0];
    expect(_.isEqual(expectedReq.getValues(), calledObj.getValues())).toBeTruthy();

    // Putting in text into dest IP
    testingUtility.setText(destIPInput, '10.1.1.8');
    testingUtility.sendEnterKeyup(sourceIPInput);
    expectedReq = new SearchPolicySearchRequest();
    expectedReq["from-ip-address"] = '10.1.1.1';
    expectedReq["tenant"] = 'default';
    expectedReq["to-ip-address"] = "10.1.1.8";
    calledObj = postQuerySpy.calls.mostRecent().args[0];
    expect(_.isEqual(expectedReq.getValues(), calledObj.getValues())).toBeTruthy();

    // Putting text into APP
    testingUtility.setText(portInput, 'tcp/80');
    testingUtility.sendEnterKeyup(sourceIPInput);
    expectedReq = new SearchPolicySearchRequest();
    expectedReq["from-ip-address"] = '10.1.1.1';
    expectedReq["tenant"] = 'default';
    expectedReq["to-ip-address"] = "10.1.1.8";
    expectedReq["app"] = "tcp/80";
    calledObj = postQuerySpy.calls.mostRecent().args[0];
    expect(_.isEqual(expectedReq.getValues(), calledObj.getValues())).toBeTruthy();

  });

  it('should show search/cancel buttons when there is input and error message when input searched is not a valid IP', () => {
    fixture.detectChanges();

    const invokePolicySearchSpy = spyOn(component, 'invokePolicySearch').and.callThrough();
    const service = TestBed.get(SearchService);
    const querySpy = spyOn(service, 'PostPolicyQuery').and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MISS'
        }
      }
      ));

    // No inputs have text, so search and clear buttons are undefined
    // There should be no error message
    let searchButton = fixture.debugElement.query(By.css('.sgpolicy-search-button'));
    let searchClearButton = fixture.debugElement.query(By.css('.sgpolicy-search-clear-button'));
    let errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(searchButton).toBeNull();
    expect(searchClearButton).toBeNull();
    expect(errorMessageDiv).toBeNull();

    // Putting in text into source IP
    const inputs = fixture.debugElement.queryAll(By.css('mat-form-field'));
    const sourceIPInput = inputs[0].query(By.css('input'));
    const portInput = inputs[2].query(By.css('input'));
    testingUtility.setText(sourceIPInput, '192');


    // The search and cancel buttons should appear
    searchButton = fixture.debugElement.query(By.css('.sgpolicy-search-button'));
    searchClearButton = fixture.debugElement.query(By.css('.sgpolicy-search-clear-button'));
    expect(searchButton).toBeTruthy();
    expect(searchClearButton).toBeTruthy();

    // Click the search button should invoke a search
    testingUtility.sendClick(searchButton);
    expect(component.invokePolicySearch).toHaveBeenCalled();
    expect(service.PostPolicyQuery).toHaveBeenCalledTimes(0)

    // There should be an invalid IP message
    errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeTruthy();
    expect(errorMessageDiv.children[1].nativeElement.textContent).toContain('Invalid IP');

    // Typing again should remove the message if the content is different
    testingUtility.setText(sourceIPInput, '192');
    errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeTruthy();
    testingUtility.setText(sourceIPInput, '192.10');
    // Listens for key up, so we trigger one with a random keyCode
    sourceIPInput.triggerEventHandler('keyup', { keyCode: 20 });
    fixture.detectChanges();
    errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeNull();

    // Clicking clear button should empty out the results, but not reset the scroll
    // since we don't have a match
    // TODO: Find a way to check scroll
    testingUtility.sendClick(searchClearButton)
    expect(sourceIPInput.nativeElement.value).toBe('');

    // Should allow port only search
    testingUtility.setText(sourceIPInput, '');
    testingUtility.setText(portInput, 'tcp/88');
    searchButton = fixture.debugElement.query(By.css('.sgpolicy-search-button'));
    expect(searchButton).toBeTruthy();
    // Click the search button should invoke a search
    testingUtility.sendClick(searchButton);
    expect(service.PostPolicyQuery).toHaveBeenCalledTimes(1)
    const req = querySpy.calls.first().args[0];
    expect(req['from-ip-address']).toBe('any');
    expect(req['to-ip-address']).toBe('any');
    expect(req.app).toBe('tcp/88');
  });

  it('should display sgpolicy meta/rules and highlight matching row or display there is none on search', () => {
    const securityService = TestBed.get(SecurityService);
    const rules = [
      {
        "apps": [
          "tcp/80",
          "udp/53"
        ],
        "action": "PERMIT",
        "from-ip-addresses": [
          "172.0.0.1"
        ],
        "to-ip-addresses": [
          "192.168.1.1/16"
        ]
      },
      {
        "apps": [
          "tcp/84",
        ],
        "action": "DENY",
        "from-ip-addresses": [
          '10.1.1.2'
        ],
        "to-ip-addresses": [
          '10.1.1.3'
        ]
      },
      {
        "apps": [
          "tcp/8080",
        ],
        "action": "PERMIT",
        "from-ip-addresses": [
          '10.1.1.0'
        ],
        "to-ip-addresses": [
          '10.1.1.8'
        ]
      }
    ];

    const sgPolicy = {
      meta: {
        name: 'policy1',
        "mod-time": '2018-08-23T17:35:08.534909931Z',
        "creation-time": '2018-08-23T17:30:08.534909931Z'
      },
      spec: {
        "rules": rules
      }
    }

    const sgPolicyObserver = new BehaviorSubject({
      body: {
        result: {
          Events: [
            {
              Type: "Created",
              Object: sgPolicy
            }
          ]
        }
      }
    });

    const sgPolicyWatch = spyOn(securityService, 'WatchSGPolicy').and.returnValue(
      sgPolicyObserver
    );

    fixture.detectChanges();
    const fields = fixture.debugElement.queryAll(By.css('.sgpolicy-summary-panel-content-value'));
    expect(fields.length).toBe(3);
    expect(fields[0].nativeElement.textContent).toContain(sgPolicy.meta.name);
    // Creation time
    const formattedCreationTime = new PrettyDatePipe('en-US').transform(sgPolicy.meta["creation-time"]);
    expect(fields[1].nativeElement.textContent).toContain(formattedCreationTime);
    // Mod time
    const formattedModTime = new PrettyDatePipe('en-US').transform(sgPolicy.meta["mod-time"]);
    expect(fields[2].nativeElement.textContent).toContain(formattedModTime);
    let tableBody = fixture.debugElement.query(By.css('.ui-table-scrollable-body-table tbody'));
    verifyTable(sgPolicy.spec.rules, component.cols, tableBody);


    // Mocking enough entries to test search highlighting
    setNewRuleData(component, fixture, 40);
    const searchService = TestBed.get(SearchService);
    const postQuerySpy = spyOn(searchService, 'PostPolicyQuery').and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MATCH',
          results: {
            "policy1": {
              "index": 10
            }
          }
        }
      }
      ));
    spyOn(component.lazyRenderWrapper, 'scrollToRowNumber');
    component.invokePolicySearch('10.1.1.1');
    fixture.detectChanges();
    expect(searchService.PostPolicyQuery).toHaveBeenCalled();
    // 10th row should be highlighted and we should have scrolled to it
    expect(component.selectedRuleIndex).toBe(10);
    let rule = fixture.debugElement.query(By.css('.sgpolicy-match'));
    expect(rule).toBeTruthy();
    expect(rule.children[0].nativeElement.textContent).toContain(10);
    expect(component.lazyRenderWrapper.scrollToRowNumber).toHaveBeenCalledWith(10);
    // If user starts to modify search, the highlight row disappears
    const inputs = fixture.debugElement.queryAll(By.css('mat-form-field'));
    const sourceIPInput = inputs[0].query(By.css('input'));
    testingUtility.setText(sourceIPInput, '192');
    // Listens for key up, so we trigger one with a random keyCode
    sourceIPInput.triggerEventHandler('keyup', { keyCode: 20 });
    fixture.detectChanges();
    rule = fixture.debugElement.query(By.css('.sgpolicy-match'));
    expect(rule).toBeNull();

    // Search with a miss or missing policy name
    postQuerySpy.and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MISS'
        }
      }
      ));
    component.invokePolicySearch('10.1.1.1');
    fixture.detectChanges();
    let errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeTruthy();
    expect(errorMessageDiv.children[1].nativeElement.textContent).toContain('No Matching Rule');

    postQuerySpy.and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MATCH',
          results: {
            "randomPolicy": {
              "index": 10
            }
          }
        }
      }
      ));
    component.invokePolicySearch('10.1.1.1');
    fixture.detectChanges();
    errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeTruthy();
    expect(errorMessageDiv.children[1].nativeElement.textContent).toContain('No Matching Rule');

    // Setting a valid search for the following test examples
    postQuerySpy.and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MATCH',
          results: {
            "policy1": {
              "index": 10
            }
          }
        }
      }
      ));
    component.invokePolicySearch('10.1.1.1');
    fixture.detectChanges();
    errorMessageDiv = fixture.debugElement.query(By.css('.sgpolicy-search-error'));
    expect(errorMessageDiv).toBeNull();

    // Causing new data button
    component.lazyRenderWrapper.hasUpdate = true;

    // Setting up for the next search
    postQuerySpy.calls.reset();
    postQuerySpy.and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MATCH',
          results: {
            "policy1": {
              "index": 5
            }
          }
        }
      }
      ));

    // if user clicks new data button, we should redo the search
    component.lazyRenderWrapper.resetTableView(); // Equivalent of clicking update data button
    fixture.detectChanges();
    expect(searchService.PostPolicyQuery).toHaveBeenCalled();
    // 10th row should be highlighted and we should have scrolled to it
    expect(component.selectedRuleIndex).toBe(5);
    rule = fixture.debugElement.query(By.css('.sgpolicy-match'));
    expect(rule).toBeTruthy();
    expect(rule.children[0].nativeElement.textContent).toContain(5);

    // If there is new data, and user tries to perform a search
    // we should force an update to the new data
    component.lazyRenderWrapper.hasUpdate = true;
    component.invokePolicySearch('10.1.1.1');
    fixture.detectChanges();
    // Should have switched to new data and invoked search
    expect(component.lazyRenderWrapper.hasUpdate).toBeFalsy();
    expect(searchService.PostPolicyQuery).toHaveBeenCalledTimes(2);

    postQuerySpy.calls.reset();
    postQuerySpy.and.returnValue(
      new BehaviorSubject<any>({
        body: {
          status: 'MATCH',
          results: {
            "policy1": {
              "index": 0
            }
          }
        }
      }
      ));

    // Test deleting the sgpolicy clears the table
    sgPolicyObserver.next({
      body: {
        result: {
          Events: [
            {
              Type: "Deleted",
              Object: sgPolicy
            }
          ]
        }
      }
    });
    fixture.detectChanges();
    tableBody = fixture.debugElement.query(By.css('.ui-table-scrollable-body-table tbody'));
    verifyTable([], component.cols, tableBody);
  });

});
