package items

import (
	"crypto/rand"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"webapp-example/db"
)

type Item struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	IID         int64 `json:"IID,string" bson:"iid"`
	Name        string `json:"Name" bson:"name"`
	Description string `json:"Description" bson:"description"`
	ImageURL    string `json:"ImageURL" bson:"imageurl"`
}

const Clitems = "items"

type ItemConn struct {
	db *mgo.Database
	cl *mgo.Collection
}

func NewItemConn(db *mgo.Database) *ItemConn {
	var iConn = ItemConn{}
	iConn.db = db
	iConn.cl = iConn.db.C(Clitems)
	return &iConn
}

// Should not need this..
func (iConn *ItemConn) CloseItemConn() {
	db.CloseDB(iConn.db.Session)
}

// Random id generator (Temporary)
var maxRand = big.NewInt(1<<63 - 1)

func randomID() (int64, error) {
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	return n.Int64() + 1, nil
}

// Assign a new item ID, and add Item type to Items collection on database
func (iConn *ItemConn) AddItem(i *Item) (id int64, err error) {
	// Generate randomID
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("AddItem: couldn't assign a new ID: %v", err)
	}

	i.IID = id
	// Add item to items collection
	if err := iConn.cl.Insert(i); err != nil {
		return 0, fmt.Errorf("AddItem: couldn't add item: %v", err)
	}
	return id, nil
}

// Delete item from Items collection given the item ID
func (iConn *ItemConn) DeleteItem(iid int64) error {
	return iConn.cl.Remove(bson.D{{Name: "iid", Value: iid}})
}

// Delete item from Items collection given the item ID
func (iConn *ItemConn) DeleteAllItems() (*mgo.ChangeInfo,error) {
	info, err := iConn.cl.RemoveAll(bson.D{})
	return  info, err
}

// Update item  in Items collection
func (iConn *ItemConn) UpdateItem(i *Item) error {
	return iConn.cl.Update(bson.D{{Name: "iid", Value: i.IID}}, i)
}

// Returns an Item given the item ID
func (iConn *ItemConn) GetItem(id int64) (*Item, error) {
	i := &Item{}
	if err := iConn.cl.Find(bson.D{{Name: "iid", Value: id}}).One(i); err != nil {
		return nil, err
	}
	return i, nil
}

func (iConn *ItemConn) ListItems() ([]*Item, error) {
	var ii []*Item
	if err := iConn.cl.Find(nil).Sort("Name").All(&ii); err != nil {
		return nil, err
	}
	return ii, nil
}
