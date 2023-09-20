package routes

import (
	"giftjob-backend/database"
	"giftjob-backend/graph"
	customMiddleware "giftjob-backend/middleware"
	"giftjob-backend/utils"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{utils.Getenv("FRONTEND_URL"), "http://localhost:3000", "https://test-nextauth-apollo.vercel.app"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
			AllowCredentials: true,
		}),
		customMiddleware.JWEAuthentication,
	)

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
