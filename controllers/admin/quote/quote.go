package quote

import (
	"appengine"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/quote"
	"net/http"
	"strconv"
)

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	ctx := appengine.NewContext(req)

	quotes, err := quote.All(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	heading, _ := quote.GetHeading(ctx)

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Quotes"] = quotes
	bag["Heading"] = heading
	bag["ActiveNav"] = "quotes"
	r.HTML(200, "admin/quote/index", bag)

	return
}

func View(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)
	var q quote.Quote

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		q.ID = int64(intID)
	}

	if q.ID > 0 {
		if err := q.Get(c); err != nil {
			http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
			return
		}
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Quote"] = q
	bag["ActiveNav"] = "quotes"
	r.HTML(200, "admin/quote/view", bag)

	return
}

func Delete(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	q := quote.Quote{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		q.ID = int64(intID)
	}

	if err := q.Get(ctx); err != nil {
		http.Error(rw, "failed to delete quote", http.StatusInternalServerError)
		return
	}

	if err := q.Delete(ctx); err != nil {
		http.Error(rw, "failed to delete quote", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}

func SetHeading(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	h := quote.Heading{}
	h.Heading = req.FormValue("heading")

	if err := h.Save(ctx); err != nil {
		http.Error(rw, "failed to save heading", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}
