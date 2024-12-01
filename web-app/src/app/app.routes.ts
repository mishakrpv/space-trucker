import { Routes } from '@angular/router';
import { HomeComponent } from './features/home/components/home/home.component';
import { SigninOptionsComponent } from './features/home/components/signin-options/signin-options.component';
import { PageNotFoundComponent } from './features/page-not-found/components/page-not-found/page-not-found.component';
import { SignupComponent } from './features/home/components/signup/signup.component';

const formatTitle = (title: string) => `${title} — Space Trucker`;

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: formatTitle('The secure way'),
    children: [
      {
        path: '',
        component: SigninOptionsComponent,
      },
      {
        path: 'signup',
        component: SignupComponent,
      },
    ],
  },
  { path: '**', component: PageNotFoundComponent },
];
