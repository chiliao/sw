import { AbstractService } from './abstract.service';
import { HttpClient } from '../../../../webapp/node_modules/@angular/common/http';
import { Observable } from '../../../../webapp/node_modules/rxjs';
import { Injectable } from '../../../../webapp/node_modules/@angular/core';

import { IRolloutRolloutList,IApiStatus,IRolloutRollout,IRolloutRolloutAction,IRolloutAutoMsgRolloutWatchHelper,IRolloutAutoMsgRolloutActionWatchHelper } from '../../models/generated/rollout';

@Injectable()
export class Rolloutv1Service extends AbstractService {
  constructor(protected _http: HttpClient) {
    super(_http);
  }

  /**
   * Override super
   * Get the service class-name
  */
  getClassName(): string {
    return this.constructor.name;
  }

  /** List Rollout objects */
  public ListRollout(queryParam: any = null, stagingID: string = ""):Observable<{body: IRolloutRolloutList | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rollout';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXGetCall(url, queryParam, 'ListRollout') as Observable<{body: IRolloutRolloutList | IApiStatus | Error, statusCode: number}>;
  }
  
  public DoRollout(body: IRolloutRollout, stagingID: string = ""):Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rollout/DoRollout';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXPostCall(url, body, 'DoRollout') as Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get Rollout object */
  public GetRollout(O_Name, queryParam: any = null, stagingID: string = ""):Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rollout/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXGetCall(url, queryParam, 'GetRollout') as Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete Rollout object */
  public DeleteRollout(O_Name, stagingID: string = ""):Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rollout/{O.Name}';
    url = url.replace('{O.Name}', O_Name);
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXDeleteCall(url, 'DeleteRollout') as Observable<{body: IRolloutRollout | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Get RolloutAction object */
  public GetRolloutAction(queryParam: any = null, stagingID: string = ""):Observable<{body: IRolloutRolloutAction | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rolloutAction';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXGetCall(url, queryParam, 'GetRolloutAction') as Observable<{body: IRolloutRolloutAction | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Delete RolloutAction object */
  public DeleteRolloutAction(stagingID: string = ""):Observable<{body: IRolloutRolloutAction | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/rolloutAction';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXDeleteCall(url, 'DeleteRolloutAction') as Observable<{body: IRolloutRolloutAction | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch Rollout objects. Supports WebSockets or HTTP long poll */
  public WatchRollout(queryParam: any = null, stagingID: string = ""):Observable<{body: IRolloutAutoMsgRolloutWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/watch/rollout';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXGetCall(url, queryParam, 'WatchRollout') as Observable<{body: IRolloutAutoMsgRolloutWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
  /** Watch RolloutAction objects. Supports WebSockets or HTTP long poll */
  public WatchRolloutAction(queryParam: any = null, stagingID: string = ""):Observable<{body: IRolloutAutoMsgRolloutActionWatchHelper | IApiStatus | Error, statusCode: number}> {
    let url = this['baseUrlAndPort'] + '/configs/rollout/v1/watch/rolloutAction';
    if (stagingID != null && stagingID.length != 0) {
      url = url.replace('configs', 'staging/' + stagingID);
    }
    return this.invokeAJAXGetCall(url, queryParam, 'WatchRolloutAction') as Observable<{body: IRolloutAutoMsgRolloutActionWatchHelper | IApiStatus | Error, statusCode: number}>;
  }
  
}