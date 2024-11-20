# Employee Management System

## 1.Project Overview

The **Employee Management System** is a web application designed to manage employee data for organizations. This application allows administrators to view, add, edit, and delete employee information.The system provides a clean and efficient interface for managing employee records, improving organizational efficiency.

### Features:
- View a list of all employees.
- Add new employees.
- Edit existing employee details.
- Delete employee records.
- Data fetched and manipulated through a RESTful API built with Go.

## Setup Instructions

### Backend Setup (Go + MongoDB)

1. **Clone the Repository**:  
   First, clone the project repository to your local machine.

   ```bash
   git clone https://github.com/CXaymae/Test-Selection.git
   cd employee-management/backend
   ```

   ## 2.Docker Setup for Backend and Frontend
    
    #### Docker Configuration Overview
    The project is containerized using Docker to streamline deployment and ensure consistent development environments across machines. Docker images are built for both the backend (Go + MongoDB) and frontend (Angular), and these services are orchestrated using Docker Compose.

    Docker Compose allows you to define and run multi-container Docker applications. In this project, Docker Compose is used to set up and link the frontend, backend, and MongoDB database.

       ##### Backend (Go + MongoDB) Dockerfile
        In the backend/ directory, you will find the Dockerfile for the backend service. This Dockerfile defines how the Go backend application is built and run in a container.

      ###### Key steps in the Dockerfile:
      - Uses the official Golang image as the base.
      - Sets up the working directory and copies the Go modules and application files.
      - Installs dependencies and builds the Go application.
      - Exposes port 8080 for the Go application.
      - Specifies the command to run the Go application when the container starts.
 
       ##### Frontend (Angular) Dockerfile
        In the frontend/ directory, you will find the Dockerfile for the frontend service. This Dockerfile defines how the Angular application is built and served in a container.

      ###### Key steps in the Dockerfile:
      - Uses the official Node.js image to build the Angular application.
      - Installs dependencies and builds the application for production.
      - Exposes port 4200 to allow access to the Angular application.
      - Specifies the command to run the Angular application when the container starts.
      - Docker Compose Configuration
      - The docker-compose.yml file is located in the root directory of the project. This file defines the multi-container setup for the backend, frontend, and MongoDB services.

        ###### Key sections of the Docker Compose file:
      -  backend: Builds the Go backend container and exposes port 8080. It depends on MongoDB to ensure the 
         database is up and running before the backend starts.
      - mongodb: Uses the official MongoDB image and creates a volume to persist the database data.
      - frontend: Builds the Angular frontend container and exposes port 4200 for the UI.

    ```bash
    docker-compose up --build
    ```
### RESTful API Design
The backend API is designed to interact with employee data via RESTful endpoints. Here are the main API routes/ endpoints:

   - GET /employees – Retrieve a list of all employees.
   - GET /employees/{id} – Retrieve details of a specific employee.
   - POST /employees – Add a new employee.
   - PUT /employees/{id} – Update an existing employee.
   - DELETE /employees/{id} – Delete an employee.
 - 
    Example of RESTful API Requests:

    ##### GET all employees:

    ```bash
    GET http://localhost:8080/employees
    ```
    Response: JSON array of all employees.

    ##### GET employee by ID

    ```bash
    GET http://localhost:8080/employees/{id}
    ```

    ##### POST a new employee

    ```bash
    POST http://localhost:8080/employees
    
    ```


### Technologies Used:
- Backend: Go (Golang), MongoDB
- Frontend: Angular
- Testing: Go testing libraries, Karma, Jasmine (Currnet to use them)
  
### Future Enhancements:
- Add user authentication and authorization.
- Improve UI/UX with modern design components.
- Implement pagination and sorting for the employee list.
- Add search functionality to easily find employees by name, department.
- Implement intuitive dashboard

