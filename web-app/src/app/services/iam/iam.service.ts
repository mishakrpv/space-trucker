import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

interface Result {
  succeed: boolean;
}

@Injectable({
  providedIn: 'root',
})
export class IamService {
  private baseUrl: string = 'https://localhost:6412';

  constructor(private http: HttpClient) {}

  login(email: string, password: string) {
    this.http
      .post<Result>(`${this.baseUrl}/auth/login`, {
        email: email,
        password: password,
      })
      .subscribe((r) => {
        console.log('Login result:', r);
      });
  }

  signup() {}

  logout() {}
}
