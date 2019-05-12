package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func main() {
	var data map[string]interface{}
	resp, _ := http.Get("https://reqres.in/api/users?page=2")
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &data)
	//stringBody := string(bytes)
	result := data["data"]
	fmt.Println(result)
	resp.Body.Close()
}