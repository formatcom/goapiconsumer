package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func main() {

	var response map[string]interface{}

	raw, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer raw.Body.Close()

	rawBody, _ := ioutil.ReadAll(raw.Body)

	json.Unmarshal(rawBody, &response)

/* OUTPUT
{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}
id: 1, title: delectus aut autem
*/
	fmt.Println(string(rawBody))

	fmt.Printf("id: %v, title: %s\n\r",
		response["id"],
		response["title"])
}
