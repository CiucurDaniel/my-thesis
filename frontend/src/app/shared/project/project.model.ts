import {UserModel} from "../../user/user.model";

export interface ProjectModel {
  Id: string;
  Name: string;
  Teacher: UserModel;
  Student: UserModel;
  Description: string;
  Tags: string[];
}
