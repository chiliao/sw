import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { RolloutsComponent } from '@app/components/settings-group/systemupgrade/rollouts/rollouts.component';
import { RolloutstatusComponent } from '@app/components/settings-group/systemupgrade/rollouts/rolloutstatus/rolloutstatus.component';

const routes: Routes = [
  {
    path: '',
    component: RolloutsComponent
  },
  {
    path: ':id',
    component: RolloutstatusComponent
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class RolloutsRoutingModule { }
