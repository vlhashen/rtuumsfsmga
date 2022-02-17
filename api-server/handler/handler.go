package handler

import (
	"net/http"
  //"github.com/gorilla/sessions"
)

/*
type SessionHandler func(w http.ResponseWriter, r *http.Request, s *sessions.Session)

func (sh SessionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func PermissionMiddleware(sh SessionHandler, store *sessions.CookieStore) SessionHandler {
  return func(w http.ResponseWriter, r *http.Request) {
    session, err := store.Get(r, "r_Cookie")
    if err != nil {
      return
    }
    
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	  w.Header().Set("Access-Control-Allow-Credentials", "true")  
    sh(w, r, session)

  }
}
*/

func PermissionMiddleware(hf http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	  w.Header().Set("Access-Control-Allow-Credentials", "true")  
    hf(w, r)

  }
}


