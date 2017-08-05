package app

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"webapp-example/apis"
	"webapp-example/db"

	"github.com/dyxj/gomod"
	"github.com/gorilla/mux"
)

// Apl : Application struct
type Apl struct {
	router *mux.Router
	feRoot string
	fs     http.Handler
}

// InitApp function : Initialize application
// feRoot : refers to the root path of front end files.
func (a *Apl) InitApp(feRoot string) {
	// Initialize router
	a.router = mux.NewRouter()
	// Init root path
	a.feRoot = feRoot
	// Init static file server
	// Serve default file index.html if file not found (Angular requirement)
	a.fs = http.FileServer(gomod.DirD{Dir: a.feRoot, Def: "index.html"})
	// Init routes
	a.initializeRoutes()
}

func (a *Apl) initializeRoutes() {
	// Api sub routes
	apis.InitApiRoutes(a.router)

	// Static file server
	a.router.PathPrefix("/").Handler(a.fs)
}

// Run function : Runs application
// addr : port for listen and serve
func (a *Apl) Run(addr string) {
	fmt.Println("Connect to database")
	db.Connect_iDB()
	defer db.Close_iDB()

	fmt.Println("Server start")
	log.Fatal(http.ListenAndServe(addr, a.router))

	fmt.Println("Server end") // Should not run..
}

// Was used for debugging
func (a *Apl) staticFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("static file handler: " + r.RequestURI)
	newPath := filepath.Join(a.feRoot, r.RequestURI)
	http.ServeFile(w, r, newPath)
}
