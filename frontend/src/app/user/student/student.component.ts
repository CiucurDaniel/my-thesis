import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-student',
  templateUrl: './student.component.html',
  styleUrls: ['./student.component.css']
})
export class StudentComponent implements OnInit {

  hasAssignedProject: boolean = true;
  constructor() { }

  ngOnInit(): void {
  }

  // todo: call service to check if project is assigned
  // or use jwt information

  // TODO: For now create 3 users and for each hard code a value.

}
