package dynamic

import (
	"controllers/admin"
	bannerAdmin "controllers/admin/banners"
	contentAdmin "controllers/admin/content"
	quoteAdmin "controllers/admin/quote"
	"controllers/auth"
	"controllers/banners"
	"controllers/content"
	"controllers/quote"
	"flag"
	"fmt"
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
				"ShortenString": func(s string, l int) string {
					if len(s) <= l {
						return s
					}
					return fmt.Sprintf("%s...", s[:l])
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
		r.Get("", backend.Home)
		r.Get("/auth", auth.Index)
		r.Post("/auth", auth.Login)
		r.Get("/auth/out", auth.Logout)
	})

	m.Group("/admin/banners", func(r martini.Router) {
		r.Get("", auth.Check, bannerAdmin.Index)
		r.Get("/:id", auth.Check, bannerAdmin.Edit)
		r.Post("/:id", auth.Check, bannerAdmin.Save)
		r.Delete("/:id", auth.Check, bannerAdmin.Delete)
	})
	m.Group("/admin/content", func(r martini.Router) {
		r.Get("", auth.Check, contentAdmin.Index)
		r.Get("/:id", auth.Check, contentAdmin.Edit)
		r.Post("/:id", auth.Check, contentAdmin.Save)
		r.Delete("/:id", auth.Check, contentAdmin.Delete)
	})
	m.Group("/admin/quotes", func(r martini.Router) {
		r.Get("", auth.Check, quoteAdmin.Index)
		r.Post("/heading", auth.Check, quoteAdmin.SetHeading)
		r.Get("/:id", auth.Check, quoteAdmin.View)
		r.Delete("/:id", auth.Check, quoteAdmin.Delete)
	})

	m.Get("/blob/:id", banners.Serve)
	m.Group("/api/banners", func(r martini.Router) {
		r.Get("", banners.All)
		r.Get("/:id", banners.Get)
	})
	m.Group("/api/content", func(r martini.Router) {
		r.Get("", content.All)
		r.Get("/:id", content.Get)
	})

	m.Group("/api/quote", func(r martini.Router) {
		r.Post("", quote.Submit)
		r.Get("/heading", quote.Heading)
	})

	m.Get("/adduser", auth.AddUser)

	// Serve Frontend
	m.Get("/**", func(rw http.ResponseWriter, req *http.Request, r render.Render) {
		bag := make(map[string]interface{}, 0)
		bag["Host"] = req.URL.Host
		r.HTML(200, "index", bag)
	})

	http.Handle("/", m)
}
