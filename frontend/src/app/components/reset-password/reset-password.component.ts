import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { RouterModule, Router, ActivatedRoute } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { ToasterService } from '../../services/toaster.service';

@Component({
  selector: 'app-reset-password',
  standalone: true,
  imports: [CommonModule, FormsModule, RouterModule],
  templateUrl: './reset-password.component.html',
  styleUrls: ['./reset-password.component.css']
})
export class ResetPasswordComponent implements OnInit {
  resetToken: string = '';
  newPassword: string = '';
  confirmPassword: string = '';
  isLoading: boolean = false;
  isInvalidToken: boolean = false;

  constructor(
    private authService: AuthService,
    private toaster: ToasterService,
    private router: Router,
    private route: ActivatedRoute
  ) {}

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.resetToken = params['token'] || '';
      if (!this.resetToken) {
        this.isInvalidToken = true;
        this.toaster.showError('Invalid reset link');
      }
    });
  }

  onSubmit() {
    if (!this.newPassword || !this.confirmPassword) {
      this.toaster.showError('Please fill in all fields');
      return;
    }

    if (this.newPassword !== this.confirmPassword) {
      this.toaster.showError('Passwords do not match');
      return;
    }

    if (this.newPassword.length < 6) {
      this.toaster.showError('Password must be at least 6 characters');
      return;
    }

    this.isLoading = true;
    this.authService.resetPassword(this.resetToken, this.newPassword).subscribe({
      next: (res:any) => {
        this.isLoading = false;
        console.log('reset-password response:', res);
        this.toaster.showSuccess('Password reset successfully');
        setTimeout(() => {
          this.router.navigate(['/login']);
        }, 2000);
      },
      error: (err) => {
        this.isLoading = false;
        console.error('reset-password error:', err);
        const serverMsg = err?.error?.error || err?.error?.message || err?.message;
        this.toaster.showError(serverMsg || 'Failed to reset password. Link may have expired.');
      }
    });
  }

  goToLogin() {
    this.router.navigate(['/login']);
  }
}
