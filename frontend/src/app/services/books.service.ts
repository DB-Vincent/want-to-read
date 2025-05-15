import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Book } from '../../types/book';
import { environment } from '../../environments/environment';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class BookService {
  private apiUrl = environment.apiUrl;
  private headers;

  constructor(private http: HttpClient, private authService: AuthService) {
    this.headers = { Authorization: `Bearer ${this.authService.getToken()}` };
  }

  listBooks(): Observable<Book[]> {
    return this.http.get<Book[]>(`${this.apiUrl}/books`, {
      headers: this.headers,
    });
  }

  addBook(book: Book): Observable<Book> {
    return this.http.post<Book>(`${this.apiUrl}/book`, book, {
      headers: this.headers,
    });
  }

  markBookAsRead(book: Book): Observable<Book> {
    return this.http.patch<Book>(
      `${this.apiUrl}/book/${book.id}`,
      {
        completed: !book.completed,
      },
      {
        headers: this.headers,
      }
    );
  }

  deleteBook(id: Number): Observable<string> {
    return this.http.delete<string>(`${this.apiUrl}/book/${id}`, {
      headers: this.headers,
    });
  }
}
