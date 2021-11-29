package presentation

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/spa_login_learning_backend/domain"
	"github.com/pocket7878/spa_login_learning_backend/infrastructure/middleware"
)

// New sets up our routes and returns a *gin.Engine.
func NewRouter(u domain.UserUsecase, t domain.TodoUsecase) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{
				"http://localhost:3000",
				"https://spa-login-learning-frontend.vercel.app",
			},
			AllowCredentials: true,
			AllowHeaders:     []string{"Authorization", "Content-Type"},
		},
	))

	// Public endpoints
	router.Any("/", func(ctx *gin.Context) {
		response := map[string]string{
			"message": "Hello, World!",
		}
		ctx.JSON(http.StatusOK, response)
	})

	router.Any("/greeting", func(ctx *gin.Context) {
		response := map[string]string{
			"message": "Howdy?",
		}
		ctx.JSON(http.StatusOK, response)
	})

	// Require authz
	// TODO: Retrieve user real todos from usecase.
	router.GET(
		"/todos",
		middleware.EnsureValidToken(),
		TodosGet(u, t),
	)

	router.POST(
		"/todos",
		middleware.EnsureValidToken(),
		TodoPost(u, t),
	)

	return router
}
