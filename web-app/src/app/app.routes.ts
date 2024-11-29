import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';
import { SigninComponent } from './components/signin/signin.component';
import { SignupComponent } from './components/signup/signup.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';

const formatTitle = (title: string) => `${title} — Space Trucker`;

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: formatTitle('Welcome'),
    children: [
      {
        path: '',
        component: SigninComponent,
        title: formatTitle('The secure way'),
      },
      {
        path: 'signup',
        component: SignupComponent,
        title: formatTitle('Sign up'),
      },
    ],
  },
  { path: '**', component: PageNotFoundComponent },
];
