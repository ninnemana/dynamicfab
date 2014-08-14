package banners

import (
	"appengine"
	"appengine/blobstore"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/content"
	"net/http"
	"strconv"
	"strings"
)

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	c := appengine.NewContext(req)

	con, err := content.All(c)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Content"] = con
	r.HTML(200, "admin/content/index", bag)

	return
}

func Edit(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)
	var con content.Content

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		con.ID = int64(intID)
	}

	if con.ID > 0 {
		if err := con.Get(c); err != nil {
			http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
			return
		}
	}

	uploadURL, err := blobstore.UploadURL(c, "/admin/content/"+strconv.Itoa(intID), nil)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Content"] = con
	bag["URL"] = uploadURL
	r.HTML(200, "admin/content/edit", bag)

	return
}

func Save(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	blobs, vals, err := blobstore.ParseUpload(req)
	if err != nil {
		http.Redirect(rw, req, "/admin/content/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	var title string
	var body string
	var link string
	if len(vals["title"]) > 0 {
		title = vals["title"][0]
	}
	if len(vals["body"]) > 0 {
		body = vals["body"][0]
	}
	if len(vals["link"]) > 0 {
		link = vals["link"][0]
	}

	c := content.Content{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		c.ID = int64(intID)
	}

	c.Get(ctx)
	c.Title = title
	c.Body = body
	c.Link = link

	file := blobs["image"]
	if len(file) != 0 {
		c.Image = fmt.Sprintf("/blob/%s", string(file[0].BlobKey))
	}

	err = c.Save(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/content/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	http.Redirect(rw, req, "/admin/content/"+strconv.Itoa(int(c.ID))+"?success=Content saved", http.StatusFound)
	return
}

func Delete(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	c := content.Content{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		c.ID = int64(intID)
	}

	if err := c.Get(ctx); err != nil {
		http.Error(rw, "failed to delete content", http.StatusInternalServerError)
		return
	}

	segs := strings.Split(c.Image, "/")
	blobstore.Delete(ctx, appengine.BlobKey(segs[len(segs)-1]))

	if err := c.Delete(ctx); err != nil {
		http.Error(rw, "failed to delete content", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}
