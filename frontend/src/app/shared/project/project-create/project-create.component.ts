import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";

export interface TagChip {
  name: string;
}

@Component({
  selector: 'app-project-create',
  templateUrl: './project-create.component.html',
  styleUrls: ['./project-create.component.css']
})
export class ProjectCreateComponent implements OnInit {

  public projectCreateForm!: FormGroup;

  constructor() {
  }

  ngOnInit(): void {
    this.projectCreateForm = new FormGroup({
      name: new FormControl('', Validators.required),
      description: new FormControl('', Validators.required),
      tags: new FormControl('', Validators.required)
    });
  }

  public onSubmit() {
    // this.authenticationService.login(
    //   this.loginForm.get('email')!.value,
    //   this.loginForm!.get('password')!.value
    // );
  }
}
