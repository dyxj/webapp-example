# webapp-example
An example of a basic CRUD web application written with **Go** and **Angular2/4** using **MongoDB**.  
  
Saw some questions on building a web application with **Go** and **Angular2/4** but couldn't find an example online, so I wrote one. Written to help people new to **Go**.  
Looking to get critique from **forum.golangbridge.org** will post link to discussion when available.

## Prerequisites ##
* Go
* mongodb  

## Documentation ##  
To run the web application : 
* Start up mongodb:- mongod
* go run main.go
    * connects to default mongodb port
    * check db/db.go for futher details

**main.go**  
Entry point of web app. Initializes and run web application (app.go).

**app/app.go**  
Application base, has functions required initialize and run web application.

**db/db.go**  
Functions for database call.

**apis/api.go**  
Base for api's, contains general api functions that are meant to be used by other api's.

**apis/item.go**  
CRUD api's related to *items* and function to specify routes.

**models/items/items.go**  
MongoDB queries for *items*.

**dist/\***  
Contains compiled Angular2/4 frontend.  

**AngularItemComponent**  
Contains Angular components and service related to items.

## Web App Demo ##
![Alt text](/webapp-example.png?at=master&fileviewer=file-view-default)