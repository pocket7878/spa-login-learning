package middleware

import (
	"net/http"
	"net/url"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/auth0/go-jwt-middleware/validate/josev2"
	"github.com/gin-gonic/gin"
	"gopkg.in/square/go-jose.v2"
)

const signatureAlgorithm = "RS256"

// EnsureValidToken is a gin.HandlerFunc middleware that will check the validity of our JWT.
func EnsureValidToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN"))

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		provider := josev2.NewCachingJWKSProvider(issuerURL, 5*time.Minute)
		validator, err := josev2.New(
			provider.KeyFunc,
			jose.RS256,
		)

		if err != nil {
			ctx.AbortWithError(500, err)
		}

		m := jwtmiddleware.New(validator.ValidateToken)

		var encounteredError = true
		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r
			ctx.Next()
		}

		m.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				map[string]string{"message": "Failed to validate JWT."},
			)
		}
	}
}
