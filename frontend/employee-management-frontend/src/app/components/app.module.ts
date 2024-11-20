import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';  // Import FormsModule for ngModel

import { AppComponent } from './app.component';
import { EmployeeAddComponent } from '../components/employee-add/employee-add.component';  // Corrected import path
import { EmployeeListComponent } from '../components/employee-list/employee-list.component';  // Corrected import path
import { RouterModule } from '@angular/router';
import { DatePipe } from '@angular/common';

@NgModule({
  declarations: [
    AppComponent,
    EmployeeAddComponent,  // Declare EmployeeAddComponent
    EmployeeListComponent  // Declare EmployeeListComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    FormsModule,  // Include FormsModule in the imports array
    RouterModule.forRoot([  // <-- Setup routing here
      { path: '', component: EmployeeListComponent },  // Default path for employee list
      { path: 'edit-employee/:id', component: EmployeeAddComponent },  // Edit employee route
      { path: 'add-employee', component: EmployeeAddComponent }  // Add employee route
    ])
  ],
  providers: [DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule {}
