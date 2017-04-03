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
		total := dir.WalkDirs(p.Dir)
		fmt.Printf("%.2f GB\n", float32(total)/1e9)
	})

	fmt.Println("Inicializando servidor")
	fmt.Println(http.ListenAndServe(":8080", nil))

}
