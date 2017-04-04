package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/sergioggonzalez/godu/dir"
)

type ReqParam struct {
	Dir string
}


func main() {

	http.HandleFunc("/size", func(rw http.ResponseWriter, req *http.Request) {
		var p ReqParam
		json.NewDecoder(req.Body).Decode(&p)
		dir.WalkDirs(p.Dir)
	})

	fmt.Println("Inicializando servidor")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
