import { Component, inject } from '@angular/core';
import { IamService } from '../../services/iam/iam.service';

@Component({
  selector: 'app-home',
  imports: [],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css',
})
export class HomeComponent {
  private iam: IamService = inject(IamService);
}
