package aboutus

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/aboutus"
	"net/http"
)

func Edit(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)

	au, err := aboutus.Get(c)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["AboutUs"] = au
	bag["ActiveNav"] = "about"
	r.HTML(200, "admin/aboutus/edit", bag)

	return
}

func Save(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	au, err := aboutus.Get(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/about?error="+err.Error(), http.StatusFound)
		return
	}
	au.Title = req.FormValue("title")
	au.Body = req.FormValue("body")

	err = au.Save(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/about?error="+err.Error(), http.StatusFound)
		return
	}

	http.Redirect(rw, req, "/admin/about?success=About Us saved", http.StatusFound)
	return
}
