package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// example in https://github.com/gin-gonic/examples/blob/master/assets-in-binary/example02/main.go
//
//go:embed templates/todo.gohtml
var f embed.FS

func main() {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/todo.gohtml"))
	router.SetHTMLTemplate(templ)

	router.GET("/todos/:list-id", func(c *gin.Context) {
		listId := c.Param("list-id")
		type Todo struct {
			Title string
			Done  bool
		}

		type TodoPageData struct {
			PageTitle string
			Todos     []Todo
		}
		data := TodoPageData{
			PageTitle: "My list " + listId,
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		c.HTML(http.StatusOK, "todo.gohtml", data)
	})

	router.GET("/foo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bar.tmpl", gin.H{
			"title": "Foo website",
		})
	})

	router.Run(":8080")
}
