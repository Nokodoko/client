package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	name, number, err := get("string")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name, number)
}

/*
1. Get Method
2. Error and response check
3. Decode / unmarshal
4. error check
5. user defined type (out of func scope)
6. return info
*/
func get(name string) (string, int, error) {
	url := fmt.Sprintf("https://api.endpoint")

	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf(resp.StatusCode)
	}

	var user info //unused
	defer resp.Body.Close()

	//decode -> &struct (err handle)
	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&user); err != nil {
		return "", 0, err
	}

	return user.Name, user.Number, nil
}

type info struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}
