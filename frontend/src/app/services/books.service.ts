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

  listBooks(userId?: number | null): Observable<Book[]> {
    let actualUserId

    if (!userId)
      actualUserId = this.authService.getUserId()
    else
      actualUserId = userId

    return this.http.get<Book[]>(`${this.apiUrl}/users/${actualUserId}/books`, {
      headers: this.headers,
    });
  }

  addBook(book: Book): Observable<Book> {
    return this.http.post<Book>(`${this.apiUrl}/users/${this.authService.getUserId()}/books`, book, {
      headers: this.headers,
    });
  }

  markBookAsRead(book: Book): Observable<Book> {
    return this.http.patch<Book>(
      `${this.apiUrl}/users/${this.authService.getUserId()}/books/${book.id}`,
      {
        completed: !book.completed,
      },
      {
        headers: this.headers,
      }
    );
  }

  deleteBook(id: Number): Observable<string> {
    return this.http.delete<string>(`${this.apiUrl}/users/${this.authService.getUserId()}/books/${id}`, {
      headers: this.headers,
    });
  }
}
