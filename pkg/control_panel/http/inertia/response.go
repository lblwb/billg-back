package inertia

import (
	"backend/pkg/render/inertia-fiber"
	"github.com/gofiber/fiber/v2"
)

type ResponseInertia struct {
}

func NewRespInertia() *ResponseInertia {
	return &ResponseInertia{}
}

func (ResponseInertia) RenderEngine(ctx *fiber.Ctx) *inertia.Engine {
	return ctx.Locals("renderEngine").(*inertia.Engine)
}
