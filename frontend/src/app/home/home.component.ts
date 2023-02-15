import { Component, OnInit } from '@angular/core';
import {AuthenticationService} from "../_services/auth.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  userRole: string | null = this.authenticationService.getRole();
  constructor(private authenticationService: AuthenticationService) { }

  ngOnInit(): void {
    // this.authenticationService.getRole();
    // this won't work, because this is called after the components have
    // already been initialized
  }

}
