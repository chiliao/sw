import { Component, EventEmitter, Input, OnInit, Output, OnDestroy, OnChanges, ViewEncapsulation, SimpleChanges } from '@angular/core';
import { FormArray, AbstractControl, Validators, FormGroup, ValidatorFn, ValidationErrors } from '@angular/forms';
import { required } from '@sdk/v1/utils/validators';
import { UsersComponent } from '../users.component';
import { ErrorStateMatcher } from '@angular/material';
import { Animations } from '@app/animations';
import { ControllerService } from '@app/services/controller.service';
import { AuthService } from '@app/services/generated/auth.service';
import { StagingService } from '@app/services/generated/staging.service';
import { AuthRoleBinding, AuthUser, AuthRole } from '@sdk/v1/models/generated/auth';
import { Utility } from '@app/common/Utility';


export interface RolebindingUISelectItem {
  name: string;
  label: string;
  value: string;
}

/**
 * This component lets user to create and update role-binding objects.
 * In UI, we let admin to pick one Auth-role to bind to and assign users to role-bindings.
 *
 * Role-binding CRUD operations don't require commit buffer.
 */
@Component({
  selector: 'app-newrolebinding',
  templateUrl: './newrolebinding.component.html',
  styleUrls: ['./newrolebinding.component.scss'],

  animations: [Animations],
  encapsulation: ViewEncapsulation.None
})
export class NewrolebindingComponent extends UsersComponent implements OnInit, OnDestroy, OnChanges {

  newAuthRolebinding: AuthRoleBinding;

  // define data structures for UI
  sourceUsers: RolebindingUISelectItem[] = [];
  targetUsers: RolebindingUISelectItem[] = [];
  sourceRoles: RolebindingUISelectItem[] = [];

  @Input() selectedAuthRolebinding: AuthRoleBinding;
  @Output() formClose: EventEmitter<any> = new EventEmitter();
  @Input() veniceUsers: AuthUser[] = [];  // hosted UI will pass in available AuthUsers
  @Input() veniceRoles: AuthRole[] = [];  // hosted UI will pass in available AuthRoles
  @Input() veniceRolebindings: AuthRoleBinding[] = [];  // hosted UI will pass in existing AuthRoleBindings

  constructor(protected _controllerService: ControllerService,
    protected _authService: AuthService,
    protected stagingService: StagingService
  ) {
    super(_controllerService, _authService, stagingService, null);
  }

  /**
   * override super.api()
   */
  getClassName(): string {
    return this.constructor.name;
  }

  ngOnInit() {
    if (!this.newAuthRolebinding) {
      this.setupData();
    }
  }

  ngOnChanges(changes: SimpleChanges) {
    this.setupData();
  }

  ngOnDestroy() { }

  setupData() {
    if (this.isEditMode()) {
      this.newAuthRolebinding = this.getRoleFromSelectedRoleBinding();
    } else {
      this.newAuthRolebinding = new AuthRoleBinding();
      this.newAuthRolebinding.$formGroup.get(['meta', 'name']).setValidators([
        this.newAuthRolebinding.$formGroup.get(['meta', 'name']).validator,
        this.isRolebindingnameValid(this.veniceRolebindings)
      ]);

      this.onUsersAvailable();
    }
    this.onRolesAvailable();
  }

  isRolebindingnameValid(authRolebindings: AuthRoleBinding[]): ValidatorFn {
    return (control: AbstractControl): ValidationErrors | null => {
      if (  Utility.isEmpty(control.value) || this.isRoleBindingAlreadyExist(control.value, authRolebindings)) {
        return {
          'rolebinding-name': {
            required: true,
            message: 'Rolebinding name is required and must be unique'
          }
        };
      }
      return null;
    };
  }

  onUsersAvailable() {
    if (this.veniceUsers) {
      this.sourceUsers.length = 0;
      this.veniceUsers.forEach(authUser => {
        this.sourceUsers.push(
          {
            name: authUser.meta.name,
            label: authUser.meta.name,
            value: authUser.meta.name
          }
        );
      });
    }
  }

  onRolesAvailable() {
    if (this.veniceRoles) {
      this.sourceRoles.length = 0;
      this.veniceRoles.forEach(authRole => {
        this.sourceRoles.push(
          {
            name: authRole.meta.name,
            label: authRole.meta.name,
            value: authRole.meta.name
          }
        );
      });
    }
  }

  isEditMode(): boolean {
    return (this.selectedAuthRolebinding) ? true : false;
  }

  getRoleFromSelectedRoleBinding(): AuthRoleBinding {
    const roleBindingData = this.selectedAuthRolebinding.getFormGroupValues();
    // build targetUser list for picklist
    this.targetUsers.length = 0;
    roleBindingData.spec.users.forEach(username => {
      this.targetUsers.push(
        {
          name: username,
          label: username,
          value: username
        }
      );
    });
    // get all users first
    this.onUsersAvailable();
    // compute available users for picklist
    const avalUsers = this.sourceUsers.filter(user => {
      return (roleBindingData.spec.users.indexOf(user.name) < 0);
    });
    this.sourceUsers = avalUsers;
    return new AuthRoleBinding(roleBindingData);
  }

  /**
   * This API serves HTML template
   */
  isAllInputsValidated() {
    const hasFormGroupError = Utility.getAllFormgroupErrors(this.newAuthRolebinding.$formGroup);
    return (hasFormGroupError === null);
  }

  /**
   * This API serves HTML template
   */
  onCancelAddRolebinding($event) {
    this.newAuthRolebinding.$formGroup.reset();
    this.formClose.emit(false);
  }

  /**
   * This API serves HTML template
   * User clicks [Save] button
   */
  onSaveAddRolebinding($event) {
    const errors = Utility.getAllFormgroupErrors(this.newAuthRolebinding.$formGroup);
    if (errors === null) {
      if (this.isEditMode()) {
        this.updateRolebinding();
      } else {
        if (this.newAuthRolebinding.getFormGroupValues().spec.role === null) {
          // Make sure user select a role
          this._controllerService.invokeErrorToaster('Invalid Input', 'Please specify role');
        } else {
          this.addRoleBinding();
        }
      }
    }
  }

  /**
   * Update a role-binding
   */
  updateRolebinding() {
    const newRolebinding = Utility.TrimDefaultsAndEmptyFields(this.getRolebindingFromUI());
    this._authService.UpdateRoleBinding(newRolebinding.meta.name, newRolebinding).subscribe(
      response => {
        this._controllerService.invokeSuccessToaster(Utility.UPDATE_SUCCESS_SUMMARY, 'Update Rolebinding ' + newRolebinding.meta.name);
        this.formClose.emit(true);
      },
      this._controllerService.restErrorHandler(Utility.UPDATE_FAILED_SUMMARY)
    );
  }

  /**
   * Add a role-binding
   */
  addRoleBinding() {
    const newRolebinding = Utility.TrimDefaultsAndEmptyFields(this.getRolebindingFromUI());
    this._authService.AddRoleBinding(newRolebinding).subscribe(
      response => {
        this._controllerService.invokeSuccessToaster(Utility.CREATE_SUCCESS_SUMMARY, 'Created Rolebinding ' + newRolebinding.meta.name);
        this.formClose.emit(true);
      },
      this._controllerService.restErrorHandler(Utility.CREATE_FAILED_SUMMARY)
    );
  }

  getRolebindingFromUI(): any {
    const _myRolebinding = this.newAuthRolebinding.getFormGroupValues();
    _myRolebinding.spec.users.length = 0;
    this.targetUsers.forEach(user => {
      _myRolebinding.spec.users.push(user.name);
    });
    return _myRolebinding;
  }

}
