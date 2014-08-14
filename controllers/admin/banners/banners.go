package banners

import (
	"appengine"
	"appengine/blobstore"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"models/banner"
	"net/http"
	"strconv"
	"strings"
)

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	c := appengine.NewContext(req)

	bns, err := banner.All(c)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Banners"] = bns
	r.HTML(200, "admin/banners/index", bag)

	return
}

func Edit(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	c := appengine.NewContext(req)
	var b banner.Banner

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		b.ID = int64(intID)
	}

	if b.ID > 0 {
		if err := b.Get(c); err != nil {
			http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
			return
		}
	}

	uploadURL, err := blobstore.UploadURL(c, "/admin/banners/"+strconv.Itoa(intID), nil)
	if err != nil {
		http.Redirect(rw, req, "/admin?error="+err.Error(), http.StatusFound)
		return
	}

	bag := make(map[string]interface{}, 0)
	bag["Host"] = req.URL.Host
	bag["Admin"] = true
	bag["Banner"] = b
	bag["URL"] = uploadURL
	r.HTML(200, "admin/banners/edit", bag)

	return
}

func Save(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	blobs, vals, err := blobstore.ParseUpload(req)
	if err != nil {
		http.Redirect(rw, req, "/admin/banners/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	var title string
	var caption string
	if len(vals["title"]) > 0 {
		title = vals["title"][0]
	}
	if len(vals["caption"]) > 0 {
		caption = vals["caption"][0]
	}

	b := banner.Banner{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		b.ID = int64(intID)
	}

	b.Get(ctx)
	b.Title = title
	b.Caption = caption

	file := blobs["image"]
	if len(file) != 0 {
		b.Image = fmt.Sprintf("/blob/%s", string(file[0].BlobKey))
	} else if intID == 0 {
		http.Redirect(rw, req, "/admin/banners/"+params["id"]+"?error=You must upload an image", http.StatusFound)
		return
	}

	err = b.Save(ctx)
	if err != nil {
		http.Redirect(rw, req, "/admin/banners/"+params["id"]+"?error="+err.Error(), http.StatusFound)
		return
	}

	http.Redirect(rw, req, "/admin/banners/"+strconv.Itoa(int(b.ID))+"?success=Banner saved", http.StatusFound)
	return
}

func Delete(rw http.ResponseWriter, req *http.Request, r render.Render, params martini.Params) {
	ctx := appengine.NewContext(req)

	b := banner.Banner{}

	intID, err := strconv.Atoi(params["id"])
	if err == nil {
		b.ID = int64(intID)
	}

	if err := b.Get(ctx); err != nil {
		http.Error(rw, "failed to delete banner", http.StatusInternalServerError)
		return
	}

	segs := strings.Split(b.Image, "/")
	blobstore.Delete(ctx, appengine.BlobKey(segs[len(segs)-1]))

	if err := b.Delete(ctx); err != nil {
		http.Error(rw, "failed to delete banner", http.StatusInternalServerError)
		return
	}

	r.Status(200)
	return
}
