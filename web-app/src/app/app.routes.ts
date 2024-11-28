import { Routes } from '@angular/router';
import { HomeComponent } from './components/home/home.component';

const formatTitle = (title: string) => `${title} — Space Trucker`;

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: formatTitle('Welcome'),
  },
];
