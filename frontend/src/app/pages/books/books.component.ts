import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { BookService } from '../../services/books.service';
import { Book } from '../../../types/book';
import { BookListComponent } from '../../components/book-list/book-list.component';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterModule, BookListComponent, FormsModule],
  templateUrl: './books.component.html',
  styleUrl: './books.component.scss',
})
export class BooksComponent implements OnInit {
  books: Book[] = [];
  booksLoading: boolean = true;
  newBook: Book = {
    title: '',
    author: '',
    completed: false,
  };

  constructor(private bookService: BookService) {}

  ngOnInit() {
    this.listBooks();
  }

  private listBooks() {
    this.bookService.listBooks().subscribe({
      next: (response) => {
        this.books = response;
        this.booksLoading = false;
      },
      error: () => {
        this.books = [];
        this.booksLoading = false;
      },
    });
  }

  onDeleteBook(book: Book) {
    const dialog = document.getElementById('delete_book') as HTMLDialogElement;
    dialog.showModal();

    const form = dialog.querySelector('form');
    form?.addEventListener('submit', (e) => {
      e.preventDefault();
      dialog.close();

      this.bookService.deleteBook(book.id!).subscribe({
        next: () => {
          this.listBooks();
        },
        error: (err) => {
          console.error('Failed to delete book:', err);
        },
      });
    });
  }

  onMarkAsRead(book: Book) {
    this.bookService.markBookAsRead(book).subscribe({
      next: () => {
        this.listBooks();
      },
      error: (err) => {
        console.error('Failed to mark book as read:', err);
      },
    });
  }

  addBook() {
    this.bookService.addBook(this.newBook).subscribe({
      next: () => {
        this.newBook = { title: '', author: '', completed: false };
        this.closeAddBookDialog();

        this.listBooks();
      },
      error: (error) => {
        console.error('Error adding book:', error);
      },
    });
  }

  closeAddBookDialog() {
    const dialog = document.getElementById('add_book') as HTMLDialogElement;
    dialog?.close();
  }

  closeDeleteBookDialog() {
    const dialog = document.getElementById('delete_book') as HTMLDialogElement;
    dialog?.close();
  }
}
