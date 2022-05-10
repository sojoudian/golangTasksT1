package user

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func userDel(w http.ResponseWriter, r *http.Request) {

	//get user ID from URL
	fmt.Println("r.URL.String() is :", r.URL.String())
	fields := strings.Split(r.URL.String(), "/")
	id, err := strconv.ParseUint(fields[len(fields)-1], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Request to delete user with id: %v", id)

	// start code protection
	lock.Lock()
	var tmp = []*User{} // create a temporary database
	for _, u := range db {
		if id == u.ID {
			continue
		}
		tmp = append(tmp, u)
	}
	db = tmp
	// end code protection
	lock.Unlock()
}
