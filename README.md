# CRUD-Restful-API-in-GO
A simple rest api in Go, using mux router and net for http request. 

## Run 
You can build and run the project with one command 
`Go build && ./restapi`

## Endpoints
* GET - http://localhost:8000/api/books
* GET - http://localhost:8000/api/book/{id}
* POST - http://localhost:8000/api/book
  Request 
  ```
  {
    "Isbn": "4333",
    "Title": "We are the gods",
    "Author": {
        "Firstname": "Nans",
        "Lastname": "Adeku",
        "Email": "Nans@gmail.com"
    }
}
  ```
  
* PUT - http://localhost:8000/api/book/{id}
  Request 
  ```
  {
    "Isbn": "4333",
    "Title": "We are the gods",
    "Author": {
        "Firstname": "Nans",
        "Lastname": "Adeku",
        "Email": "Nans@gmail.com"
    }
}
  ```
* DELETE - http://localhost:8000/api/book/{id}
