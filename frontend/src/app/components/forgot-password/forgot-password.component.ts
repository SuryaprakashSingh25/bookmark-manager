import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule, Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { ToasterService } from '../../services/toaster.service';

@Component({
  selector: 'app-forgot-password',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterModule],
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent {
  email: string = '';
  isLoading: boolean = false;
  isSubmitted: boolean = false;

  constructor(
    private authService: AuthService,
    private toaster: ToasterService,
    private router: Router
  ) {}

  onSubmit() {
    if (!this.email) {
      this.toaster.showError('Please enter your email address');
      return;
    }

    this.isLoading = true;
    this.authService.forgotPassword(this.email).subscribe({
      next: (res: any) => {
        this.isLoading = false;
        this.isSubmitted = true;
        this.toaster.showSuccess('Password reset link has been sent to your email');
        setTimeout(() => {
          this.router.navigate(['/login']);
        }, 4000);
      },
      error: (err) => {
        this.isLoading = false;
        console.error(err);
        const serverMsg = err?.error?.error || err?.message;
        this.toaster.showError(serverMsg || 'Failed to send reset email. Please try again.');
      }
    });
  }

  goToLogin() {
    this.router.navigate(['/login']);
  }
}
