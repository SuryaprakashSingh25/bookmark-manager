import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ToasterService } from '../../services/toaster.service';

@Component({
  selector: 'app-toaster',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './toaster.component.html',
  styleUrls: ['./toaster.component.css'],
})
export class ToasterComponent {
  constructor(public toaster: ToasterService) {}
}
