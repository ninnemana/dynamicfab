package equipment

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Equipment struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" xml:"-"`
	Description string    `json:"desc" xml:"-"`
	Components  []string  `json:"components" xml:"-"`
	Created     time.Time `json:"-" xml:"-"`
}

func All(ctx appengine.Context) ([]Equipment, error) {
	q := datastore.NewQuery("Equipment").Order("-Created")

	var eq []Equipment
	_, err := q.GetAll(ctx, &eq)
	return eq, err
}

func (e *Equipment) Get(ctx appengine.Context) error {
	return datastore.Get(ctx, e.key(ctx), e)
}

func (e *Equipment) Save(ctx appengine.Context) error {

	var err error
	k, err := datastore.Put(ctx, e.key(ctx), e)
	if k != nil {
		e.ID = k.IntID()
		datastore.Put(ctx, k, e)
	}
	return err
}

func (e *Equipment) Delete(ctx appengine.Context) error {
	return datastore.Delete(ctx, e.key(ctx))
}

func (e *Equipment) key(ctx appengine.Context) *datastore.Key {
	if e.ID == 0 {
		e.Created = time.Now()
		return datastore.NewIncompleteKey(ctx, "Equipment", defaultEquipment(ctx))
	}
	return datastore.NewKey(ctx, "Equipment", "", e.ID, defaultEquipment(ctx))
}

func defaultEquipment(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Equipment", "default", 0, nil)
}
