import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { ApiService } from './services/api.service';
import { Book } from '../types/book';
import { NavbarComponent } from './components/navbar/navbar.component';
import { BookListComponent } from './components/book-list/book-list.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterModule, NavbarComponent, BookListComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent implements OnInit {
  books: Book[] = [];
  booksLoading: boolean = true;

  constructor(private apiService: ApiService) {}

  ngOnInit() {
    this.listBooks();
  }

  private listBooks() {
    this.apiService.listBooks().subscribe({
      next: (response) => {
        this.books = response;
        this.booksLoading = false;
      },
      error: () => {
        this.books = [];
        this.booksLoading = false;
      }
    })
  }

  onDeleteBook(book: Book) {
    console.log("Deleting", book.title)
    this.apiService.deleteBook(book.id).subscribe({
      next: () => {
        this.listBooks();
      },
      error: (err) => {
        console.error('Failed to delete book:', err);
      }
    });
  }

  onMarkAsRead(book: Book) {
    console.log("Marking", book.title)
    this.apiService.markBookAsRead(book).subscribe({
      next: () => {
        this.listBooks();
      },
      error: (err) => {
        console.error('Failed to mark book as read:', err);
      }
    });
  }
}
