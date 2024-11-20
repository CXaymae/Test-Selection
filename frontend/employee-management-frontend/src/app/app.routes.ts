// src/app/app-route.ts
import { Routes } from '@angular/router';
import { EmployeeListComponent } from './components/employee-list/employee-list.component';
import { EmployeeAddComponent } from './components/employee-add/employee-add.component';
import { EmployeeEditComponent } from './components/employee-edit/employee-edit.component';

export const appRoutes: Routes = [
  { path: '', component: EmployeeListComponent },       // Home route
  { path: 'add', component: EmployeeAddComponent },     // Add employee route
  { path: 'edit/:id', component: EmployeeEditComponent }, // Edit employee route
  { path: '**', redirectTo: '' }                        // Wildcard route for unmatched paths
];
