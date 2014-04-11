package onemember

import (
	"os"

	"github.com/aminjam/onemember"
	oa2 "github.com/aminjam/onememberoauth2"
	"github.com/codegangsta/martini"

	"github.com/aminjam/martini-onemember/routes/onemember/data"
)

var as onemember.AccountService
var oa2Consumer oa2.Consumer
var adaptor string

func RouteHandler(r martini.Router) {
	r.Get("/:username", Read)
	r.Post("", Create)
	r.Get("/request/:provider", oa2Consumer.Request)
	r.Get("/callback/:provider", oa2Consumer.Callback)
}

func Handler(c martini.Context) {
	switch adaptor {
	case "couchbase":
		data.SetCouchbaseDB()
		as = onemember.New(data.CouchbaseDB)
	case "mongo":
		as = onemember.New(data.MongoDB)
	}
	c.MapTo(as, (*onemember.AccountService)(nil))
}

func init() {
	options := make(map[string]*oa2.Client)
	options["google"] = oa2.Google(oa2.Client{
		ClientId:     "myClient",
		ClientSecret: "mySecret",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL:  "http://localhost:3000/onemember/callback/google",
		ClaimsBuilder: func(accessToken string) (claims map[string]string, err error) {
			claims = make(map[string]string)
			claims["accessToken"] = accessToken
			return
		},
	})
	oa2Consumer = oa2.New(options)
	adaptor = os.Getenv("ONEMEMBER_ADAPTOR")
}
