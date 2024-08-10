package response

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"task-management-system/internal/api/model"
)

func GenerateMetadata(ctx *fiber.Ctx) model.ApiMetadata {
	return model.ApiMetadata{
		Timestamp: time.Now(),
		Path:      ctx.Path(),
		Method:    ctx.Method(),
	}
}

func Response(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(model.ApiResponse{
		Success: true,
		Meta:    GenerateMetadata(ctx),
		Data:    data,
	})
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusOK, data)
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return Response(ctx, fiber.StatusCreated, data)
}
