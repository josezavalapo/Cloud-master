import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule }    from '@angular/http';
import { RouterModule, Routes } from '@angular/router';


import { AppComponent } from './app.component';
import { LoginComponentComponent } from './login-component/login-component.component';
import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { DomsComponent } from './doms/doms.component';


const appRoutes: Routes = [
  { path: 'signup', component: AppComponent },
  { path: 'login',      component: LoginComponentComponent },
  { path: 'dashboard',      component: DashboardComponent },
  { path: '**', component: PageNotFoundComponent }
];


@NgModule({
  declarations: [
    AppComponent,
    LoginComponentComponent,
    PageNotFoundComponent,
    DashboardComponent,
    DomsComponent
  ],
  imports: [

    HttpModule,
    RouterModule.forRoot(
    appRoutes,
    { enableTracing: true } // <-- debugging purposes only
  ),
    BrowserModule

  ],

  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
