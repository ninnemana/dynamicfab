package banner

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type Banner struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Caption   string    `json:"caption"`
	DateAdded time.Time `json:"date_added"`
}

func All(ctx appengine.Context) ([]Banner, error) {

	q := datastore.NewQuery("Banner").Order("-DateAdded")

	// To retrieve the results,
	// you must execute the Query using its GetAll or Run methods.
	var bns []Banner
	_, err := q.GetAll(ctx, &bns)
	if err != nil {
		return bns, err
	}

	return bns, nil
}

func (b *Banner) Get(ctx appengine.Context) error {
	return datastore.Get(ctx, b.key(ctx), b)
}

func (b *Banner) Save(ctx appengine.Context) error {
	var err error
	k, err := datastore.Put(ctx, b.key(ctx), b)
	if k != nil {
		b.ID = k.IntID()
		datastore.Put(ctx, k, b)
	}
	return err
}

func (b *Banner) Delete(ctx appengine.Context) error {
	err := datastore.Delete(ctx, b.key(ctx))
	return err
}

func NewKey(ctx appengine.Context, str_id string, int_id int64) *datastore.Key {
	if str_id != "" {
		return datastore.NewKey(ctx, "Banner", str_id, 0, bannerKey(ctx))
	}
	return datastore.NewKey(ctx, "Banner", "", int_id, bannerKey(ctx))
}

func NewIncompleteKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewIncompleteKey(ctx, "Banner", bannerKey(ctx))
}

func (b *Banner) key(c appengine.Context) *datastore.Key {
	if b.ID == 0 {
		b.DateAdded = time.Now()
		return datastore.NewIncompleteKey(c, "Banner", defaultBanners(c))
	}
	return datastore.NewKey(c, "Banner", "", b.ID, defaultBanners(c))
}

func defaultBanners(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Banners", "default", 0, nil)
}

func bannerKey(ctx appengine.Context) *datastore.Key {
	return datastore.NewKey(ctx, "Banner", "default_banners", 0, nil)
}
