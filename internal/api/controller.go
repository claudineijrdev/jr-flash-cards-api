package api

import (
	fiber "github.com/gofiber/fiber/v2"
)

type ControllerInterface interface {
	CreateDeck(ctx *fiber.Ctx) error
	ListDecks(ctx *fiber.Ctx) error
	GetDeck(ctx *fiber.Ctx) error
}

type Controller struct {
	service struct {}
}

func NewController(service struct{}) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateDeck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("started")
}
