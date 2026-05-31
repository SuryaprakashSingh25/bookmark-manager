import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { AuthService } from "../../services/auth.service";
import { Router, RouterLink } from "@angular/router";
import { ToasterService } from "../../services/toaster.service";

@Component({
    selector:'app-signup',
    standalone: true,
    imports:[FormsModule, RouterLink],
    templateUrl:'./signup.component.html',
    styleUrls: ['./signup.component.css']
})
export class SignupComponent {
    name = ''
    email = ''
    password = ''

    constructor(
        private authService: AuthService,
        private router: Router,
        private toaster: ToasterService
    ) {}

    signup() {
        this.authService.signup({
            name: this.name,
            email: this.email,
            password: this.password
        })
        .subscribe({
            next: () => {
                this.router.navigate(['/login']);
            },
            error: (err) => {
                console.error(err);
                this.toaster.showError('Signup failed. Please check your information and try again.');
            }
        });
    }
}