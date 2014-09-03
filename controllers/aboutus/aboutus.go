package aboutus

import (
	"appengine"
	"github.com/martini-contrib/render"
	"models/aboutus"
	"net/http"
)

func Get(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)

	au, err := aboutus.Get(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, au)
}
