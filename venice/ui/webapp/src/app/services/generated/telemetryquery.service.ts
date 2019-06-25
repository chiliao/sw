import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Eventtypes } from '@app/enum/eventtypes.enum';
import { VeniceResponse } from '@app/models/frontend/shared/veniceresponse.interface';
import { ControllerService } from '@app/services/controller.service';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';
import { Utility } from '../../common/Utility';
import { GenServiceUtility } from './GenUtility';
import { UIConfigsService } from '../uiconfigs.service';
import { UIRolePermissions } from '@sdk/v1/models/generated/UI-permissions-enum';
import { NEVER } from 'rxjs';
import { MethodOpts } from '@sdk/v1/services/generated/abstract.service';


import { Telemetry_queryv1Service } from '@sdk/v1/services/generated/telemetry_queryv1.service';
import { ITelemetry_queryMetricsQueryList } from '@sdk/v1/models/generated/telemetry_query';
import { SearchUtil } from '@app/components/search/SearchUtil';
import { IAuthUser } from '@sdk/v1/models/generated/auth';

@Injectable()
export class TelemetryqueryService extends Telemetry_queryv1Service {

  // Attributes used by generated services
  protected O_Tenant: string = this.getTenant();
  protected baseUrlAndPort = window.location.protocol + '//' + window.location.hostname + ':' + window.location.port;
  protected oboeServiceMap: { [method: string]: Observable<VeniceResponse> } = {};
  protected serviceUtility: GenServiceUtility;

  constructor(protected _http: HttpClient,
              protected _controllerService: ControllerService,
              protected uiconfigsService: UIConfigsService) {
    super(_http);
    this.serviceUtility = new GenServiceUtility(
      _http,
      (payload) => { this.publishAJAXStart(payload); },
      (payload) => { this.publishAJAXEnd(payload); }
    );
  }

  /**
   * Get the service class-name
   */
  getClassName(): string {
    return this.constructor.name;
  }

  protected publishAJAXStart(eventPayload: any) {
    this._controllerService.publish(Eventtypes.AJAX_START, eventPayload);
  }

  protected publishAJAXEnd(eventPayload: any) {
    this._controllerService.publish(Eventtypes.AJAX_END, eventPayload);
  }

  protected invokeAJAX(method: string, url: string, payload: any, opts: MethodOpts, forceReal: boolean = false): Observable<VeniceResponse> {

    const isMetrics = opts.eventID.includes('Metrics');
    if (isMetrics && method === 'POST') {
      let metricAllowed = true;

      const body = payload as ITelemetry_queryMetricsQueryList;
      body.queries.forEach( q => {
        // If one query has already failed, we return early
        if (metricAllowed === false) {
          return;
        }
        const cat = Utility.findCategoryByKind(q.kind);
        if (cat == null) {
          metricAllowed = false;
          return;
        }
        if (!this.uiconfigsService.isAuthorized(UIRolePermissions[cat + q.kind + '_read'])) {
          metricAllowed = false;
          return;
        }
      });

      // Check if admin, admin always has right to read any metrics
      if (!metricAllowed && !Utility.getInstance().isAdmin()) {
        return NEVER;
      }
    } else {
      const key = this.convertEventID(opts);
      if (!this.uiconfigsService.isAuthorized(key)) {
        return NEVER;
      }
    }

    const isOnline = !this.isToMockData() || forceReal;
    return this.serviceUtility.invokeAJAX(method, url, payload, opts.eventID, isOnline);
  }

  convertEventID(opts: MethodOpts): UIRolePermissions {
    let key: string;
    if (opts.eventID.includes('Fwlogs')) {
      key = 'fwlogsquery' + '_' + 'read';
    }
    return UIRolePermissions[key];
  }


  /**
   * Override-able api
   */
  public isToMockData(): boolean {
    const isUseRealData = this._controllerService.useRealData;
    return (!isUseRealData) ? isUseRealData : environment.isRESTAPIReady;
  }

  /**
   * Get login user tenant information
   */
  getTenant(): string {
    return Utility.getInstance().getTenant();
  }

}
