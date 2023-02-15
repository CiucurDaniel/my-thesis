import {ProjectModel} from "../shared/project/project.model";
import {Observable, Subject} from "rxjs";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {map} from "rxjs/operators";
import {UserModel} from "../user/user.model";
import {Injectable} from "@angular/core";

@Injectable({providedIn: 'root'})
export class ProjectService {
  private projects: ProjectModel[] = [];

  private projectsUpdated = new Subject<ProjectModel[]>();

  constructor(private http: HttpClient) {
  }

  getProjects() {
    this.http.get<ProjectModel[]>(environment.apiUrl + "/project/")
      .pipe(
        map((projectData) => {
          return projectData.map(project => {
            return {
              // Field din Interfata Angular: Field din JSON trimis de backend
              Id: project.Id,
              Name: project.Name,
              Description: project.Description,
              Teacher: project.Teacher,
              Student: project.Student,
              Tags: project.Tags,
            };
          });
        }))
      .subscribe(transformedProjects => {
        this.projects = transformedProjects;
        this.projectsUpdated.next([...this.projects]);
      });
  }

  getPostUpdateListener() {
    return this.projectsUpdated.asObservable();
  }
}
