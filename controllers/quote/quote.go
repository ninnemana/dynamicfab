package quote

import (
	"appengine"
	"encoding/json"
	"io/ioutil"
	"models/quote"
	"net/http"
)

func Submit(rw http.ResponseWriter, req *http.Request) {

	var q quote.Quote
	defer req.Body.Close()

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, &q)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := appengine.NewContext(req)
	if err := q.Save(ctx); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(q)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)

	return
}

func Heading(rw http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)

	heading, err := quote.GetHeading(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/plain")
	rw.Write([]byte(heading.Heading))

	return
}
