import { HttpClientTestingModule } from '@angular/common/http/testing';
import { Component } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { MatIconRegistry } from '@angular/material';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { RouterTestingModule } from '@angular/router/testing';
import { LogPublishersService } from '@app/services/logging/log-publishers.service';
import { LogService } from '@app/services/logging/log.service';
import { MessageService } from '@app/services/message.service';
import { MetricsqueryService } from '@app/services/metricsquery.service';
import { SharedModule } from '@components/shared/shared.module';
import { PrimengModule } from '@lib/primeng.module';
import { ConfirmationService } from 'primeng/primeng';
import { WidgetsModule } from 'web-app-framework';
import { TelemetrychartComponent } from './telemetrychart.component';
import { MaterialdesignModule } from '@app/lib/materialdesign.module';
import { ControllerService } from '@app/services/controller.service';
import { configureTestSuite } from 'ng-bullet';
import { UIConfigsService } from '@app/services/uiconfigs.service';
import { LicenseService } from '@app/services/license.service';
import { AuthService } from '@app/services/auth.service';
import { ClusterService } from '@app/services/generated/cluster.service';
import { NetworkService } from '@app/services/generated/network.service';
import { AuthService as AuthServiceGen } from '@app/services/generated/auth.service';
import { MetricsMetadataIris } from '@sdk/metrics/generated/iris_metadata';
import { setMetricsMetadata } from '@sdk/metrics/generated/metadata';


@Component({
  template: ''
})
class DummyComponent { }

describe('TelemetrychartComponent', () => {
  let component: TelemetrychartComponent;
  let fixture: ComponentFixture<TelemetrychartComponent>;

  configureTestSuite(() => {
    TestBed.configureTestingModule({
      declarations: [
        DummyComponent,
      ],
      imports: [
        RouterTestingModule.withRoutes([
          { path: 'login', component: DummyComponent }
        ]),
        FormsModule,
        HttpClientTestingModule,
        WidgetsModule,
        PrimengModule,
        MaterialdesignModule,
        NoopAnimationsModule,
        SharedModule
      ],
      providers: [
        ControllerService,
        ClusterService,
        NetworkService,
        ConfirmationService,
        MatIconRegistry,
        LogService,
        LogPublishersService,
        MessageService,
        MetricsqueryService,
        UIConfigsService,
        LicenseService,
        AuthService,
        AuthServiceGen
      ]
    });
  });

  beforeEach(() => {
    setMetricsMetadata(MetricsMetadataIris);
    fixture = TestBed.createComponent(TelemetrychartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
