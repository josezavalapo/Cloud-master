import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ImageUploadModule } from "angular2-image-upload";


import { AppComponent } from './app.component';


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
  	ImageUploadModule.forRoot(),
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
