import {UserModel} from "../../user/user.model";

export interface ProjectModel {
  Id: string;
  Name: string;
  // teacher: UserModel;
  // student: UserModel;
  Description: string;
  Tags: string[];
}
