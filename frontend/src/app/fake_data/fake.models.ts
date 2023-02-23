import {UserModel} from "../user/user.model";

export interface FakeAssignment {
  title: string;
  isDone: boolean;
  description: string;
  comments: FakeComments[];
}

export interface FakeComments {
  author: UserModel;
  comment: string;
}


