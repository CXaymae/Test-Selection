// src/app/services/employee.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Employee } from '../Model/employee.model';

@Injectable({
  providedIn: 'root'
})
export class EmployeeService {
  private apiUrl = 'http://localhost:3000/employees'; // Backend API URL

  constructor(private http: HttpClient) {}

  getEmployees(): Observable<Employee[]> {
    return this.http.get<Employee[]>(this.apiUrl); // Fetch employee data from backend
  }

  deleteEmployee(id: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`); // Delete employee by ID
  }
}
