import { HttpClientTestingModule } from '@angular/common/http/testing';
import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatIconRegistry } from '@angular/material';
import { By } from '@angular/platform-browser';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { AuthService } from '@app/services/auth.service';
import { ControllerService } from '@app/services/controller.service';
import { ClusterService } from '@app/services/generated/cluster.service';
import { SearchService } from '@app/services/generated/search.service';
import { WorkloadService } from '@app/services/generated/workload.service';
import { LogPublishersService } from '@app/services/logging/log-publishers.service';
import { LogService } from '@app/services/logging/log.service';
import { MessageService } from '@app/services/message.service';
import { UIConfigsService } from '@app/services/uiconfigs.service';
import { LicenseService } from '@app/services/license.service';
import { RouterLinkStubDirective } from '@common/RouterLinkStub.directive.spec';
import { TestingUtility } from '@common/TestingUtility';
import { Utility } from '@common/Utility';
import { SharedModule } from '@components/shared/shared.module';
import { MaterialdesignModule } from '@lib/materialdesign.module';
import { PrimengModule } from '@lib/primeng.module';
import { ClusterHost, ClusterDistributedServiceCard } from '@sdk/v1/models/generated/cluster';
import { UIRolePermissions } from '@sdk/v1/models/generated/UI-permissions-enum';
import { configureTestSuite } from 'ng-bullet';
import { ConfirmationService } from 'primeng/api';
import { WidgetsModule } from 'web-app-framework';
import { HostsComponent } from './hosts.component';
import { NewhostComponent } from './newhost/newhost.component';
import { WorkloadsComponent } from '@app/components/dashboard/workloads/workloads.component';
import { WorkloadWorkload } from '@sdk/v1/models/generated/workload';


