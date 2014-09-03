package aboutus

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type AboutUs struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"title"`
	DateAdded time.Time `json:"date_added"`
	Body      string    `json:"body" datastore:",noindex"`
}

func Get(ctx appengine.Context) (AboutUs, error) {
	q := datastore.NewQuery("AboutUs").Limit(1)

	var au AboutUs
	var aus []AboutUs

	_, err := q.GetAll(ctx, &aus)
	if err != nil {
		return au, err
	}

	if len(aus) == 0 {
		return au, nil
	}

	return aus[0], nil
}

func (a *AboutUs) Save(ctx appengine.Context) error {
	var err error
	k, err := datastore.Put(ctx, a.key(ctx), a)
	if k != nil {
		a.ID = k.IntID()
		datastore.Put(ctx, k, a)
	}
	return err
}

func (a *AboutUs) Delete(ctx appengine.Context) error {
	err := datastore.Delete(ctx, a.key(ctx))
	return err
}

func NewKey(ctx appengine.Context, str_id string, int_id int64) *datastore.Key {
	if str_id != "" {
		return datastore.NewKey(ctx, "AboutUs", str_id, 0, aboutUsKey(ctx))
	}
	return datastore.NewKey(ctx, "AboutUs", "", int_id, aboutUsKey(ctx))
}

func NewIncompleteKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewIncompleteKey(ctx, "AboutUs", aboutUsKey(ctx))
}

func (a *AboutUs) key(ctx appengine.Context) *datastore.Key {
	if a.ID == 0 {
		a.DateAdded = time.Now()
		return datastore.NewIncompleteKey(ctx, "AboutUs", defaultAboutUs(ctx))
	}
	return datastore.NewKey(ctx, "AboutUs", "", a.ID, defaultAboutUs(ctx))
}

func defaultAboutUs(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "AboutUs", "default", 0, nil)
}

func aboutUsKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "AboutUs", "default_aboutus", 0, nil)
}
