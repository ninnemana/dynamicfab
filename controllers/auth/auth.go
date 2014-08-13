package auth

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Username string
	Password string
}

func Check(w http.ResponseWriter, r *http.Request, sess sessions.Session) {
	auth := sess.Get("auth")
	if auth == nil {
		http.Redirect(w, r, "/admin/auth", http.StatusFound)
		return
	}
	log.Println(auth)
}

func Index(rw http.ResponseWriter, req *http.Request, r render.Render) {
	bag := make(map[string]interface{}, 0)
	bag["Admin"] = true
	bag["Auth"] = true
	r.HTML(200, "admin/auth/index", bag)
}

func Login(rw http.ResponseWriter, req *http.Request, sess sessions.Session) {
	username := req.FormValue("username")
	password := req.FormValue("pass")

	c := appengine.NewContext(req)
	q := datastore.NewQuery("User")

	var key *datastore.Key
	for t := q.Run(c); ; {
		var u User
		k, err := t.Next(&u)
		c.Debugf("username: %s", u.Username)
		c.Debugf("password: %s", u.Password)
		if err != nil || k == nil {
			http.Redirect(rw, req, "/admin/auth?error=Failed to login", http.StatusFound)
			return
		}

		if !strings.EqualFold(username, u.Username) {
			http.Redirect(rw, req, "/admin/auth?error=Failed to login", http.StatusFound)
			return
		}

		if !strings.EqualFold(password, u.Password) {
			http.Redirect(rw, req, "/admin/auth?error=Failed to login", http.StatusFound)
			return
		}
		key = k
		break
	}

	if key == nil {
		http.Redirect(rw, req, "/admin/auth?error=Failed to login", http.StatusFound)
		return
	}

	sess.Set("auth", key.Encode())

	http.Redirect(rw, req, "/admin", http.StatusFound)
	return
}
