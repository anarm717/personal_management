//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     Security:
//     - api_key:
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Authorization
//          in: header
// swagger:meta
package main

import (
	"personal_management/api"
)

func main() {
	api.Run()
}
