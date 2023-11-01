package temlates_test

import (
	"bytes"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func TestTemplates(t *testing.T) {
	type Templ struct {
		Name    string
		Content string
	}
	type TestDatum struct {
		Name             string
		Templates        []Templ
		Data             any
		ParseErr         string
		RenderedTemplate string
	}

	for _, td := range []TestDatum{
		{
			Name: "Simple template",
			Templates: []Templ{
				{
					Content: "{{ .Message }}",
				},
			},
			Data:             map[string]any{"Message": "Hello, world!"},
			ParseErr:         "",
			RenderedTemplate: "Hello, world!",
		},
		{
			Name: "Access top level scope from local scope",
			Templates: []Templ{
				{
					Content: `
{{- with $root := . -}}
{{- with .Nested.Data -}}
<p>{{.}}</p><p>{{$root.Hello}}</p>
{{- end -}}
{{- end -}}
`,
				},
			},
			Data: map[string]any{
				"Hello":  "Hello",
				"Nested": map[string]any{"Data": "NestedData"},
			},
			RenderedTemplate: "<p>NestedData</p><p>Hello</p>",
		},
		{
			Name: "Access top level scope from sub-template local scope is not possible",
			// Just pass the data as the template expects it, change the model
			//
			// Otherwise you can create a pipeline that creates maps and send that map, but this
			// adds too much logic in the view
			Templates: []Templ{
				{
					Content: `
{{- with $root := . -}}
{{- with .Nested.Data -}}
{{ template "T1" $root . }}
{{- end -}}
{{- end -}}
{{define "T1"}}<p>{{$root}}</p><p>{{.}}</p>{{end -}}
`,
				},
			},
			Data: map[string]any{
				"Hello":  "Hello",
				"Nested": map[string]any{"Data": "NestedData"},
			},
			ParseErr: `undefined variable "$root"`,
		},
	} {
		t.Run(td.Name, func(t *testing.T) {
			templ := template.New("")
			var parseErrs []error
			for _, te := range td.Templates {
				var err error
				templ, err = templ.New(te.Name).Parse(te.Content)
				if err != nil {
					parseErrs = append(parseErrs, err)
				}
			}
			if td.ParseErr != "" {
				err := multierror.Append(nil, parseErrs...)
				require.ErrorContains(t, err, td.ParseErr)
				return
			} else {
				require.Empty(t, parseErrs)
			}
			var buf bytes.Buffer
			err := templ.Execute(&buf, td.Data)

			require.NoError(t, err)
			require.Equal(t, td.RenderedTemplate, buf.String())
		})
	}
}
