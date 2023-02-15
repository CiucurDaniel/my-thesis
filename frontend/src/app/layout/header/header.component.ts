import { Component, OnInit } from '@angular/core';
import {AuthenticationService} from "../../_services/auth.service";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  userRole: string | null = this.authenticationService.getRole();
  isLoggedIn: boolean | null = this.authenticationService.isLoggedIn();
  constructor(private authenticationService: AuthenticationService) { }

  ngOnInit(): void {
    // this.authenticationService.getRole();
    // this won't work, because this is called after the components have
    // already been initialized
  }

  public onSubmit() {
    this.authenticationService.logout();
  }

}
