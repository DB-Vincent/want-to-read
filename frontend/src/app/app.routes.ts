import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { BooksComponent } from './pages/books/books.component';
import { AuthGuard } from './guards/auth/auth.guard';
import { AdminGuard } from './guards/auth/admin.guard';
import { AdminComponent } from './pages/admin/admin.component';
import { AddUserComponent } from './pages/admin/add-user/add-user.component';

export const routes: Routes = [
  {
    path: '',
    component: BooksComponent,
    canActivate: [AuthGuard],
  },
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'admin',
    component: AdminComponent,
    canActivate: [AuthGuard, AdminGuard],
  },
  {
    path: 'admin/add-user',
    component: AddUserComponent,
    canActivate: [AuthGuard, AdminGuard],
  }
];
