package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

var Todos = []Todo{
	{
		Id:        1,
		Title:     "Hello World",
		Completed: false,
	},
	{
		Id:        2,
		Title:     "Fixed the wardbrode",
		Completed: false,
	},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"todos": Todos})
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		todoId, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid todo_id provided"})
		}

		var todo Todo

		for i := range Todos {
			if Todos[i].Id == todoId {
				todo = Todos[i]
			}
		}

		if (todo == Todo{}) {
			return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
		}

		return c.JSON(fiber.Map{"todo": todo})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		todo := new(Todo)

		c.BodyParser(todo)

		if todo.Title == "" {
			return c.Status(400).JSON(fiber.Map{"message": "Todo title cannot be empty"})
		}

		newTodo := Todo{Title: todo.Title, Completed: false, Id: len(Todos) + 1}

		Todos = append(Todos, newTodo)

		return c.JSON(fiber.Map{"message": "Todo successfully created", "todo": newTodo})

	})

	app.Patch("/complete/:id", func(c *fiber.Ctx) error {
		todo_id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Invalid todo_id"})
		}

		var todo *Todo

		for i := range Todos {
			if Todos[i].Id == todo_id {
				todo = &Todos[i]
			}
		}

		if (*todo == Todo{}) {
			return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
		}

		todo.Completed = true

		return c.JSON(fiber.Map{"message": "Todo completed successfully"})
	})
	fmt.Println("Server listening on Port 3000")
	app.Listen(":3000")
}
