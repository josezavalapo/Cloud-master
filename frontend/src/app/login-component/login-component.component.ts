import { Component, OnInit } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import { Router } from '@angular/router';
import 'rxjs/add/operator/map';



@Component({
  selector: 'app-login-component',
  templateUrl: './login-component.component.html',
  styleUrls: ['./login-component.component.css']
})
export class LoginComponentComponent implements OnInit {

  public email = document.getElementById("email");
  public password = document.getElementById("password");
  users: any = null;

  constructor(private _http: Http, private router:Router) {
 this.getUsers();
}

  ngOnInit() {
  //  console.log(email.value, password)

  }

  private getUsers() {
    return this._http.get('http://localhost:8000/info')
                .map((res: Response) => res.json())
                 .subscribe(users => {
                        this.users = users;
                        console.log("Users", this.users);
                });
  }


  onClickMe() {

    for(var i = 0; i < this.users.length; i++) {
      if(this.users[i].email == this.email && this.users[i].password == this.password) {
        console.log("email", this.users[i].email, "password", this.users[i].password)
        console.log("router", this.router)
         this.router.navigate(["/dashboard"]);
        break;
      } else {
        alert("User not recognized");
      }
    }
    console.log("users", this.users)
     console.log("login", this.email)
   }


}
