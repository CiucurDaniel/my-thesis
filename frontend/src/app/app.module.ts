import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import {HttpClientModule} from "@angular/common/http";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatButtonModule} from "@angular/material/button";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatCardModule} from "@angular/material/card";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatToolbarModule} from "@angular/material/toolbar";
import {MatProgressSpinnerModule} from '@angular/material/progress-spinner';
import {MatExpansionModule} from '@angular/material/expansion';
import { MatPaginatorModule } from "@angular/material/paginator";
import { MatSortModule } from "@angular/material/sort";
import { MatTableModule } from "@angular/material/table";
import {MatChipsModule} from '@angular/material/chips';

import { StudentComponent } from './user/student/student.component';
import { TeacherComponent } from './user/teacher/teacher.component';
import { ProjectAssignedComponent } from './user/student/project-assigned/project-assigned.component';
import { ProjectMissingComponent } from './user/student/project-missing/project-missing.component';
import { ProjectListComponent } from './shared/project/project-list/project-list.component';
import { ProjectCreateComponent } from './shared/project/project-create/project-create.component';
import { FooterComponent } from './layout/footer/footer.component';
import { HeaderComponent } from './layout/header/header.component';
import {MatSidenavModule} from "@angular/material/sidenav";


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    StudentComponent,
    TeacherComponent,
    ProjectAssignedComponent,
    ProjectMissingComponent,
    ProjectListComponent,
    ProjectCreateComponent,
    FooterComponent,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule,
    MatCardModule,
    MatInputModule,
    MatToolbarModule,
    MatProgressSpinnerModule,
    MatExpansionModule,
    MatInputModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatChipsModule,
    MatSidenavModule
  ],

  // this is for services
  //providers: [ProjectService],
  // either add services here or with @Injectable({}) on
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
