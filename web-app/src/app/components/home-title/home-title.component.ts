import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-home-title',
  imports: [],
  templateUrl: './home-title.component.html',
  styleUrl: './home-title.component.css',
})
export class HomeTitleComponent {
  @Input({ alias: 'primary', required: true }) primaryTitle!: string;
  @Input({ alias: 'secondary', required: true }) secondaryTitle!: string;
}
