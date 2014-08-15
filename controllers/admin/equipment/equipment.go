package equipment

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/equipment"
	"net/http"
	"strconv"
)

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	c := appengine.NewContext(req)

	eq, err := equipment.All(c)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Equipment"] = eq
	bag["ActiveNav"] = "equipment"
	r.HTML(200, "admin/equipment/index", bag)

	return
}

func Edit(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)
	var eq equipment.Equipment

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		eq.ID = int64(intID)
	}

	if eq.ID > 0 {
		if err := eq.Get(c); err != nil {
			http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
			return
		}
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Equipment"] = eq
	bag["ActiveNav"] = "equipment"
	r.HTML(200, "admin/equipment/edit", bag)

	return
}

func Save(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	e := equipment.Equipment{}
	req.ParseForm()
	e.Title = req.FormValue("title")
	e.Description = req.FormValue("desc")
	delete(req.Form, "title")
	delete(req.Form, "desc")
	for _, val := range req.Form {
		e.Components = append(e.Components, val...)
	}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		e.ID = int64(intID)
	}

	err = e.Save(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/equipment/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	http.Redirect(rw, req, "/admin/equipment/"+strconv.Itoa(int(e.ID))+"?success=Equipment saved", http.StatusFound)
	return
}

func Delete(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	e := equipment.Equipment{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		e.ID = int64(intID)
	}

	if err := e.Delete(ctx); err != nil {
		http.Error(rw, "failed to delete equipment", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}
