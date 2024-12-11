package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          int
	Name        string
	IsCompleted bool
}

func newTodo(nextID *int, name string, isCompleted bool) Todo {
	*nextID++
	return Todo{
		ID:          *nextID,
		Name:        name,
		IsCompleted: isCompleted,
	}
}

func main() {
	r := gin.Default()

	nextID := 0
	todos := []Todo{
		newTodo(&nextID, "lavar a louça", false),
		newTodo(&nextID, "estudar", false),
		newTodo(&nextID, "jogar futebol", true),
	}

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"Todos": todos,
		})
	})

	r.GET("/todos/:id/edit", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)

		var todo Todo

		for i := range todos {
			if todos[i].ID == idInt {
				todo = todos[i]
				break
			}
		}

		c.HTML(200, "todo-edit", todo)
	})

	r.POST("/todos", func(c *gin.Context) {
		name := c.PostForm("Name")

		newTodo := newTodo(&nextID, name, false)

		todos = append(todos, newTodo)

		c.HTML(200, "form", gin.H{})
		c.HTML(200, "oob-todo", newTodo)
	})

	r.Run(":80")
}
