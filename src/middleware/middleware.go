package middleware

import (
	"otumian-empire/go-ecom/src/config"
	"otumian-empire/go-ecom/src/handlers"
	"otumian-empire/go-ecom/src/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Repository[T interface{}] interface {
	FindOneById(id config.IdType) (T, error)
}

func AuthorizeJWT[T interface{}](repo Repository[T]) gin.HandlerFunc {
	return func(context *gin.Context) {
		env, _ := config.GetEnvirons()

		authToken := context.GetHeader("authorization")
		if !utils.ValidateApiKey(authToken) {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		justTheToken := strings.Split(authToken, " ")[1]
		token, err := utils.VerifyJwt(justTheToken)

		if err != nil {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization" + " " + err.Error()))
			return
		}

		if !token.Valid {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// Check if the token has expired
		if utils.IsExpiredToken(claims) {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		// Check the issuer
		if !utils.IsValidIssuer(claims, env.JwtIssuer) {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		// Get the user of this token
		userId, idErr := utils.GetIdFromClaim(claims)
		if idErr != nil {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		user, err := repo.FindOneById(userId)

		if err != nil {
			context.Abort()
			context.JSON(handlers.AuthenticationErrorResponse("Invalid authorization"))
			return
		}

		// Remove the password
		if user, ok := any(user).(map[string]interface{}); ok {
			if _, exists := user["Password"]; exists {
				delete(user, "Password")
			}
		}

		// Set the user in the Gin context
		context.Set("user", user)
		context.Next()
	}
}
