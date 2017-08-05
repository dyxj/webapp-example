import { Observable } from 'rxjs/Observable';
import { Component, OnInit, OnDestroy, OnChanges, DoCheck, SimpleChanges, Input } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ItemsService } from './../services/items.service';
import { Subscription } from 'rxjs/Subscription';

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent implements OnInit, OnDestroy {

  items = [];
  itemsForm: FormGroup;
  itemSub: Subscription;
  itemSubRefresh: Subscription;
  formSubmitted = false;
  // @Input() mIID: string;

  constructor(private itemService: ItemsService) { }

  ngOnInit() {
    this.itemsForm = new FormGroup({
      'IID': new FormControl({ value: null, disabled: true }),
      // 'IID': new FormControl(null),
      'Name': new FormControl(null, Validators.required),
      'Description': new FormControl(null, Validators.required),
      'ImageURL': new FormControl(null)
    });
    this.onGetItems();
    // For testing, set default value
    // this.itemsForm.setValue({ 'IID': null, 'Name': 'Hello', 'Description': 'Description hello', 'ImageURL': 'asset/hello' });
    this.itemsForm.patchValue({ 'Name': 'Hello', 'Description': 'Description hello', 'ImageURL': 'asset/hello' });
  }

  clearIID() {
    this.itemsForm.patchValue({'IID': null});
  }

  onGetItems() {
    this.itemSub = this.itemService.getItems()
      .subscribe(
      (items: any[]) => {
        if (items != null) {
          this.items = items;
        } else {
          this.items = [];
        }
        console.log(this.items);
        // "Live" update
        this.ItemsRefresh();
      },
      (error) => console.log(error)
      );
  }

  ItemsRefresh() {
    if (this.itemSubRefresh) {
      this.itemSubRefresh.unsubscribe();
    }
    this.itemSubRefresh = Observable.timer(5000).first().subscribe(
      () => this.onGetItems()
    );
  }

  onAddItem() {
    let j = JSON.stringify(this.itemsForm.getRawValue());
    // let j = JSON.stringify(this.itemsForm.value);
    // console.log(j);
    this.formSubmitted = true;
    this.itemService.addItem(j).subscribe(
      (item: any) => {
        if (item != null) {
          // console.log(this.items);
          // Add item to array
          this.items.push(item);
        }
      },
      (error) => console.log(error)
    );
  }

  onUpdItem() {
    let j = JSON.stringify(this.itemsForm.getRawValue());
    // let j = JSON.stringify(this.itemsForm.value);
    // console.log(j);
    this.formSubmitted = true;
    this.itemService.updateItem(j).subscribe(
      (item: any) => {
        for (let x = 0; x < this.items.length; x++) {
          if (this.items[x].IID == item.IID) {
            // Replace Item
            this.items[x] = item;
          }
        }
      },
      (error) => console.log(error)
    );
  }

  onClearItems() {
    this.items = [];
  }

  onDelete(delId: string) {
    const i = { IID: delId };
    const j = JSON.stringify(i);
    this.itemService.deleteItem(j).subscribe(
      (item: any) => {
        if (item.IID == 'ALL') {
          // Remove all items
          this.items = [];
          return;
        }
        for (let x = 0; x < this.items.length; x++) {
          if (this.items[x].IID == item.IID) {
            // Remove item from array
            this.items.splice(x, 1);
          }
        }
      },
      (error) => console.log(error)
    );
  }

  onEdit(item: any) {
    this.itemsForm.setValue({ 'IID': item.IID, 'Name': item.Name, 'Description': item.Description, 'ImageURL': item.ImageURL });
    // let j = JSON.stringify(this.itemsForm.value);
    console.log(this.itemsForm.get('IID').value);
    let j = JSON.stringify(this.itemsForm.getRawValue());
    console.log(j);
  }

  btnControls() {
    const x = this.itemsForm.get('IID').value;
    let b = false;
    if (x === null) {
      b = true;
    }
    return b;
  }

  btnClearForm() {
    this.itemsForm.setValue({ 'IID': null, 'Name': '', 'Description': '', 'ImageURL': '' });
  }

  ngOnDestroy() {
    console.log('Unsubscribe');
    if (this.itemSub) {
      this.itemSub.unsubscribe();
    }
    if (this.itemSubRefresh) {
      this.itemSubRefresh.unsubscribe();
    }
  } // End ngOnDestroy()

}
