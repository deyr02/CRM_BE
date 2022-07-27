package http

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deyr02/bnzlcrm/graph"
	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	var srv *handler.Server = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "Login") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		srv.ServeHTTP(c.Writer, c.Request)

	}
}
