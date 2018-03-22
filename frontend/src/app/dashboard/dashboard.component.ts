import { Component, OnInit } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import { Router } from '@angular/router';
import 'rxjs/add/operator/map';

var img;
var newInput;
var newimage;

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})



export class DashboardComponent implements OnInit {

  public imageDom = document.getElementById("image");
  public b = document.getElementById("list");

  constructor(private _http: Http, private router:Router) {
 this.getEncodeImage();
}

  ngOnInit() {


  }

  encode() {

      console.log("hello")
       newInput="<p>Item: <input type="+"text"+" name="+"item"+"/></p>";
       newimage= img;
       document.getElementById("image").style.backgroundColor = "red";
       document.getElementById("image").innerHTML=newimage;
  }

  private getEncodeImage() {
    return this._http.get('http://localhost:8000/encode')
                .map((res: Response) => res.json())
                 .subscribe(image => {
                   img = image
                });
  }

}
