// employee-list.component.ts
import { Component, OnInit } from '@angular/core';

import { EmployeeService } from '../../services/employee.service';  // Adjusted import path


@Component({
  selector: 'app-employee-list',
  templateUrl: './employee-list.component.html',
  styleUrls: ['./employee-list.component.css']
})
export class EmployeeListComponent implements OnInit {
  employees: any[] = [];

  constructor(private employeeService: EmployeeService) {}

  ngOnInit(): void {
    this.loadEmployees();
  }

  loadEmployees(): void {
    this.employeeService.getEmployees()
      .subscribe({
        next: (response) => {
          this.employees = response;
        },
        error: (error) => {
          console.error('Error fetching employees:', error);
        }
      });
  }

  deleteEmployee(id: number): void {
    this.employeeService.deleteEmployee(id)
      .subscribe({
        next: () => {
          this.loadEmployees();
        },
        error: (error) => {
          console.error('Error deleting employee:', error);
        }
      });
  }
}
