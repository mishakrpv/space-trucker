import { Component, inject } from '@angular/core';
import { IamService } from '../../services/iam/iam.service';
import { RouterOutlet } from '@angular/router';

interface FooterNavLink {
  link: string;
  text: string;
}

@Component({
  selector: 'app-home',
  imports: [RouterOutlet],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  private iam: IamService = inject(IamService);

  protected footerNavLinks: FooterNavLink[] = [
    {
      link: '/',
      text: 'About',
    },
    {
      link: '/',
      text: 'Terms of Service',
    },
    {
      link: '/',
      text: 'Privacy Policy',
    },
    {
      link: '/',
      text: 'Cookie Policy',
    },
    {
      link: '/',
      text: 'Help Center',
    },
    {
      link: '/',
      text: '© 2024 Space Trucker',
    },
  ];
}
