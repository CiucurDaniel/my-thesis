import {Component, OnDestroy, OnInit} from '@angular/core';
import {ProjectModel} from "../project.model";
import {Subscription} from "rxjs";
import {ProjectService} from "../../../_services/project.service";

@Component({
  selector: 'app-project-list',
  templateUrl: './project-list.component.html',
  styleUrls: ['./project-list.component.css']
})
export class ProjectListComponent implements OnInit, OnDestroy {

  projects: ProjectModel[] = [];
  isLoading = false;

  displayedColumns: string[] = ['Project Name', 'Description', 'Teacher', 'Tags', 'Status'];
  // @ts-ignore
  private projectSubscription: Subscription;
  constructor(public projectService: ProjectService) { }

  // recommended to do basic initialization tasks in onInit
  ngOnInit() {
    this.isLoading = true;
    this.projectService.getProjects();
    this.projectSubscription = this.projectService.getPostUpdateListener()
      .subscribe( (posts: ProjectModel[]) => {
        this.isLoading = false;
        this.projects = posts;
        console.log(this.projects); //DEBUG
      });
  }

  ngOnDestroy(): void {
    this.projectSubscription.unsubscribe();
  }

  onDetails(postId: string){
    // TODO: Implement me
    //this.postService.detailsPost(postId);
  }

  assignProjectToStudent(studentId: string){
    // TODO: This should actually take the ID from the localStorage
  }
}
