import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { ApiService } from './services/api.service';
import { Book } from '../types/book';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
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
}
