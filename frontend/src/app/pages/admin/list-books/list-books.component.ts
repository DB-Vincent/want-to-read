import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { BookListComponent } from '../../../components/book-list/book-list.component';
import { Book } from '../../../../types/book';
import { BookService } from '../../../services/books.service';

@Component({
  selector: 'app-list-books',
  imports: [BookListComponent],
  templateUrl: './list-books.component.html',
  styleUrl: './list-books.component.scss'
})
export class ListBooksComponent implements OnInit {
  userId: number | null = null;
  books: Book[] = []

  constructor(private route: ActivatedRoute, private bookService: BookService) {}

  ngOnInit() {
    this.userId = Number(this.route.snapshot.paramMap.get("user_id"))

    this.bookService.listBooks(this.userId).subscribe({
      next: (response) => {
        this.books = response;
      },
      error: () => {
        this.books = [];
      },
    });
  }
}
