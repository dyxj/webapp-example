package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"net/http"
	"webapp-example/db"
	"webapp-example/models/items"
)

func ItemsRoutesInit(rt *mux.Router) {
	rt.StrictSlash(true).Path("/items").Handler(withACAO(ItemListApi)).
		Methods(http.MethodGet)
	rt.StrictSlash(true).Path("/items").Handler(withACAO(ItemAddApi)).
		Methods(http.MethodPost, http.MethodOptions)
	// Sub Router
	sr := rt.PathPrefix("/items").Subrouter()
	sr.StrictSlash(true).HandleFunc("/delete", withACAO(ItemDeleteApi)).
		Methods(http.MethodPost, http.MethodOptions)
	sr.StrictSlash(true).HandleFunc("/update", withACAO(ItemUpdateApi)).
		Methods(http.MethodPost, http.MethodOptions)
}

func ItemListApi(w http.ResponseWriter, r *http.Request) {
	cdb := db.Copy_iDB()
	defer db.CloseDB(cdb.Session)
	iConn := items.NewItemConn(cdb)

	// Get item list from database
	ii, err := iConn.ListItems()
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//for _, i := range ii {
	//	fmt.Printf("%+v\n", i)
	//}

	js, err := json.Marshal(ii)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}

func ItemAddApi(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	var i items.Item
	err := d.Decode(&i)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check mandatory fields if required
	fmt.Printf("%v\n", i)

	// Add to DB
	cdb := db.Copy_iDB()
	defer db.CloseDB(cdb.Session)
	iConn := items.NewItemConn(cdb)
	_, err = iConn.AddItem(&i)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Should respond with some success message here (not done)
	// and the item object
	w.Header().Add("Content-Type", "application/json")
	js, err := json.Marshal(i)
	w.Write(js)
}

func ItemUpdateApi(w http.ResponseWriter, r *http.Request) {
	// Get Body
	d := json.NewDecoder(r.Body)
	var i items.Item
	err := d.Decode(&i)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// DB
	cdb := db.Copy_iDB()
	defer db.CloseDB(cdb.Session)
	iConn := items.NewItemConn(cdb)
	err = iConn.UpdateItem(&i)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write to client
	w.Header().Add("Content-Type", "application/json")
	js, err := json.Marshal(i)
	w.Write(js)
}

type itemDelInput struct {
	IID json.Number `json:"IID,omitempty"`
}

func ItemDeleteApi(w http.ResponseWriter, r *http.Request) {
	// Another way parse body
	// body, err := ioutil.ReadAll(r.Body)
	// err = json.Unmarshal(body, &i)
	var i itemDelInput
	d := json.NewDecoder(r.Body)
	err := d.Decode(&i)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cdb := db.Copy_iDB()
	defer db.CloseDB(cdb.Session)
	iConn := items.NewItemConn(cdb)
	var info *mgo.ChangeInfo
	if i.IID == "ALL" {
		info, err = iConn.DeleteAllItems()
	} else {
		iid, _ := i.IID.Int64()
		err = iConn.DeleteItem(iid)
	}
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(info)
	temp := map[string]string{"IID": i.IID.String()}
	js, err := json.Marshal(temp)
	w.Header().Add("Content-Type", "application/json")
	w.Write(js)
}
