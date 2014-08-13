package backend

import (
	"github.com/martini-contrib/render"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request, r render.Render) {

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	r.HTML(200, "admin/index", bag)
}
