import { Component, Input, Output, EventEmitter } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Book } from '../../../types/book';

@Component({
  selector: 'app-book-list',
  templateUrl: './book-list.component.html',
  standalone: true,
  imports: [CommonModule]
})
export class BookListComponent {
  @Input() books: Book[] = [];
  @Output() delete = new EventEmitter<Book>();
  @Output() markAsRead = new EventEmitter<Book>();

  onDelete(book: Book) {
    this.delete.emit(book);
  }

  onMarkAsRead(book: Book) {
    this.markAsRead.emit(book);
  }
} 