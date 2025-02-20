package api

import fiber "github.com/gofiber/fiber/v2"

type Route struct {
	Method  string
	Path    string
	Handler func(ctx *fiber.Ctx) error
}

func Router(controller ControllerInterface) []Route {
	routes := make([]Route, 0)
	routes = append(routes, Route{
		Method:  "POST",
		Path:    "deck",
		Handler: controller.CreateDeck,
	})
	routes = append(routes, Route{
		Method:  "GET",
		Path:    "/deck/{id}",
		Handler: controller.GetDeck,
	})
	routes = append(routes, Route{
		Method:  "GET",
		Path:		"/decks",
		Handler: controller.ListDecks,
	})
	
	return routes
}
