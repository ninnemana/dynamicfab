package dynamic

import (
	"controllers/admin"
	"controllers/auth"
	"flag"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"html/template"
	"net/http"
	"strings"
)

var (
	listenAddr  = flag.String("port", ":8080", "http listen address")
	environment = flag.String("env", "", "current app environment")
)

func init() {
	flag.Parse()
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(martini.Static("public"))
	store := sessions.NewCookieStore([]byte("dynamic-fab"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{
			{
				"StringsEqual": func(a, b string) bool {
					if strings.EqualFold(a, b) {
						return true
					}
					return false
				},
			},
		},
		Delims:          render.Delims{"{{", "}}"},
		Charset:         "UTF-8",
		IndentJSON:      true,
		HTMLContentType: "text/html",
	}))

	// Backend tasks
	m.Group("/admin", func(r martini.Router) {
		r.Get("", auth.Check, backend.Home)
		r.Get("/auth", auth.Index)
		r.Post("/auth", auth.Login)
	})

	// Serve Frontend
	m.Get("/**", func(rw http.ResponseWriter, req *http.Request, r render.Render) {
		bag := make(map[string]interface{}, 0)
		bag["Host"] = req.URL.Host
		r.HTML(200, "index", bag)
	})

	http.Handle("/", m)
}
