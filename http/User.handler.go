package http

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deyr02/bnzlcrm/graph"
	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/gin-gonic/gin"
)

func GetAllUserHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetAllUser") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GetUserByID() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetUserByID") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AddNewUser() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "AddNewUser") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ModifyUser() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "ModifyUser") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func DeleteUser() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "DeleteUser") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}
