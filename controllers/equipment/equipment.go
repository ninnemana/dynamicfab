package equipment

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/equipment"
	"net/http"
	"strconv"
)

func All(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)
	eq, err := equipment.All(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if eq == nil {
		eq = make([]equipment.Equipment, 0)
	}

	r.JSON(200, eq)
}

func Get(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	var e equipment.Equipment
	ctx := appengine.NewContext(req)

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		e.ID = int64(intID)
	}

	err = e.Get(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	r.JSON(200, e)
}
