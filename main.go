package main

import (
	"fmt"
	"github.com/sojoudian/t1/api/user"
	"log"
	"net/http"
)

const port = "3003"

func main() {
	http.Handle("/users/", &user.UserAPI{})
	fmt.Printf("web server started at port %v ...", port)
	log.Fatalln(http.ListenAndServe(":"+"3003", nil))

}
