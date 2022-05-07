package main

import (
	"context"
	"fmt"

	"github.com/ilovejs/autodemoapi/client"
)

func main() {
	s := client.New()

	// postman.
	// 1. drag swagger json spec to create collection.
	// 2. create mock server (public) base on collection.
	s.BaseURI = "https://ed02420f-f462-4b7e-a0d1-7398c0c4c9cf.mock.pstmn.io"

	ctx := context.Background()

	res, err := s.GetDog(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(*res.Name, *res.WagsTailWhenHappy)

	// Assume that postman can convert things to Swagger docs..
	return
}
