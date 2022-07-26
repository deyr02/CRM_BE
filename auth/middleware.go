package auth

import (
	"context"
	"net/http"

	"github.com/deyr02/bnzlcrm/graph/model"
	"github.com/deyr02/bnzlcrm/jwt"
	"github.com/gin-gonic/gin"
)

const userCtxKey = "username"

var userAuthRepo UserAuthRepository = NewUserAuthRepository()

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		method := c.Request.Method
		if method == "DELETE" {
			method = "DEL"
		}

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No token",
			})
			return
		}

		//validate jwt token
		tokenStr := authHeader
		_userTokenDto, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Invalid token",
			})
			return
		}

		_user, err := userAuthRepo.GetUserByUserName(_userTokenDto.UserName)

		if err != nil {
			if _user == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Invalid User",
				})
			}
			return
		}

		if _user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid User",
			})
			return
		}

		var isAuthorized bool = userAuthRepo.IsUserAuthorized(_user.RoleID, (*model.Operation)(&method))

		if !isAuthorized {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"Message": "Access Denied",
			})
			return
		}

		ctx := context.WithValue(c.Request.Context(), userCtxKey, _user.UserName)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "No token",
			})
			return
		}

		//validate jwt token
		tokenStr := authHeader
		_userTokenDto, err := jwt.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Invalid token",
			})
			return
		}

		_user, err := userAuthRepo.GetUserByUserName(_userTokenDto.UserName)

		if err != nil {
			if _user == nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Invalid User",
				})
			}
			return
		}

		if _user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid User",
			})
			return
		}

		_isAdmin := userAuthRepo.IsAdmin(_user.RoleID)
		if !_isAdmin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"Message": "Access Denied",
			})
			return
		}
		ctx := context.WithValue(c.Request.Context(), userCtxKey, _user.UserName)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
