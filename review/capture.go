package review

import (
	"net/http"

	"github.com/itsubaki/appstore-api/appstoreurl"
	"github.com/itsubaki/appstore-api/model"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func Capture(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	if len(r.Header.Get("X-Appengine-Cron")) == 0 {
		log.Warningf(ctx, "X-Appengine-Cron not found.")
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Warningf(ctx, "query[\"id\"] is empty.")
		return
	}

	appname := r.URL.Query().Get("name")
	if appname == "" {
		log.Warningf(ctx, "query[\"name\"] is empty.")
		return
	}

	_, _, country := appstoreurl.Parse(r.URL.Query())
	url := appstoreurl.ReviewURL(id, country)
	log.Infof(ctx, url)

	b, err := appstoreurl.Fetch(ctx, url)
	if err != nil {
		log.Warningf(ctx, err.Error())
		return
	}

	f := model.NewReviewFeed(b)
	log.Infof(ctx, f.Stats())
	for _, r := range f.ReviewList {
		log.Debugf(ctx, r.String())
	}

	Taskq(ctx, id, appname, f)
}
