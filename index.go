package dynamic

import (
	"controllers/aboutus"
	"controllers/admin"
	aboutAdmin "controllers/admin/aboutus"
	bannerAdmin "controllers/admin/banners"
	contentAdmin "controllers/admin/content"
	equipmentAdmin "controllers/admin/equipment"
	quoteAdmin "controllers/admin/quote"
	testimonialAdmin "controllers/admin/testimonials"
	"controllers/auth"
	"controllers/banners"
	"controllers/content"
	"controllers/equipment"
	"controllers/quote"
	"controllers/testimonials"
	"flag"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"html/template"
	"net/http"
	"reflect"
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
				"IntegerGreater": func(x interface{}, y interface{}) bool {

					if x == nil || y == nil {
						return false
					}

					var xint int = 0
					var yint int = 0

					xtyp := reflect.TypeOf(x)
					switch xtyp.Kind() {
					case reflect.Int:
						xint = int(x.(int))
					case reflect.Int32:
						xint = int(x.(int32))
					case reflect.Int16:
						xint = int(x.(int16))
					case reflect.Int64:
						xint = int(x.(int64))
					}

					ytyp := reflect.TypeOf(y)
					switch ytyp.Kind() {
					case reflect.Int:
						yint = int(y.(int))
					case reflect.Int32:
						yint = int(y.(int32))
					case reflect.Int16:
						yint = int(y.(int16))
					case reflect.Int64:
						yint = int(y.(int64))
					}

					if xint <= yint {
						return false
					}

					return true
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
	m.Group("/admin/testimonials", func(r martini.Router) {
		r.Get("", auth.Check, testimonialAdmin.Index)
		r.Get("/:id", auth.Check, testimonialAdmin.Edit)
		r.Post("/:id", auth.Check, testimonialAdmin.Save)
		r.Delete("/:id", auth.Check, testimonialAdmin.Delete)
	})
	m.Group("/admin/equipment", func(r martini.Router) {
		r.Get("", auth.Check, equipmentAdmin.Index)
		r.Get("/:id", auth.Check, equipmentAdmin.Edit)
		r.Post("/:id", auth.Check, equipmentAdmin.Save)
		r.Delete("/:id", auth.Check, equipmentAdmin.Delete)
	})
	m.Group("/admin/quotes", func(r martini.Router) {
		r.Get("", auth.Check, quoteAdmin.Index)
		r.Post("/heading", auth.Check, quoteAdmin.SetHeading)
		r.Get("/:id", auth.Check, quoteAdmin.View)
		r.Delete("/:id", auth.Check, quoteAdmin.Delete)
	})

	m.Group("/admin/about", func(r martini.Router) {
		r.Get("", auth.Check, aboutAdmin.Edit)
		r.Post("", auth.Check, aboutAdmin.Save)
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
	m.Group("/api/testimonials", func(r martini.Router) {
		r.Get("", testimonial.All)
		r.Get("/:id", testimonial.Get)
	})
	m.Group("/api/equipment", func(r martini.Router) {
		r.Get("", equipment.All)
		r.Get("/:id", equipment.Get)
	})

	m.Group("/api/quote", func(r martini.Router) {
		r.Post("", quote.Submit)
		r.Get("/heading", quote.Heading)
	})
	m.Get("/api/aboutus", aboutus.Get)

	m.Get("/adduser", auth.AddUser)

	// Serve Frontend
	m.Get("/**", func(rw http.ResponseWriter, req *http.Request, r render.Render) {
		bag := make(map[string]interface{}, 0)
		bag["Host"] = req.URL.Host
		r.HTML(200, "index", bag)
	})

	http.Handle("/", m)
}
