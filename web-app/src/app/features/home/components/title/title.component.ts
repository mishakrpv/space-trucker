import { Component, Input } from '@angular/core';

@Component({
  selector: 'home-title',
  imports: [],
  templateUrl: './title.component.html',
  styleUrl: './title.component.css',
})
export class TitleComponent {
  @Input({ alias: 'primary', required: true }) primaryTitle!: string;
  @Input({ alias: 'secondary', required: true }) secondaryTitle!: string;
}
