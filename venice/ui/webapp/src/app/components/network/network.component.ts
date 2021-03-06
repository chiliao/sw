import { Component, OnInit, ViewEncapsulation, ChangeDetectorRef, ViewChild } from '@angular/core';
import { Animations } from '@app/animations';
import { Icon } from '@app/models/frontend/shared/icon.interface';
import { IApiStatus, NetworkNetwork, INetworkNetwork, NetworkOrchestratorInfo } from '@sdk/v1/models/generated/network';
import { CustomExportMap, TableCol } from '@app/components/shared/tableviewedit';
import { Subscription, Observable } from 'rxjs';
import { HttpEventUtility } from '@app/common/HttpEventUtility';
import { NetworkService } from '@app/services/generated/network.service';
import { OrchestrationService } from '@app/services/generated/orchestration.service';
import { UIConfigsService } from '@app/services/uiconfigs.service';
import { ControllerService } from '@app/services/controller.service';
import { UIRolePermissions } from '@sdk/v1/models/generated/UI-permissions-enum';
import { Utility } from '@app/common/Utility';
import { OrchestrationOrchestrator } from '@sdk/v1/models/generated/orchestration';
import { SelectItem } from 'primeng/api';
import { WorkloadWorkload } from '@sdk/v1/models/generated/workload';
import { WorkloadService } from '@app/services/generated/workload.service';
import { NetworkWorkloadsTuple, ObjectsRelationsUtility } from '@app/common/ObjectsRelationsUtility';
import { PentableComponent } from '@app/components/shared/pentable/pentable.component';
import { DataComponent } from '@app/components/shared/datacomponent/datacomponent.component';
import { WorkloadUtility, WorkloadNameInterface } from '@app/common/WorkloadUtility';

interface NetworkUIModel {
  associatedWorkloads: WorkloadWorkload[];
}

@Component({
  selector: 'app-network',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './network.component.html',
  styleUrls: ['./network.component.scss'],
  animations: [Animations]
})

export class NetworkComponent extends DataComponent implements OnInit {

  @ViewChild('networkTable') networkTable: PentableComponent;

  bodyicon: Icon = {
    margin: {
      top: '9px',
      left: '8px'
    },
    svgIcon: 'host'
  };

  headerIcon: Icon = {
    margin: {
      top: '0px',
      left: '10px',
    },
    matIcon: 'computer'
  };

  exportFilename: string = 'PSM-networks';

  exportMap: CustomExportMap = {
    'associatedWorkloads': (opts): string => {
      return (opts.data._ui.associatedWorkloads) ? opts.data._ui.associatedWorkloads.map(wkld => wkld.meta.name).join(', ') : '';
    },
    'spec.orchestrators': (opts): string => {
      return (opts.data.spec.orchestrators) ? opts.data.spec.orchestrators.map(or => or['orchestrator-name'] + ' Datacenter: ' + or.namespace).join(', ') : '';
    }
  };

  vcenters: ReadonlyArray<OrchestrationOrchestrator> = [];
  vcenterOptions: SelectItem[] = [];

  workloadList: WorkloadWorkload[] = [];

  subscriptions: Subscription[] = [];
  dataObjects: ReadonlyArray<NetworkNetwork> = [];
  networkEventUtility: HttpEventUtility<NetworkNetwork>;

  disableTableWhenRowExpanded: boolean = true;
  isTabComponent: boolean = false;

  // Used for the table - when true there is a loading icon displayed
  tableLoading: boolean = false;

  cols: TableCol[] = [
    { field: 'meta.name', header: 'Name', class: 'network-column-name', sortable: true, width: 20 },
    { field: 'spec.vlan-id', header: 'VLAN', class: 'network-column-vlan', sortable: true, width: '80px'},
    { field: 'spec.orchestrators', header: 'Orchestrators', class: 'network-column-orchestrators', sortable: false, width: 35 },
    { field: 'associatedWorkloads', header: 'Workloads', class: '', sortable: false, width: 35 },
    { field: 'meta.creation-time', header: 'Creation Time', class: 'vcenter-integration-column-date', sortable: true, width: '180px' }
  ];

  constructor(private networkService: NetworkService,
    protected cdr: ChangeDetectorRef,
    protected uiconfigsService: UIConfigsService,
    protected orchestrationService: OrchestrationService,
    protected workloadService: WorkloadService,
    protected controllerService: ControllerService) {
    super(controllerService, uiconfigsService);
  }

  getNetworks() {
    const hostSubscription = this.networkService.ListNetworkCache().subscribe(
      (response) => {
        if (response.connIsErrorState) {
          return;
        }
        this.dataObjects = this.buildNetworkWorkloadsMap(response.data);
        this.tableLoading = false;
      },
      (error) => {
        this.tableLoading = false;
        this.controllerService.invokeRESTErrorToaster('Error', 'Failed to get networks');
      }
    );
    this.subscriptions.push(hostSubscription);
  }

