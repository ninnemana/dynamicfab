package content

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Content struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"title"`
	DateAdded time.Time `json:"date_added"`
	Body      string    `json:"body" datastore:",noindex"`
	Image     string    `json:"image"`
	Link      string    `json:"link"`
}

func All(ctx appengine.Context) ([]Content, error) {

	q := datastore.NewQuery("Content").Order("-DateAdded")

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	var con []Content
	_, err := q.GetAll(ctx, &con)
	if err != nil {
		return con, err
	}

	return con, nil
}

func (c *Content) Get(ctx appengine.Context) error {
	return datastore.Get(ctx, c.key(ctx), c)
}

func (c *Content) Save(ctx appengine.Context) error {
	var err error
	k, err := datastore.Put(ctx, c.key(ctx), c)
	if k != nil {
		c.ID = k.IntID()
		datastore.Put(ctx, k, c)
	}
	return err
}

func (c *Content) Delete(ctx appengine.Context) error {
	err := datastore.Delete(ctx, c.key(ctx))
	return err
}

func NewKey(ctx appengine.Context, str_id string, int_id int64) *datastore.Key {
	if str_id != "" {
		return datastore.NewKey(ctx, "Content", str_id, 0, contentKey(ctx))
	}
	return datastore.NewKey(ctx, "Content", "", int_id, contentKey(ctx))
}

func NewIncompleteKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewIncompleteKey(ctx, "Content", contentKey(ctx))
}

func (c *Content) key(ctx appengine.Context) *datastore.Key {
	if c.ID == 0 {
		c.DateAdded = time.Now()
		return datastore.NewIncompleteKey(ctx, "Content", defaultContent(ctx))
	}
	return datastore.NewKey(ctx, "Content", "", c.ID, defaultContent(ctx))
}

func defaultContent(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Content", "default", 0, nil)
}

func contentKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Content", "default_content", 0, nil)
}
