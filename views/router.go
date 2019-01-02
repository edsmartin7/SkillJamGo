package views

import (
   "fmt"
   "html/template"
   "net/http"

   "github.com/gorilla/mux"
   "github.com/gorilla/sessions"
)

const (

)

var (
   store = sessions.NewCookieStore(os.Getenv("SESSION-KEY"))
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("views/logintemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
}

func CredentialHandler(w http.ResponseWriter, r *http.Request) {
   username := r.FormValue("username")
   password := r.FormValue("password")

   redirectTarget := "/"
   if username=="admin" && password=="admin" { //TODO:  Set to !="", then check creds in func
      //set session
      //
      session, err := store.Get(r, "seession-name")
      if err != nil {
         http.Error(w, err.Error(), http.StatusInteranlServerError)
         return
      }
      session.Values["foo"] = "bar"
      session.Save(r, w)
      //
      redirectTarget = "/main"
   }
   http.Redirect(w, r, redirectTarget, 302)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("views/maintemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome to the Home Page")
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

}

func StartServer() {

   router := mux.NewRouter()

   //load static assets
   fileServer := http.FileServer(http.Dir("./views/"))
   router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", fileServer))

   //routes
   router.HandleFunc("/", HomeHandler)
   router.HandleFunc("/login", LoginHandler)
   router.HandleFunc("/checkLogin", CredentialHandler)
   router.HandleFunc("/main", MainHandler)

   //start server
   http.Handle("/", router)
   http.ListenAndServe(":8080", router)

}

