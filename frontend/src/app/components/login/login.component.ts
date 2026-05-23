import { Component } from "@angular/core";
import { FormsModule } from "@angular/forms";
import { AuthService } from "../../services/auth.service";
import { Router, RouterLink } from "@angular/router";

@Component({
    selector:'app-login',
    standalone: true,
    imports:[FormsModule, RouterLink],
    templateUrl:'./login.component.html',
    styleUrl: './login.component.css'
})
export class LoginComponent{
    email=''
    password=''

    constructor(
        private authService:AuthService,
        private router:Router
    ){}

    login(){
        this.authService.login({
            email:this.email,
            password:this.password
        })
        .subscribe({
            next:(res)=>{
                this.authService.saveToken(res.token);
                this.router.navigate(['/']);
            }
        });
    }
}