import {Injectable} from "@angular/core";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {User} from "../Model/user.model";
import {catchError, map, Observable, throwError} from "rxjs";
import { jwtDecode } from 'jwt-decode';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  user: User | undefined;
  headers: HttpHeaders = new HttpHeaders();
  token: string = '';
  private apiUrl = 'http://localhost:8090/api/v1/auth';
  private userid = 'http://localhost:8090/api/v1/auth/userId';

  constructor(private http: HttpClient) {
     this.headers=this.getHeaders();
     this.token=this.getToken();
  }



  userId(): Observable<number> {
    return this.http.get<number>(`${this.userid}`, { headers: this.getHeaders() });
  }



  login(credentials: { email: string; password: string }): Observable<any> {
    const loginUrl = `${this.apiUrl}/authenticate`;

    return this.http.post<any>(loginUrl, credentials).pipe(
      map(response => {
        if (response && response.token) {
          localStorage.setItem('JwtToken', JSON.stringify(response.token));
        }
        return response;
      }),
      catchError(error => {
        return throwError(error);
      })
    );
  }

  logout(): void {
    localStorage.removeItem('JwtToken');
  }

  getToken(): string {
    return JSON.parse(localStorage.getItem('JwtToken') || 'null');
  }

  isAuthenticated(): boolean {
    return !!this.getToken();
  }


  getHeaders(): HttpHeaders {
    const token = this.getToken();
    return new HttpHeaders({
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    });
  }

  currentUser(): User | null {
    const token = this.getToken();
    if (token) {
      const tokenPayload: any = jwtDecode(token);
      if (tokenPayload && tokenPayload.roles && tokenPayload.roles.length > 0) {
        const user: User = {
          email: tokenPayload.sub,
          password: '' // Not storing password here
        };
        return user;
      }
    }
    return null;
  }


  autoLogout(dateExpiration: number): void {
    setTimeout(() => {
      this.logout();
    }, dateExpiration);
  }
}

