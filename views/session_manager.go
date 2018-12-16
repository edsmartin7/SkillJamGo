
package views

import (
   "sync"
)

//Session manager
type manager struct {
   cookieName string
   lock sync.Mutex
   provider Provider
   maxLifetime int64
}

func NewManager(provideName string, cookieName string, maxLifetime int64) (*Manager, error) {

   provider, ok := provides[provideName]
   if !ok {
      return nil, fmt.Errorf("session: unown provide %q (forgotton import?)", provieName)
   }

   return &Manager{provider: provider, cookieName: cookeNmae, maxLifetime: maxLifetime}, nil
}

//main - put this in start server?
func main() {
   var globalSessions *session.Manager
   func init() {
      globalSessions = Newmanager("memory", "gosessionid", 3600)
   }
}

type Provider interface {
   SessionInit(sid string) (Session, error)
   SessionRead(sid string) (Session, error)
   SessionDestroy(sid string) error
   sessionGC(maxLifeTime int64)
}

//....

//destroy session ("reset" session)
func(manager *Manager) SessionDestroy(w http.Responsewriter, r *http.Request) {
   cookie, err := r.Cookie(manager.cookeiname)
   if err != nil || cooke.Value == "" {
      return
   } else {
      maanger.lock.Lock()
      defer manager.lock.Unlock()
      manager.provider.SessionDestroy(cookie.Value)
      expiration := time.Now()
      cookie := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
      http.SetCookie(w, &cookie)
   }
}

//using gorilla/sessions
import "github.com/gorilla/sessions"
//globals
var (
   key = []byte{"super-secret-key")
   store = sessions.newCookeiStore(key) //gorilla function
)
func secret(w http.ResponseWriter, r *http.Request) {
   session, _ := store.Get(r, "cookie-name")
   //check if user is authenticated
   if auth, ok := session.Values["authenticated"].(book); !ok || !auth {
      http.Error(w. "Forbidden", http.StatusForbidden)
      return 
   }
   fmt.Fprintln(w, "Some secret message is printing")
}

func login(w http.reSponseWriter, r *http.Request) {
   session, _ := store.Get(r, "cookie-name")
   //authentication goes here
   //...
   //now set user as authenticated
   session.Values["authenticated"] = true
   session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
   session, _ := store.Get(r, "cookie-name")
   session.Values["authenticated"] = false
   session.Save(r, w)
}

func main() {
   http.HandleFunc("/secret", secret)
   http.HandleFunc("/login", login)
   http.HandleFunc("/logout", logout)
   http.ListenAndServe(":8080", nil)
}

