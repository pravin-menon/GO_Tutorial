// package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Println("hello worlds")
	var nameOne string = "John Doe"
	const nameTwo string = "Jane Doe"
	nameThree := "Whoever Doe"
	fmt.Println(nameOne, "\n", nameTwo, "\n", nameThree)
	
	app := fiber.New()
	
	app.Get("/", func(c *fiber.Ctx) error {
			return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
		})

	log.Fatal(app.Listen(":4000"))
	
}