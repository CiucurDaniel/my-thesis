import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { AuthenticationClient } from '../_clients/authentication.client';

@Injectable({
  providedIn: 'root',
})
export class AuthenticationService {
  private tokenKey = 'token';
  private role = 'role';

  constructor(
    private authenticationClient: AuthenticationClient,
    private router: Router
  ) {}

  public login(username: string, password: string): void {
    this.authenticationClient.login(username, password).subscribe((token) => {
      localStorage.setItem(this.tokenKey, token);
      //localStorage.setItem(this.role, 'student')
      localStorage.setItem(this.role, 'teacher')
      this.router.navigate(['/']);
    });
  }

  public register(username: string, email: string, password: string): void {
    this.authenticationClient
      .register(username, email, password)
      .subscribe((token) => {
        localStorage.setItem(this.tokenKey, token);
        //localStorage.setItem(this.role, 'student')
        this.router.navigate(['/']);
      });
  }

  public logout() {
    localStorage.removeItem(this.tokenKey);
    localStorage.removeItem(this.role)
    this.router.navigate(['/login']);
  }

  public isLoggedIn(): boolean {
    let token = localStorage.getItem(this.tokenKey);
    return token != null && token.length > 0;
  }

  public getToken(): string | null {
    return this.isLoggedIn() ? localStorage.getItem(this.tokenKey) : null;
  }

  public getRole(): string | null {
    return this.isLoggedIn() ? localStorage.getItem(this.role) : null;
  }
}
