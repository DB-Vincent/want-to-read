import { Component, OnInit } from '@angular/core';
import { AdminNavbarComponent } from '../../../components/admin-navbar/admin-navbar.component';
import { AuthService } from '../../../services/auth.service';
import { User } from '../../../../types/user';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
@Component({
  selector: 'app-list-users',
  imports: [AdminNavbarComponent, CommonModule],
  templateUrl: './list-users.component.html',
  styleUrl: './list-users.component.scss',
})
export class ListUsersComponent implements OnInit {
  users: User[] = [];
  errorMessage: string = '';

  constructor(private authService: AuthService, private router: Router) {}

  ngOnInit() {
    this.listUsers();
  }

  listUsers() {
    this.authService.listUsers().subscribe(
      (response) => {
        this.users = response;
      },
      (error) => {
        console.error('Error fetching users:', error);
      }
    );
  }

  makeSuperUser(user: User, superUser: boolean = true) {
    const updatedUser: User = { ...user, is_super: superUser };

    this.authService.updateUser(updatedUser).subscribe(
      () => {
        this.listUsers();
      },
      (error) => {
        this.errorMessage = error.error.error;
      }
    );
    console.log('Making user a superuser:', user.id);
  }

  navigateToUsersBooks(userId: number | undefined) {
    this.router.navigate(['/admin/list-books/', userId]);
  }
}
