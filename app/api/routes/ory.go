package routes

import (
	handlersOry "github.com/alfiejones/panda-ci/app/api/handlers/ory"
	"github.com/alfiejones/panda-ci/app/queries"
	"github.com/labstack/echo/v4"
)

func RegisterOryRoutes(e *echo.Echo, queries *queries.Queries) error {

	oryHandler, err := handlersOry.NewHandler(queries)
	if err != nil {
		return err
	}

	v1 := e.Group("/v1/ory")

	v1.POST("/after/registration", oryHandler.HandleAfterRegistration)

	return nil
}
