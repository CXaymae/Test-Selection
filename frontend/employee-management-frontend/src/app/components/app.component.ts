import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template: `
    <header>
      <h1>Employee Management System</h1>
    </header>
    <div class="container">
    
    </div>
  `,
  styles: [`
    header {
      background-color: #007bff;
      color: white;
      padding: 10px;
      text-align: center;
    }
    .container {
      margin: 20px auto;
      max-width: 900px;
    }
  `]
})
export class AppComponent {}
