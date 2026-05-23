import { Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { SignupComponent } from './components/signup/signup.component';
import { BookmarkList } from './components/bookmark-list/bookmark-list';
import { BookmarkForm } from './components/bookmark-form/bookmark-form';
import { authGuard } from './auth/auth.guard';

export const routes: Routes = [
  {
    path: 'login',
    component: LoginComponent,
  },
  {
    path: 'signup',
    component: SignupComponent,
  },
  {
    path: '',
    canActivate: [authGuard],
    children: [
      {
        path: '',
        component: BookmarkList,
      },
      {
        path: 'create',
        component: BookmarkForm,
      },
    ],
  },
  {
    path: '**',
    redirectTo: '',
  },
];
