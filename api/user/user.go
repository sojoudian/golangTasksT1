package user

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UserAPI struct{}

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
}

var db = []*User{}
var nextUserID uint64
var lock sync.Mutex

func (u *UserAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userGet(w, r)
	case http.MethodPost:
		userPost(w, r)
	case http.MethodDelete:
		userDel(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unsupported method %v for %v\n", r.Method, r.URL)
		log.Printf("Unsupported method '%v' to '%v'\n", r.Method, r.URL)

	}
}
