package http

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/deyr02/bnzlcrm/graph"
	"github.com/deyr02/bnzlcrm/graph/generated"
	"github.com/gin-gonic/gin"
)

//----------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------
//------------------------------------------ Activity ----------------------------------------
//----------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------
//----------------------------------------------------------------------------------------

func GetAllAccountHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetAllAccount") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GetAccountByID() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "GetAccount") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func AddNewAccount() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "AddAccount") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ModifyAccount() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "ModifyAccount") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func DeleteAccount() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		if !CheckRequest(c, "DeleteAccount") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Request",
			})
			return
		}
		h.ServeHTTP(c.Writer, c.Request)
	}
}
