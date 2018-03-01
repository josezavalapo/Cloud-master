import { Component } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import 'rxjs/add/operator/map';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'EC3: Equipo chingón';
  users: any = null;

   constructor(private _http: Http) {
  this.getUsers();
}

private getUsers() {
  return this._http.get('http://localhost:8000/info')
              .map((res: Response) => res.json())
               .subscribe(users => {
                      this.users = users;
                      console.log(this.users);
              });
}

}
