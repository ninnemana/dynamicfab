package backend

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Home(rw http.ResponseWriter, req *http.Request, r render.Render, sess sessions.Session) {

	var key string
	k := sess.Get("auth")
	if k != nil {
		key = k.(string)
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Key"] = key
	bag["ActiveNav"] = ""
	r.HTML(200, "admin/index", bag)
	return
}
