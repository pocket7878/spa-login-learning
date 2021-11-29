package presentation

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/auth0/go-jwt-middleware/validate/josev2"
	"github.com/gin-gonic/gin"
	"github.com/pocket7878/spa_login_learning_backend/domain"
)

func extractProviderAndUID(ctx *gin.Context) (string, string) {
	claims := ctx.Request.Context().Value(jwtmiddleware.ContextKey{}).(*josev2.UserContext)
	subject := claims.RegisteredClaims.Subject
	subjectParts := strings.Split(subject, "|")
	subjectProvider := subjectParts[0]
	subjectUID := subjectParts[1]

	return subjectProvider, subjectUID
}

func ensureUser(ctx context.Context, u domain.UserUsecase, provider, uid string) (*domain.User, error) {
	var result *domain.User

	existsUser, err := u.GetByProviderWithUID(ctx, provider, uid)
	if err != nil {
		return nil, err
	}

	if existsUser != nil {
		return existsUser, nil
	}

	// User not exists
	result = &domain.User{
		Provider: provider,
		UID:      uid,
	}
	err = u.Store(ctx, result)
	if err != nil {
		return nil, fmt.Errorf("Failed to store new user with (%s,%s): %e", provider, uid, err)
	}

	return result, nil
}

func TodosGet(u domain.UserUsecase, t domain.TodoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Retrieve & Ensure User
		provider, uid := extractProviderAndUID(c)
		user, err := ensureUser(c, u, provider, uid)
		if err != nil {
			c.AbortWithError(500, fmt.Errorf("failed to ensure user: %w", err))
			return
		}

		todos, err := t.GetTodos(c, user.ID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		response := make([]map[string]string, 0)
		for _, t := range todos {
			response = append(response, map[string]string{
				"id":          fmt.Sprint(t.ID),
				"user_id":     fmt.Sprint(t.UserID),
				"description": t.Description,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}

type PostTodoData struct {
	Description string `json:"description"`
}

type PostTodoInput struct {
	Data PostTodoData `json:"todo"`
}

func TodoPost(u domain.UserUsecase, t domain.TodoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Retrieve & Ensure User
		provider, uid := extractProviderAndUID(c)
		user, err := ensureUser(c.Request.Context(), u, provider, uid)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		var jsonInput PostTodoInput
		err = c.BindJSON(&jsonInput)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		todo, err := t.Create(c, user.ID, jsonInput.Data.Description)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		response := make(map[string]string)
		response["id"] = fmt.Sprint(todo.ID)
		response["user_id"] = fmt.Sprint(todo.UserID)
		response["description"] = todo.Description

		c.JSON(http.StatusOK, response)
	}
}
