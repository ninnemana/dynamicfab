package content

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/content"
	"net/http"
	"strconv"
)

func All(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)
	con, err := content.All(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, con)
}

func Get(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	var con content.Content
	ctx := appengine.NewContext(req)

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		con.ID = int64(intID)
	}

	err = con.Get(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, con)
}
