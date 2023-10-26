package server_test

import (
	"embed"
	"fmt"
	"github.com/stretchr/testify/require"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

//go:embed templates
var templatesFS embed.FS

func TestServer(t *testing.T) {
	engine, err := template.ParseFS(templatesFS, "templates/todo.gohtml")
	require.NoError(t, err)

	templateInstances := engine.Templates()
	var templateNames []string
	for _, v := range templateInstances {
		if v == nil {
			continue
		}
		templateNames = append(templateNames, v.Name())
	}
	fmt.Println("templateNames", templateNames)

	handler := func(w http.ResponseWriter, r *http.Request) {
		listId := r.URL.Query().Get("listId")
		data := TodoPageData{
			PageTitle: "My list " + listId,
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		require.NoError(t, engine.ExecuteTemplate(w, "todo.gohtml", data))
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)

	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	require.Equal(t,
		`<h1>My list </h1>
<ul>
            <li>Task 1</li>
            <li class="done">Task 2</li>
            <li class="done">Task 3</li>
</ul>`, string(body))

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

}
