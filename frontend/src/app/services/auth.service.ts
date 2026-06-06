import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { HttpClient } from "@angular/common/http";

@Injectable({
    providedIn:'root'
})
export class AuthService{
    private apiUrl=environment.apiUrl;

    constructor(private http:HttpClient){}

    signup(data:any){
        return this.http.post(
            `${this.apiUrl}/signup`,
            data
        );
    }

    login(data:any){
        return this.http.post<any>(
            `${this.apiUrl}/login`,
            data
        );
    }

    forgotPassword(email:string){
        return this.http.post(
            `${this.apiUrl}/forgot-password`,
            { email }
        );
    }

    resetPassword(token:string, newPassword:string){
        return this.http.post(
            `${this.apiUrl}/reset-password`,
            { token, newPassword }
        );
    }

    saveToken(token:string){
        localStorage.setItem('token',token);
    }

    getToken(){
        return localStorage.getItem('token');
    }

    logout(){
        localStorage.removeItem('token');
    }

    isLoggedIn(){
        return !!this.getToken();
    }
}