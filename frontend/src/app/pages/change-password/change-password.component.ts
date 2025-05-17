import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-change-password',
  imports: [FormsModule, CommonModule],
  templateUrl: './change-password.component.html',
  styleUrl: './change-password.component.scss'
})
export class ChangePasswordComponent {
  userInput = {
    oldPassword: '',
    newPassword: '',
  };

  errorMessage = '';

  constructor(private authService: AuthService, private router: Router) {}

  changePassword() {
    this.authService.changePassword(this.userInput.oldPassword, this.userInput.newPassword)
      .subscribe({
        next: () => {         
          this.router.navigate(['/']);
        },
        error: (err) => {
          if (err.status === 401) {
            this.errorMessage = 'Invalid old password';
          } else if (err.status === 400) {
            this.errorMessage = 'New password is required';
          } else {
            this.errorMessage = 'An error occurred while changing the password';
          }
        }
      });
  }
}
