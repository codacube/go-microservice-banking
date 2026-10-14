package main

import (
	"github.com/codacube/go-microservice-banking/app"
)

// To test "curl -i http://localhost:8000/greet"

/* To test and send a header

curl.exe -i -H "Content-Type: application/xml" http://localhost:8000/customers
curl.exe -i -H "Content-Type: application/json" http://localhost:8000/customers

*/
func main() {
	app.Start()
}
