package main

import (
	"github.com/nurana88/microservices/app"
)

func main() {

	app.StartApplication()
}

// curl -X POST localhost:8080/users -d '{"id":123, "first_name":"Nura", "email":123}'
