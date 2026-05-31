import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { AuthService } from "../../services/auth.service";
import { Router, RouterLink } from "@angular/router";
import { ToasterService } from "../../services/toaster.service";

@Component({
    selector:'app-login',
    standalone: true,
    imports:[FormsModule, RouterLink],
    templateUrl:'./login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent{
    email=''
    password=''

    constructor(
        private authService:AuthService,
        private router:Router,
        private toaster: ToasterService
    ){}

    login(){
        this.authService.login({
            email:this.email,
            password:this.password
        })
        .subscribe({
            next:(res)=>{
                this.authService.saveToken(res.token);
                this.router.navigate(['/create']);
            },
            error: (err) => {
                console.error(err);
                this.toaster.showError('Login failed. Check your credentials and try again.');
            }
        });
    }
}