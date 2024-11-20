// src/app/models/employee.model.ts
export interface Employee {
    id?: string;        // Optional, for MongoDB's auto-generated ObjectId
    firstName: string;  // First name of the employee
    lastName: string;   // Last name of the employee
    email: string;      // Email address of the employee
    phone: string;      // Phone number of the employee
    position: string;   // Job position of the employee
    department: string; // Department where the employee works
    dateOfHire: string; // Date of hire, should be formatted as a string
  }
  