package routes

import (
	"giftjob-backend/database"
	"giftjob-backend/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{os.Getenv("FRONTEND_URL"), "http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders:     []string{"Authorization", "Content-Type"},
			AllowCredentials: true,
		}),
	)
	//e.Use(customMiddleware.JWTAuthentication)

	e.POST("/query", graphqlHandler())
	e.GET("/", playgroundHandler())

	return e
}

func graphqlHandler() echo.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: database.DB}}))

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func playgroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
