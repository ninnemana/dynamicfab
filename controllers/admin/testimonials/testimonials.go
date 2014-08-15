package testimonials

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/testimonials"
	"net/http"
	"strconv"
)

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	c := appengine.NewContext(req)

	ts, err := testimonials.All(c)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Testimonials"] = ts
	bag["ActiveNav"] = "testimonials"
	r.HTML(200, "admin/testimonials/index", bag)

	return
}

func Edit(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)
	var t testimonials.Testimonial

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		t.ID = int64(intID)
	}

	if t.ID > 0 {
		if err := t.Get(c); err != nil {
			http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
			return
		}
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Testimonial"] = t
	bag["ActiveNav"] = "testimonials"
	r.HTML(200, "admin/testimonials/edit", bag)

	return
}

func Save(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	t := testimonials.Testimonial{
		Job: testimonials.Job{
			Title:       req.FormValue("title"),
			Description: req.FormValue("desc"),
		},
		Mention: testimonials.Mention{
			Name: req.FormValue("name"),
			Copy: req.FormValue("copy"),
		},
	}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		t.ID = int64(intID)
	}

	err = t.Save(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/testimonials/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	http.Redirect(rw, req, "/admin/testimonials/"+strconv.Itoa(int(t.ID))+"?success=Testimonial saved", http.StatusFound)
	return
}

func Delete(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	t := testimonials.Testimonial{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		t.ID = int64(intID)
	}

	if err := t.Delete(ctx); err != nil {
		http.Error(rw, "failed to delete testimonial", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}
