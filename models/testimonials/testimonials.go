package testimonials

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Testimonial struct {
	ID      int64     `json:"id"`
	Job     Job       `json:"job"`
	Mention Mention   `json:"mention"`
	Created time.Time `json:"-"`
}

type Job struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type Mention struct {
	Name string `json:"name"`
	Copy string `json:"copy"`
}

func All(ctx appengine.Context) ([]Testimonial, error) {
	q := datastore.NewQuery("Testimonial").Order("-Created")

	var ts []Testimonial
	_, err := q.GetAll(ctx, &ts)
	return ts, err
}

func (t *Testimonial) Get(ctx appengine.Context) error {
	return datastore.Get(ctx, t.key(ctx), t)
}

func (t *Testimonial) Save(ctx appengine.Context) error {

	var err error
	k, err := datastore.Put(ctx, t.key(ctx), t)
	if k != nil {
		t.ID = k.IntID()
		datastore.Put(ctx, k, t)
	}
	return err
}

func (t *Testimonial) Delete(ctx appengine.Context) error {
	return datastore.Delete(ctx, t.key(ctx))
}

func (t *Testimonial) key(ctx appengine.Context) *datastore.Key {
	if t.ID == 0 {
		t.Created = time.Now()
		return datastore.NewIncompleteKey(ctx, "Testimonial", defaultTestimonials(ctx))
	}
	return datastore.NewKey(ctx, "Testimonial", "", t.ID, defaultTestimonials(ctx))
}

func defaultTestimonials(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Testimonial", "default", 0, nil)
}
