import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { TitleComponent } from '../title/title.component';

@Component({
  selector: 'app-signup',
  imports: [ReactiveFormsModule, TitleComponent],
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css', '../signin-options/signin-options.component.css'],
})
export class SignupComponent {
  form = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  submitForm() {}
}
