package views

import (
   "fmt"
   "html/template"
   "net/http"

   "github.com/gorilla/mux" 
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("views/logintemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome to the Home Page")
}

func StartServer() {

   router := mux.NewRouter()

   //load static assets
   fileServer := http.FileServer(http.Dir("./views/"))
   router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", fileServer))

   //routes
   router.HandleFunc("/", HomeHandler)
   router.HandleFunc("/login", LoginHandler)

   //start server
   http.Handle("/", router)
   http.ListenAndServe(":8080", router)

}

