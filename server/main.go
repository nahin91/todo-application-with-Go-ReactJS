package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Welcome")

	app := fiber.New() // here := helps to asign the type of the 'app' which is fiber.New() here.

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	todos := []Todo{}

	app.Get("/test", func(c *fiber.Ctx) error { //this is a callback function
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error { //this is a callback function
		todo := &Todo{}

		if err := c.BodyParser((todo)); err != nil {
			return err
		}

		todo.ID = len(todos) + 1 //here assigning the 'id' of the new Todo

		todos = append(todos, *todo) // appending new data into the 'todos' array. here '*todo' is a pointer to the [todo := &Todo{}]

		return c.JSON(todos)
	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") // geting the id from the URL parameter

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos { // i=index, t=todo;
			if t.ID == id {
				todos[i].Done = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Removes slice element at index(s) and returns new slice
	app.Delete("/api/todos/:id/delete", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") // geting the id from the URL parameter

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos { // i=index, t=todo;
			if t.ID == id {
				todos = remove(todos, i)
				break
			}
		}

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}

func remove[T any](todos []T, s int) []T {
	return append(todos[:s], todos[s+1:]...)
}
