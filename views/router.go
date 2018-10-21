package views

import (
   "fmt"
   "html/template"
   "net/http" 
)

func loginPage(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("views/logintemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
}

func StartServer() {
   mux := http.NewServeMux()

   //load static assets
   //fileServer := http.FileServer(http.Dir("./views/"))
   //mux.PathPrefix("/views/").Handler(http.StripPrefix("views/", fileServer)) //
   fs := http.FileServer(http.Dir("static"))
   http.Handle("/static/", http.StripPrefix("/static/", fs))

   //routes
   handler := http.HandlerFunc(loginPage)
   mux.Handle("/login", handler) 

   //start server
   //http.Handle("/", mux)
   http.ListenAndServe(":8080", mux)

}

