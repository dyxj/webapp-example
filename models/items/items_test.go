package items

import (
	"fmt"
	"testing"
	"webapp-example/db"
)

func TestAddItem(t *testing.T) {
	mdb := db.GetNewDB()
	defer db.CloseDB(mdb.Session)
	iConn := NewItemConn(mdb)

	i := Item{IID: 0,
		Name:        "item1",
		Description: "item description for item1",
		ImageURL:    "asset/abc"}

	id, err := iConn.AddItem(&i)
	if err == nil {
		fmt.Println(id, err)
	} else {
		t.Fatalf("AddItem test failed: %v", err)
	}
}

func TestListItems(t *testing.T) {
	mdb := db.GetNewDB()
	cmdb := db.CopyDB(mdb)
	defer db.CloseDB(cmdb.Session)
	db.CloseDB(mdb.Session)

	iConn := NewItemConn(cmdb)

	ii, err := iConn.ListItems()
	if err == nil {
		for _, i := range ii {
			fmt.Printf("%+v\n", i)
		}
	} else {
		t.Fatalf("ListItems test failed: %v", err)
	}
}