package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Name: "Learn Go", Done: true},
	{ID: 2, Name: "Build an API with Iris", Done: false},
	{ID: 3, Name: "Learn a new framework", Done: false},
}

func main() {
	app := iris.New()

	app.Handle("GET", "/todos", func(ctx iris.Context) {
		ctx.JSON(todos)
	})

	app.Handle("GET", "/todos/:id", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		for _, todo := range todos {
			if todo.ID == id {
				ctx.JSON(todo)
				return
			}
		}
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Todo not found")
	})

	app.Handle("POST", "/todos", func(ctx iris.Context) {
		var todo Todo
		err := ctx.ReadJSON(&todo)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("Invalid Todo")
			return
		}
		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		ctx.JSON(todo)
	})

	app.Handle("PUT", "/todos/:id", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		var todo Todo
		err := ctx.ReadJSON(&todo)
		if err != nil {
ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("Invalid Todo")
			return
		}
		for i, t := range todos {
			if t.ID == id {
				todo.ID = id
				todos[i] = todo
				ctx.JSON(todo)
				return
			}
		}
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("Todo not found")
	})

	app.Handle("DELETE", "/todos/:id", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		for i, todo := range todos {
			if todo.ID == id {
                todos = append(todos[:i], todos[i+1:]...)
                ctx.StatusCode(iris.StatusOK)
                ctx.WriteString("Todo deleted")
                return
                }
            }
            ctx.StatusCode(iris.StatusNotFound)
            ctx.WriteString("Todo not found")
            })
        
            app.Listen(":"+os.Getenv("PORT"))
        }
        
