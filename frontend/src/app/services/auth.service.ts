import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { environment } from '../../environments/environment';
import { User } from '../../types/user';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private apiUrl = environment.apiUrl;

  constructor(private http: HttpClient) {}

  login(user: User): Observable<{ token: string }> {
    return this.http.post<{ token: string }>(`${this.apiUrl}/login`, user).pipe(
      tap((response) => {
        localStorage.setItem('jwt_token', response.token);
      })
    );
  }

  register(user: User): Observable<User> {
    return this.http.post<User>(`${this.apiUrl}/register`, user, {
      headers: {
        Authorization: `Bearer ${this.getToken()}`,
      },
    });
  }

  listUsers(): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}/users`, {
      headers: {
        Authorization: `Bearer ${this.getToken()}`,
      },
    });
  }

  changePassword(oldPassword: string, newPassword: string): Observable<any> {
    return this.http.post(
      `${this.apiUrl}/change-password`,
      { old_password: oldPassword, new_password: newPassword },
      {
        headers: {
          Authorization: `Bearer ${this.getToken()}`,
        },
      }
    );
  }

  updateUser(user: User): Observable<User> {
    return this.http.patch<User>(`${this.apiUrl}/user/${user.id}`, user, {
      headers: {
        Authorization: `Bearer ${this.getToken()}`,
      },
    });
  }

  logout() {
    localStorage.removeItem('jwt_token');
  }

  getToken(): string | null {
    return localStorage.getItem('jwt_token');
  }

  getUserId(): number | null {
    const token = this.getToken();
    if (!token) return null;

    const payload = JSON.parse(atob(token.split('.')[1]));
    return payload.user_id;
  }

  isLoggedIn(): boolean {
    const token = this.getToken();
    if (!token) return false;

    const payload = JSON.parse(atob(token.split('.')[1]));
    const expirationDate = new Date(payload.exp * 1000);
    const currentDate = new Date();

    if (expirationDate < currentDate) {
      this.logout();
      return false;
    }
    
    return true;
  }

  isSuper(): boolean {
    const token = this.getToken();
    if (!token) return false;

    const payload = JSON.parse(atob(token.split('.')[1]));
    return payload.is_super === true;
  }
}
