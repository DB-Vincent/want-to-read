import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { User } from '../../../types/user';
import { AuthService } from '../../services/auth.service';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-login',
  imports: [FormsModule, CommonModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
})
export class LoginComponent {
  errorMessage = ""
  
  constructor(private authService: AuthService, private router: Router) {}

  userInput: User = {
    username: '',
    password: '',
  };

  login() {
    this.authService.login(this.userInput).subscribe({
      next: (response) => {
        console.log('Login successful:', response);

        this.router.navigate(['/']);
      },
      error: (error) => {
        console.error('Login failed:', error);
        this.errorMessage = error.error.error;
        console.log(this.errorMessage);
      },
    });
  }
}
