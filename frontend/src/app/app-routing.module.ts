import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {LoginComponent} from './login/login.component'
import {HomeComponent} from './home/home.component';
import {AuthGuard} from './_helpers/auth.guard';
import {ProjectListComponent} from "./shared/project/project-list/project-list.component";
import {HeaderComponent} from "./layout/header/header.component";
import {ProjectCreateComponent} from "./shared/project/project-create/project-create.component";

const routes: Routes = [
  {
    path: '', component: HeaderComponent, canActivate: [AuthGuard], children: [
      {path: '', component: HomeComponent},
      {path: 'projects', component: ProjectListComponent},
      {path: 'create_project', component: ProjectCreateComponent},
    ]
  },
  {path: 'login', component: LoginComponent},

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
