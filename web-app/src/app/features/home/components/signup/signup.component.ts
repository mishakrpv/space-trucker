import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { TitleComponent } from '../title/title.component';

@Component({
  selector: 'app-signup',
  imports: [ReactiveFormsModule, TitleComponent],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
})
export class SignupComponent {
  form = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  submitForm() {}
}