  getVcenterIntegrations() {
    const sub = this.orchestrationService.ListOrchestratorCache().subscribe(
      response => {
        if (response.connIsErrorState) {
          return;
        }
        this.vcenters = response.data;
        this.vcenterOptions = this.vcenters.map(vcenter => {
          return {
            label: vcenter.meta.name,
            value: vcenter.meta.name
          };
        });
        this.vcenterOptions.push({label: '', value: null});
      },
      this.controllerService.webSocketErrorHandler('Failed to get vCenters')
    );
    this.subscriptions.push(sub);
  }

  /**
   * Fetch workloads.
   */
  watchWorkloads() {
    const workloadSubscription = this.workloadService.ListWorkloadCache().subscribe(
      (response) => {
        if (response.connIsErrorState) {
          return;
        }
        this.workloadList = response.data as WorkloadWorkload[];
        this.dataObjects = this.buildNetworkWorkloadsMap(this.dataObjects);
      }
    );
    this.subscriptions.push(workloadSubscription);
  }

  setDefaultToolbar() {
    let buttons = [];
    if (this.uiconfigsService.isAuthorized(UIRolePermissions.networknetwork_create)) {
      buttons = [{
        cssClass: 'global-button-primary networks-button networks-button-ADD',
        text: 'ADD NETWORK',
        computeClass: () =>  !(this.networkTable.showRowExpand) ? '' : 'global-button-disabled',
        callback: () => { this.networkTable.createNewObject(); }
      }];
    }
    this.controllerService.setToolbarData({
      buttons: buttons,
      breadcrumb: [{ label: 'Network ', url: Utility.getBaseUIUrl() + 'networks' }]
    });
  }

  displayColumn(rowData: NetworkNetwork, col: TableCol): any {
    const fields = col.field.split('.');
    const value = Utility.getObjectValueByPropertyPath(rowData, fields);
    const column = col.field;
    switch (column) {
      case 'spec.orchestrators':
        return this.displayColumn_orchestrators(value);
      case 'spec.vlan-id':
        return value ? value : 0;
      default:
        return Array.isArray(value) ? JSON.stringify(value, null, 2) : value;
    }
  }

  hasWorkloads(rowData: NetworkNetwork): boolean {
    const workloads = rowData._ui.associatedWorkloads;
    return workloads && workloads.length > 0;
  }

  displayColumn_orchestrators(values: NetworkOrchestratorInfo[]): any {
    const map: {'vCenter': string, 'dataCenters': string[]} = {} as any;
    values.forEach((value: NetworkOrchestratorInfo) => {
      if (!map[value['orchestrator-name']]) {
        map[value['orchestrator-name']] = [value.namespace];
      } else {
        map[value['orchestrator-name']].push(value.namespace);
      }
    });
    let result: string = '';
    for (const key of Object.keys(map)) {
      const eachRow: string = 'vCenter: ' + key + ', Datacenter: ' + map[key].join(', ');
      result += '<div class="ellipsisText" title="' + eachRow + '">' + eachRow + '</div>';
    }
    return result;
  }

  ngOnInit() {
    this.tableLoading = true;
    this.setDefaultToolbar();
    this.getNetworks();
    this.getVcenterIntegrations();
    this.watchWorkloads();
  }

  creationFormClose() {
    this.networkTable.creationFormClose();
  }

  editFormClose(rowData) {
    if (this.networkTable.showRowExpand) {
      this.networkTable.toggleRow(rowData);
    }
  }

  expandRowRequest(event, rowData) {
    if (!this.networkTable.showRowExpand) {
      this.networkTable.toggleRow(rowData, event);
    }
  }

  onColumnSelectChange(event) {
    this.networkTable.onColumnSelectChange(event);
  }

  onDeleteRecord(event, object) {
    this.networkTable.onDeleteRecord(
      event,
      object,
      this.generateDeleteConfirmMsg(object),
      this.generateDeleteSuccessMsg(object),
      this.deleteRecord.bind(this),
      () => {
        this.networkTable.selectedDataObjects = [];
      }
    );
  }

  buildNetworkWorkloadsMap(responseData: any = []) {
    const networkWorkloadsTuple: NetworkWorkloadsTuple =
      ObjectsRelationsUtility.buildNetworkWorkloadsMap(this.workloadList || [], responseData);
    return responseData.map(network => {
      const associatedWorkloads: WorkloadWorkload[] =
        networkWorkloadsTuple[network.meta.name] || [];
      const uiModel: NetworkUIModel = { associatedWorkloads };
      network._ui = uiModel;
      return network;
    });
  }

  deleteRecord(object: NetworkNetwork): Observable<{ body: INetworkNetwork | IApiStatus | Error; statusCode: number }> {
    return this.networkService.DeleteNetwork(object.meta.name);
  }

  generateDeleteConfirmMsg(object: INetworkNetwork): string {
    return 'Are you sure you want to delete network ' + object.meta.name;
  }

  generateDeleteSuccessMsg(object: INetworkNetwork): string {
    return 'Deleted network ' + object.meta.name;
  }

  getClassName(): string {
    return this.constructor.name;
  }

  getSelectedDataObjects(): any[] {
    return this.networkTable.selectedDataObjects;
  }

  clearSelectedDataObjects() {
    this.networkTable.selectedDataObjects = [];
  }
}

