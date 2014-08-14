package dynamic

import (
	"controllers/admin"
	bannerAdmin "controllers/admin/banners"
	contentAdmin "controllers/admin/content"
	"controllers/auth"
	"controllers/banners"
	"controllers/content"
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
		r.Get("", backend.Home)
		r.Get("/auth", auth.Index)
		r.Post("/auth", auth.Login)
		r.Get("/auth/out", auth.Logout)
	})

	m.Group("/admin/banners", func(r martini.Router) {
		r.Get("", bannerAdmin.Index)
		r.Get("/:id", bannerAdmin.Edit)
		r.Post("/:id", bannerAdmin.Save)
		r.Delete("/:id", bannerAdmin.Delete)
	})
	m.Group("/admin/content", func(r martini.Router) {
		r.Get("", contentAdmin.Index)
		r.Get("/:id", contentAdmin.Edit)
		r.Post("/:id", contentAdmin.Save)
		r.Delete("/:id", contentAdmin.Delete)
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

	// Serve Frontend
	m.Get("/**", func(rw http.ResponseWriter, req *http.Request, r render.Render) {
		bag := make(map[string]interface{}, 0)
		bag["Host"] = req.URL.Host
		r.HTML(200, "index", bag)
	})

	http.Handle("/", m)
}
