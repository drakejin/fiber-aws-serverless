package todo

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Service) FiberHandler_PostTodo(ctx *fiber.Ctx) error {
	in := &In{}
	err := ctx.BodyParser(in)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}
	err = validator.New().Struct(in.Todo)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	todo, err := s.Insert(in.Todo)
	if err != nil {
		if err != nil {
			ctx.Status(fiber.StatusConflict)
		} else {
			ctx.Status(fiber.StatusInternalServerError)
		}
		return err
	}

	out := &Out{Todo: todo}
	if b, err := out.JSON(); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	} else {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		ctx.Status(fiber.StatusOK)
		ctx.Send(b)
	}

	return nil
}

func (s *Service) FiberHandler_GetTodos(ctx *fiber.Ctx) error {
	todos, err := s.Select()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(fiber.StatusNoContent)
		} else {
			ctx.Status(fiber.StatusInternalServerError)
		}
		return err
	}

	out := &Out{Todos: todos}
	if b, err := out.JSON(); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	} else {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		ctx.Status(fiber.StatusOK)
		ctx.Send(b)
	}

	return nil
}

func (s *Service) FiberHandler_GetTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	todo, err := s.SelectOne(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(fiber.StatusNotFound)
		} else {
			ctx.Status(fiber.StatusInternalServerError)
		}
		return err
	}

	out := &Out{Todo: todo}
	if b, err := out.JSON(); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	} else {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		ctx.Status(fiber.StatusOK)
		ctx.Send(b)
	}

	return nil
}

func (s *Service) FiberHandler_PatchTodo(ctx *fiber.Ctx) error {
	in := &In{}
	err := ctx.BodyParser(in)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}

	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return err
	}
	id := ctx.Params("id", "")

	todo, err := s.Update(id, in.Todo)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(fiber.StatusNotFound)
		} else {
			ctx.Status(fiber.StatusInternalServerError)
		}
		return err
	}

	out := &Out{Todo: todo}
	if b, err := out.JSON(); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	} else {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		ctx.Status(fiber.StatusOK)
		ctx.Send(b)
	}

	return nil
}

func (s *Service) FiberHandler_DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	todo, err := s.Delete(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.Status(fiber.StatusNotFound)
		} else {
			ctx.Status(fiber.StatusInternalServerError)
		}
		return err
	}

	out := &Out{Todo: todo}
	if b, err := out.JSON(); err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	} else {
		ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
		ctx.Status(fiber.StatusOK)
		ctx.Send(b)
	}

	return nil
}
