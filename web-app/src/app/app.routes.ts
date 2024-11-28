import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';

const formatTitle = (title: string) => `${title} — Space Trucker`;

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: formatTitle('Welcome'),
  },
  {
    path: 'login',
    component: LoginComponent,
    title: formatTitle('Log in'),
  },
  { path: '**', redirectTo: '' },
];
