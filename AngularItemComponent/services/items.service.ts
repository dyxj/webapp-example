import { Injectable } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import 'rxjs/Rx';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class ItemsService {

  constructor(private http: Http) { }

  getItems() {
    return this.http.get('http://localhost:8080/api/items')
      .map(
      (response: Response) => {
        const data = response.json();
        // Edit data if and when needed
        // for (const item of data){
        //   console.log(item.Name);
        // }
        return data;
      })
      .catch((error: Response) => {
        return Observable.throw('Error at getItems(): ' + error);
      });
  } // End getItems()

  addItem(item: any) {
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    return this.http.post('http://localhost:8080/api/items', item, { headers: headers })
      .map((response: Response) => {
        const data = response.json();
        return data;
      })
      .catch((error: Response) => {
        return Observable.throw('Error at addItem(): ' + error);
      });
  }// End addItem()

  deleteItem(item: any) {
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    return this.http.post('http://localhost:8080/api/items/delete', item, { headers: headers })
      .map((response: Response) => {
        const data = response.json();
        console.log(data);
        return data;
      })
      .catch((error: Response) => {
        return Observable.throw('Error at deleteItem(): ' + error);
      });
  }// End deleteItem()

  updateItem(item: any) {
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    return this.http.post('http://localhost:8080/api/items/update', item, {headers: headers})
    .map((response: Response) => {
      const data = response.json();
      console.log(data);
      return data;
    })
    .catch((error: Response) => {
      return Observable.throw('Error at updateItem(): '+ error);
    });
  } // End updateItem()

} // End ItemsService()
