import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { SigninComponent } from './components/signin/signin.component';
import { SignupComponent } from './components/signup/signup.component';

const formatTitle = (title: string) => `${title} — Space Trucker`;

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: formatTitle('Welcome'),
    children: [
      {
        path: 'signin',
        component: SigninComponent,
        title: formatTitle('Sign in'),
      },
      {
        path: 'signup',
        component: SignupComponent,
        title: formatTitle('Sign up'),
      },
    ],
  },
  { path: '**', redirectTo: '' },
];
