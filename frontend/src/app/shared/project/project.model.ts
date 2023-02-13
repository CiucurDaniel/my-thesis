import {UserModel} from "../../user/user.model";

export interface ProjectModel {
  id: string;
  name: string;
  teacher: UserModel;
  student: UserModel;
  description: string;
  tags: string[];
}
