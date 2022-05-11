package user

import (
	"encoding/json"
	"log"
	"net/http"
)

func userPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonDecoder := json.NewDecoder(r.Body)

	newUser := &User{}
	err := jsonDecoder.Decode(newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// start code protection and lock ID, to avoid duplicate IDs
	lock.Lock()
	nextUserID++
	newUser.ID = nextUserID
	db = append(db, newUser)
	// END code protection
	lock.Unlock()

	response := User{
		ID:        newUser.ID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Password:  newUser.Password,
	}
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(response)

}
