package testimonial

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/testimonials"
	"net/http"
	"strconv"
)

func All(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)
	ts, err := testimonials.All(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if ts == nil {
		ts = make([]testimonials.Testimonial, 0)
	}

	r.JSON(200, ts)
}

func Get(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	var t testimonials.Testimonial
	ctx := appengine.NewContext(req)

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		t.ID = int64(intID)
	}

	err = t.Get(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, t)
}
