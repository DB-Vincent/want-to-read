import { Component, OnInit } from '@angular/core';
import { AdminNavbarComponent } from "../../components/admin-navbar/admin-navbar.component";
import { AuthService } from '../../services/auth.service';
import { User } from '../../../types/user';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-admin',
  imports: [AdminNavbarComponent, CommonModule],
  templateUrl: './admin.component.html',
  styleUrl: './admin.component.scss'
})
export class AdminComponent implements OnInit {
  users: User[] = [];

  constructor(private authService: AuthService) {}

  ngOnInit() {
    this.listUsers()
  }

  listUsers() {
    this.authService.listUsers().subscribe(
      (response) => {
        this.users = response;
      },
      (error) => {
        console.error('Error fetching users:', error);
        // Handle error, show a message to the user, etc.
      }
    );
  }
}
