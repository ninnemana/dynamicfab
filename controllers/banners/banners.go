package banners

import (
	"appengine"
	"appengine/blobstore"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/banner"
	"net/http"
	"strconv"
)

func Serve(rw http.ResponseWriter, req *http.Request, params martini.Params) {
	blobstore.Send(rw, appengine.BlobKey(params["id"]))
	return
}

func All(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)
	bns, err := banner.All(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(bns) == 0 {
		bns = make([]banner.Banner, 0)
	}

	r.JSON(200, bns)
}

func Get(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	var b banner.Banner
	ctx := appengine.NewContext(req)

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		b.ID = int64(intID)
	}

	err = b.Get(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, b)
}
