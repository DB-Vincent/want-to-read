import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) { }

  checkHealth(): Observable<{ status: string }> {
    return this.http.get<{ status: string }>(`${this.apiUrl}/health`);
  }
} 