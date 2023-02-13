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
      .subscribe(data => console.log(data));
  }
  
  getPostUpdateListener() {
    return this.projectsUpdated.asObservable();
  }
}