describe('HostsComponent', () => {
  let componentCloud: HostsComponent;
  let fixture: ComponentFixture<HostsComponent>;
  let componentEnterprise: HostsComponent;

  const host1 = {
    'kind': 'Host',
    'api-version': 'v1',
    'meta': {
      'name': 'naples1-host',
      'generation-id': '1',
      'resource-version': '694',
      'uuid': '0fd7d80e-ba31-411d-a4a8-df2a47bf8cf8',
      'creation-time': '2019-04-02T18:09:37.972814339Z',
      'mod-time': '2019-04-02T18:09:37.972817316Z',
      'self-link': '/configs/cluster/v1/hosts/naples1-host'
    },
    'spec': {
      'dscs': [
        {
          'mac-address': '0242.c0a8.1c02'
        }
      ]
    },
    'status': {}
  };
  const host2 = {
    'kind': 'Host',
    'api-version': 'v1',
    'meta': {
      'name': 'test-host',
      'generation-id': '1',
      'resource-version': '112568',
      'uuid': 'aebeced0-d1a2-4d2f-9d6d-0ca6eff85681',
      'creation-time': '2019-04-03T18:57:31.727876245Z',
      'mod-time': '2019-04-03T18:57:31.727880959Z',
      'self-link': '/configs/cluster/v1/hosts/test-host'
    },
    'spec': {
      'dscs': [
        {
          'id': 'test'
        }
      ]
    },
    'status': {}
  };

  const host3 = {
    'kind': 'Host',
    'api-version': 'v1',
    'meta': {
      'name': 'test-host1',
      'generation-id': '1',
      'resource-version': '228646',
      'uuid': '57779a13-792e-4acb-8a89-93c34be739e4',
      'creation-time': '2019-04-04T20:47:41.054985496Z',
      'mod-time': '2019-04-04T20:47:41.054988288Z',
      'self-link': '/configs/cluster/v1/hosts/test-host1'
    },
    'spec': {
      'dscs': [
        {
          'id': 'test3'
        }
      ]
    },
    'status': {}
  };

  // id test => mac address 00ae.cd00.1142
  const naple1 = {
    'kind': 'DistributedServiceCard',
    'api-version': 'v1',
    'meta': {
      'name': '00ae.cd00.1142',
      'generation-id': '1',
      'resource-version': '706999',
      'uuid': '96fa49f5-ccb8-40ac-a314-41dd798fae78',
      'creation-time': '2019-04-02T18:09:39.17373748Z',
      'mod-time': '2019-04-09T07:15:59.574423516Z',
      'self-link': '/configs/cluster/v1/distributedservicecards/00ae.cd00.1142'
    },
    'spec': {
      'admit': true,
      'id': 'test',
      'ip-config': {
        'ip-address': '1.2.3.4'
      },
      'mgmt-mode': 'NETWORK',
      'network-mode': 'INBAND',
      'controllers': [
        '192.168.30.10'
      ]
    },
    'status': {
      'admission-phase': 'admitted',
      'conditions': [
        {
          'type': 'HEALTHY',
          'status': 'TRUE',
          'last-transition-time': '2019-04-09T07:15:51Z'
        }
      ],
      'serial-num': 'FLM18440006',
      'primary-mac': '00ae.cd00.1142',
      'ip-config': {
        'ip-address': '1.1.1.1'
      },
      'system-info': {
        'bios-info': {
          'version': '1.0E'
        },
        'os-info': {
          'type': 'Linux',
          'kernel-relase': '4.4.0-87-generic',
          'processor': 'ARMv7'
        },
        'cpu-info': {
          'speed': '2.0 Ghz'
        },
        'memory-info': {
          'type': 'HBM'
        }
      },
      'interfaces': [
        'lo',
        'eth0',
        'eth1',
        'eth2'
      ],
      'DSCVersion': '1.0E',
      'smartNicSku': '68-0003-02 01',
      'host': 'test-name5'
    }
  };

  const workload1 = {
    'kind': 'Workload',
    'api-version': 'v1',
    'meta': {
      'name': 'w1',
      'tenant': 'default',
      'namespace': 'default',
      'generation-id': '2',
      'resource-version': '139282',
      'uuid': 'f3b01b40-5f21-4fa2-9c63-e3bb0b243d29',
      'labels': {
        'type': 'test'
      },
      'creation-time': '2019-10-18T20:29:41.577867228Z',
      'mod-time': '2019-10-18T22:55:50.128243229Z',
      'self-link': '/configs/workload/v1/tenant/default/workloads/w1'
    },
    'spec': {
      'host-name': 'test-host',
      'interfaces': [
        {
          'mac-address': 'aaaa.bbbb.cccc',
          'micro-seg-vlan': 1,
          'external-vlan': 1,
          'ip-addresses': [
            '1.1.11.1'
          ]
        }
      ]
    }
  };

  const workload2 = {
    'kind': 'Workload',
    'api-version': 'v1',
    'meta': {
      'name': 'w2',
      'tenant': 'default',
      'namespace': 'default',
      'generation-id': '2',
      'resource-version': '139282',
      'uuid': 'f3b01b40-5f21-4fa2-9c63-e3bb0b243d29',
      'labels': {
        'type': 'test'
      },
      'creation-time': '2019-10-18T20:29:41.577867228Z',
      'mod-time': '2019-10-18T22:55:50.128243229Z',
      'self-link': '/configs/workload/v1/tenant/default/workloads/w1'
    },
    'spec': {
      'host-name': 'naples1-host',
      'interfaces': [
        {
          'mac-address': 'aaaa.bbbb.cccc',
          'micro-seg-vlan': 1,
          'external-vlan': 1,
          'ip-addresses': [
            '1.1.11.1'
          ]
        }
      ]
    }
  };

  configureTestSuite(() => {
    TestBed.configureTestingModule({
      declarations: [HostsComponent, NewhostComponent, RouterLinkStubDirective],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        NoopAnimationsModule,
        HttpClientTestingModule,
        PrimengModule,
        WidgetsModule,
        MaterialdesignModule,
        RouterTestingModule,
        SharedModule],
      providers: [
        ControllerService,
        UIConfigsService,
        LicenseService,
        AuthService,
        ConfirmationService,
        LogService,
        LogPublishersService,
        MatIconRegistry,
        ClusterService,
        MessageService,
        WorkloadService,
        SearchService
      ]
    });
  });

  beforeEach(() => {
    Utility.getInstance().clearAllVeniceObjectCacheData(); // prevent using cached data.
    fixture = TestBed.createComponent(HostsComponent);
    componentCloud = fixture.componentInstance;
    componentEnterprise = fixture.componentInstance;
  });

  it('should populate cloud table', () => {
    TestingUtility.setAllPermissions();
    TestingUtility.setCloudMode();
    const serviceCluster = TestBed.get(ClusterService);
    const serviceWorkload = TestBed.get(WorkloadService);

    spyOn(serviceCluster, 'ListHostCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterHost(host1), new ClusterHost(host2)
      ])
    );

    spyOn(serviceCluster, 'ListDistributedServiceCardCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterDistributedServiceCard( naple1)
      ])
    );

    spyOn(serviceWorkload, 'ListWorkloadCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new WorkloadWorkload (workload1), new WorkloadWorkload (workload2)
      ])
    );

    fixture.detectChanges();

    // check table header
    const title = fixture.debugElement.query(By.css('.tableheader-title'));
    expect(title.nativeElement.textContent).toContain('Hosts (2)');
    // check table contents
    const tableBody = fixture.debugElement.query(By.css('.ui-table-scrollable-body tbody'));
    expect(tableBody).toBeTruthy();

    TestingUtility.verifyTable([new ClusterHost(host2), new ClusterHost(host1)], componentCloud.cols, tableBody, {
      'spec.dscs': (fieldElem: DebugElement, rowData: any, rowIndex: number) => {
        expect(fieldElem.nativeElement.textContent).toContain(
          componentCloud.processSmartNics(rowData)[0]['text']
        );  // only works if we for one entry case
      },
      'workloads': (fieldElem: DebugElement, rowData: any, rowIndex: number) => {
        const workloads = componentCloud.getHostWorkloads(rowData);
        expect(workloads.length).toBeGreaterThanOrEqual(0);
        if (workloads.length > 0) {
          expect(fieldElem.nativeElement.textContent).toContain(
            workloads[0].meta.name
          );
        } else {
          expect(fieldElem.nativeElement.textContent.length).toEqual(0);
        }
      }
    }, '', true);  // should not have delete icon as hosts have associated workloads
  });
  it('should populate enterprise table', () => {
    TestingUtility.setAllPermissions();
    TestingUtility.setEnterpriseMode();
    const serviceCluster = TestBed.get(ClusterService);
    const serviceWorkload = TestBed.get(WorkloadService);

    spyOn(serviceCluster, 'ListHostCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterHost(host1), new ClusterHost(host2)
      ])
    );

    spyOn(serviceCluster, 'ListDistributedServiceCardCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterDistributedServiceCard( naple1)
      ])
    );

    spyOn(serviceWorkload, 'ListWorkloadCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new WorkloadWorkload (workload1), new WorkloadWorkload (workload2)
      ])
    );

    fixture.detectChanges();

    // check table header
    const title = fixture.debugElement.query(By.css('.tableheader-title'));
    expect(title.nativeElement.textContent).toContain('Hosts (2)');
    // check table contents
    const tableBody = fixture.debugElement.query(By.css('.ui-table-scrollable-body tbody'));
    expect(tableBody).toBeTruthy();

    TestingUtility.verifyTable([new ClusterHost(host2), new ClusterHost(host1)], componentEnterprise.cols, tableBody, {
      'spec.dscs': (fieldElem: DebugElement, rowData: any, rowIndex: number) => {
        expect(fieldElem.nativeElement.textContent).toContain(
          componentEnterprise.processSmartNics(rowData)[0]['text']
        );  // only works if we for one entry case
      },
      'workloads': (fieldElem: DebugElement, rowData: any, rowIndex: number) => {
        const workloads = componentEnterprise.getHostWorkloads(rowData);
        expect(workloads.length).toBeGreaterThanOrEqual(0);
        if (workloads.length > 0) {
          expect(fieldElem.nativeElement.textContent).toContain(
            workloads[0].meta.name
          );
        } else {
          expect(fieldElem.nativeElement.textContent.length).toEqual(0);
        }
      }
    }, '', true);  // should not have delete icon as hosts have associated workloads
  });

  /**
   *  hosts.c.ts  call fetchDSC --> onComplete ( call watch-A, watch-B, watch-C).
   *  So we do
   *  1. set up subject
   *  2. set up spyOn
   *  3. subject.complete
   *  4. test....
   */
  it('should have correct router links for cloud', () => {
    TestingUtility.setAllPermissions();
    TestingUtility.setCloudMode();
    const serviceCluster = TestBed.get(ClusterService);
    const serviceWorkload = TestBed.get(WorkloadService);

    spyOn(serviceCluster, 'ListHostCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterHost(host1), new ClusterHost(host2)
      ])
    );

    spyOn(serviceCluster, 'ListDistributedServiceCardCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterDistributedServiceCard( naple1)
      ])
    );

    spyOn(serviceWorkload, 'ListWorkloadCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new WorkloadWorkload (workload1), new WorkloadWorkload (workload2)
      ])
    );

    fixture.detectChanges();

    // find DebugElements with an attached RouterLinkStubDirective
    const linkDes = fixture.debugElement
      .queryAll(By.directive(RouterLinkStubDirective));
    // get attached link directive instances
    // using each DebugElement's injector
    const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
    expect(routerLinks.length).toBe(2, 'Should have 2 routerLinks');
    expect(routerLinks[0].linkParams).toBe('/cluster/dscs/00ae.cd00.1142');
    expect(routerLinks[1].linkParams).toBe('/cluster/dscs/0242.c0a8.1c02');
  });

  it('should have correct router links for enterprise', () => {
    TestingUtility.setAllPermissions();
    TestingUtility.setEnterpriseMode();
    const serviceCluster = TestBed.get(ClusterService);
    const serviceWorkload = TestBed.get(WorkloadService);

    spyOn(serviceCluster, 'ListHostCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterHost(host1), new ClusterHost(host2)
      ])
    );

    spyOn(serviceCluster, 'ListDistributedServiceCardCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new ClusterDistributedServiceCard( naple1)
      ])
    );

    spyOn(serviceWorkload, 'ListWorkloadCache').and.returnValue(
      TestingUtility.createDataCacheSubject([
        new WorkloadWorkload (workload1), new WorkloadWorkload (workload2)
      ])
    );

    fixture.detectChanges();

    // find DebugElements with an attached RouterLinkStubDirective
    const linkDes = fixture.debugElement
      .queryAll(By.directive(RouterLinkStubDirective));
    // get attached link directive instances
    // using each DebugElement's injector
    const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
    expect(routerLinks.length).toBe(2, 'Should have 2 routerLinks');
    expect(routerLinks[0].linkParams).toBe('/cluster/dscs/00ae.cd00.1142');
    // expect(routerLinks[1].linkParams).toBe('/workload');
    expect(routerLinks[1].linkParams).toBe('/cluster/dscs/0242.c0a8.1c02');
    // expect(routerLinks[3].linkParams).toBe('/workload');
  });


  describe('RBAC', () => {
    beforeEach(() => {
      const serviceCluster = TestBed.get(ClusterService);
      const serviceWorkload = TestBed.get(WorkloadService);


      spyOn(serviceCluster, 'ListHostCache').and.returnValue(
        TestingUtility.createDataCacheSubject([
          new ClusterHost(host1), new ClusterHost(host2)
        ])
      );

      spyOn(serviceCluster, 'ListDistributedServiceCardCache').and.returnValue(
        TestingUtility.createDataCacheSubject([
          new ClusterDistributedServiceCard( naple1)
        ])
      );

      spyOn(serviceWorkload, 'ListWorkloadCache').and.returnValue(
        TestingUtility.createDataCacheSubject([
          new WorkloadWorkload (workload1), new WorkloadWorkload (workload2)
        ])
      );
    });

    it('naples read permission for cloud mode', () => {
      TestingUtility.addPermissions(
        [UIRolePermissions.clusterdistributedservicecard_read]
      );
      fixture.detectChanges();
      const linkDes = fixture.debugElement
        .queryAll(By.directive(RouterLinkStubDirective));
      // get attached link directive instances
      // using each DebugElement's injector
      const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
      expect(routerLinks.length).toBe(2, 'Should have 2 routerLinks');
    });

    it('no permission  - cloud mode', () => {
      fixture.detectChanges();
      const linkDes = fixture.debugElement
        .queryAll(By.directive(RouterLinkStubDirective));
      // get attached link directive instances
      // using each DebugElement's injector
      const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
      expect(routerLinks.length).toBe(0, 'Should have no routerLinks');
    });

    it('naples read permission for enterprise mode', () => {
      TestingUtility.addPermissions(
        [UIRolePermissions.clusterdistributedservicecard_read]
      );
      fixture.detectChanges();
      const linkDes = fixture.debugElement
        .queryAll(By.directive(RouterLinkStubDirective));
      // get attached link directive instances
      // using each DebugElement's injector
      const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
      expect(routerLinks.length).toBe(2, 'Should have 2 routerLinks');
    });

    it('no permission  - enterprise mode', () => {
      fixture.detectChanges();
      const linkDes = fixture.debugElement
        .queryAll(By.directive(RouterLinkStubDirective));
      // get attached link directive instances
      // using each DebugElement's injector
      const routerLinks = linkDes.map(de => de.injector.get(RouterLinkStubDirective));
      expect(routerLinks.length).toBe(0, 'Should have no routerLinks');
    });

  });
});
