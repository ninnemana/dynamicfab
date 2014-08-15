package quote

import (
	"appengine"
	"appengine/datastore"
	"errors"
	"time"
)

type Quote struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Description string    `json:"desc"`
	Created     time.Time `json:"-"`
}

func All(ctx appengine.Context) ([]Quote, error) {
	q := datastore.NewQuery("Quote").Order("-Created")

	var qs []Quote
	_, err := q.GetAll(ctx, &qs)
	return qs, err
}

func (q *Quote) Get(ctx appengine.Context) error {
	return datastore.Get(ctx, q.key(ctx), q)
}

func (q *Quote) Save(ctx appengine.Context) error {
	if err := q.validate(); err != nil {
		return err
	}

	var err error
	k, err := datastore.Put(ctx, q.key(ctx), q)
	if k != nil {
		q.ID = k.IntID()
		datastore.Put(ctx, k, q)
	}
	return err
}

func (q *Quote) validate() error {
	if q.Name == "" {
		return errors.New("name is required")
	}
	if q.Email == "" && q.Phone == "" {
		return errors.New("email or phone number is required")
	}
	if q.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

func (q *Quote) Delete(ctx appengine.Context) error {
	return datastore.Delete(ctx, q.key(ctx))
}

func (q *Quote) key(ctx appengine.Context) *datastore.Key {
	if q.ID == 0 {
		q.Created = time.Now()
		return datastore.NewIncompleteKey(ctx, "Quote", defaultQuotes(ctx))
	}
	return datastore.NewKey(ctx, "Quote", "", q.ID, defaultQuotes(ctx))
}

func defaultQuotes(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Quote", "default", 0, nil)
}
