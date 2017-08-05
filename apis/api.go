package apis

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"io"
	"net/http"
)

type Api struct {
	mDB *mgo.Database
}

// InitApiRoutes : intialize api routes
func InitApiRoutes(rt *mux.Router) {
	sr := rt.PathPrefix("/api").Subrouter()
	sr.StrictSlash(true).HandleFunc("/test1", ApiCall)
	sr.StrictSlash(true).HandleFunc("/test2", ApiWrap(ApiCall))
	ItemsRoutesInit(sr)
}

// withACAO : Access-Control-Allow-Origin
func withACAO(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf(r)
		// Run before
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions{
			w.Header().Add("Access-Control-Allow-Headers", "Authorization, Content-Type")
			return
		}
		// Run input Handler Func
		h.ServeHTTP(w, r)
		// Run after
	})
}

// ApiWrap : Test function
func ApiWrap(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Run before
		io.WriteString(w, "api?\n")
		// Run input Handler Func
		h.ServeHTTP(w, r)
		// Run after
	})
}

// ApiCall : Test function
func ApiCall(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	//fmt.Println(r.Body)
	//fmt.Println("\n\n")
	//fmt.Printf("%+v\n", r)
	io.WriteString(w, "First Api: 1\n")
}