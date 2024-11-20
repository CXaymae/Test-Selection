import { Component } from '@angular/core';
import { EmployeeService, Employee } from '../../services/employee.service'; // Correct path to your service

@Component({
  selector: 'app-employee-add',
  templateUrl: './employee-add.component.html',
  styleUrls: ['./employee-add.component.css']
})
export class EmployeeAddComponent {
  employee: Employee = {
    id: '',
    firstName: '',
    lastName: '',
    email: '',
    phone: '',
    position: '',
    department: '',
    dateOfHire: new Date()
  };

  constructor(private employeeService: EmployeeService) {}

  addEmployee(): void {
    this.employeeService.addEmployee(this.employee).subscribe({
      next: (response) => {
        console.log('Employee added successfully!', response);
        // Optionally, redirect or clear the form after successful addition
      },
      error: (error) => {
        console.error('There was an error!', error);
      }
    });
  }
}
