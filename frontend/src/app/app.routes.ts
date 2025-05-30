import { Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { BooksComponent } from './pages/books/books.component';
import { AuthGuard } from './guards/auth/auth.guard';
import { AdminGuard } from './guards/auth/admin.guard';
import { AdminComponent } from './pages/admin/admin.component';
import { AddUserComponent } from './pages/admin/add-user/add-user.component';
import { ListUsersComponent } from './pages/admin/list-users/list-users.component';
import { ChangePasswordComponent } from './pages/change-password/change-password.component';
import { ListBooksComponent } from './pages/admin/list-books/list-books.component';

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
  },
  {
    path: 'admin/list-users',
    component: ListUsersComponent,
    canActivate: [AuthGuard, AdminGuard],
  },
  {
    path: 'admin/list-books/:user_id',
    component: ListBooksComponent,
    canActivate: [AuthGuard, AdminGuard],
  },
  {
    path: 'change-password',
    component: ChangePasswordComponent,
    canActivate: [AuthGuard],
  },
];
