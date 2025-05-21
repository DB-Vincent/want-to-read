import { Component } from '@angular/core';
import { AdminNavbarComponent } from '../../../components/admin-navbar/admin-navbar.component';
import { User } from '../../../../types/user';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../../services/auth.service';

@Component({
  selector: 'app-add-user',
  imports: [AdminNavbarComponent, FormsModule],
  templateUrl: './add-user.component.html',
  styleUrl: './add-user.component.scss',
})
export class AddUserComponent {
  userInput: User = {
    username: '',
    password: '',
    is_super: false /*  */,
  };

  constructor(private authService: AuthService) {}

  register() {
    this.authService.register(this.userInput).subscribe(
      (response) => {
        console.log('User registered successfully:', response);
        // Optionally, you can reset the form or navigate to another page
        this.userInput = {
          username: '',
          password: '',
          is_super: false,
        };
      },
      (error) => {
        console.error('Error registering user:', error);
        // Handle error, show a message to the user, etc.
      }
    );
  }
}
