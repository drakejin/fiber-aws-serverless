package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/drakejin/fiber-aws-serverless/internal/container"
	"github.com/drakejin/fiber-aws-serverless/internal/service/todo"
)

func NewHTTP(c *container.Container) *fiber.App {
	f := fiber.New()

	f.Use(cors.New())
	f.Use(logger.New())

	f.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("ok")
	})

	todoService := todo.New(c)

	f.Post("api/v1/todos", todoService.FiberHandler_PostTodo)
	f.Get("api/v1/todos", todoService.FiberHandler_GetTodos)
	f.Get("api/v1/todos/:id", todoService.FiberHandler_GetTodo)
	f.Patch("api/v1/todos/:id", todoService.FiberHandler_PatchTodo)
	f.Delete("api/v1/todos/:id", todoService.FiberHandler_DeleteTodo)

	return f
}
