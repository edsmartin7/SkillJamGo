package views

import (
   "fmt"
   "html/template"
   "net/http"
   "os"

   "github.com/gorilla/mux"
   "github.com/gorilla/sessions"
)

const (

)

var (
   store = sessions.NewCookieStore([]byte(os.Getenv("SESSION-KEY")))
)

//shows the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
   t, err := template.ParseFiles("views/logintemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
}

//shows the home page after logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}

//actual login functionality
func CredentialHandler(w http.ResponseWriter, r *http.Request) {
   username := r.FormValue("username") //or r.PostForm.Get("username")
   password := r.FormValue("password")

   redirectTarget := "/"
   if username=="admin" && password=="admin" { //TODO:  Set to !="", then check creds in func
      //set session
      session, err := store.Get(r, "session-name") //session-name -> username
      //session.Values["username"] = username
      //session.Save(r,w)
      if err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
         return
      }
      session.Values["foo"] = "bar"
      session.Save(r, w)
      
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
   t, err := template.ParseFiles("views/profiletemplate.html")
   if err != nil {
      fmt.Println(err)
      return
   }

   t.Execute(w, t)
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
   router.HandleFunc("/profile", ProfileHandler)

   //start server
   http.Handle("/", router)
   http.ListenAndServe(":8080", router)

}

