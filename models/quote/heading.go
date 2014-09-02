package quote

import (
	"appengine"
	"appengine/datastore"
	"errors"
)

type Heading struct {
	ID      int64  `json:"id"`
	Heading string `json:"heading"`
}

func GetHeading(ctx appengine.Context) (Heading, error) {
	q := datastore.NewQuery("Heading").Limit(1)

	var h Heading
	var hs []Heading

	_, err := q.GetAll(ctx, &hs)
	if err != nil {
		return h, err
	}

	if len(hs) == 0 {
		return h, nil
	}

	return hs[0], err
}

func (h *Heading) Save(ctx appengine.Context) error {
	q := datastore.NewQuery("Heading").Limit(1)

	var hs []Heading

	keys, err := q.GetAll(ctx, &hs)
	if err != nil {
		return err
	}

	if len(keys) > 0 {
		h.ID = keys[0].IntID()
	}

	k, err := datastore.Put(ctx, h.key(ctx), h)
	if k != nil {
		h.ID = k.IntID()
		datastore.Put(ctx, k, h)
	}
	return err
}

func (h *Heading) validate() error {
	if h.Heading == "" {
		return errors.New("heading cannot be blank")
	}
	return nil
}

func (h *Heading) Delete(ctx appengine.Context) error {
	return datastore.Delete(ctx, h.key(ctx))
}

func (h *Heading) key(ctx appengine.Context) *datastore.Key {
	if h.ID == 0 {
		return datastore.NewIncompleteKey(ctx, "Heading", defaultHeading(ctx))
	}
	return datastore.NewKey(ctx, "Heading", "", h.ID, defaultHeading(ctx))
}

func defaultHeading(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Heading", "default", 0, nil)
}
