package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//UserDetails structure
type UserDetails struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	TotalPages int    `json:"total_pages"`
	Users      []User `json:"data"`
}

//User structure
type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

func main() {
	var userDetails UserDetails
	resp, _ := http.Get("https://reqres.in/api/users?page=2")
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &userDetails)
	users := userDetails.Users
	for _, usr := range users {
		fmt.Println(usr.Email)
	}
	resp.Body.Close()
}
