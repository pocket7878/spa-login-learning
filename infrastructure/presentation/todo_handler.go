package presentation

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pocket7878/spa_login_learning_backend/domain"
)

func TodosGet(u domain.TodoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		todos, err := u.GetTodos(c, 1)
		if err != nil {
			c.AbortWithError(500, err)
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

func TodoPost(u domain.TodoUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jsonInput PostTodoInput
		err := c.BindJSON(&jsonInput)
		if err != nil {
			c.AbortWithError(500, err)
		}
		todo, err := u.Create(c, 1, jsonInput.Data.Description)
		if err != nil {
			c.AbortWithError(500, err)
		}

		response := make(map[string]string)
		response["id"] = fmt.Sprint(todo.ID)
		response["user_id"] = fmt.Sprint(todo.UserID)
		response["description"] = todo.Description

		c.JSON(http.StatusOK, response)
	}
}
