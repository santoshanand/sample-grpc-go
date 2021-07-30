package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santoshanand/sample-grpc-go/gen"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := gen.NewUserServiceClient(conn)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		userReq := &gen.Request{
			UserId:    100,
			FirstName: "Santosh",
			LastName:  "Anand",
			Email:     "santosh@mail.com",
			Password:  "testetee",
		}
		res, err := client.GetUsers(context.Background(), userReq)
		if err != nil {
			fmt.Println("Error")
		}
		return c.JSON(res)
	})

	app.Listen(":3000")
}
