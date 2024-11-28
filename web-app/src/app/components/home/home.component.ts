import { Component, inject } from '@angular/core';
import { IamService } from '../../services/iam/iam.service';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'app-home',
  imports: [RouterOutlet],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  private iam: IamService = inject(IamService);
}
