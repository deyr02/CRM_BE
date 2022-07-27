package http

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deyr02/bnzlcrm/graph"
	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/gin-gonic/gin"
)

func GetAllUserRoleHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetAllUserRole") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GetUserRoleByID() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetUserRoleByID") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AddNewUserRole() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "AddNewRole") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ModifyUserRole() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "ModifyUserRole") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func DeleteUserRole() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "DeleteUserRole") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}
