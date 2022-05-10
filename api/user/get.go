package user

import (
	"encoding/json"
	"net/http"
)

func userGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content/Type", "application/json")
	JsonEncoder := json.NewEncoder(w)
	JsonEncoder.Encode(db)
}
