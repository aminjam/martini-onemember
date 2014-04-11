package main

import (
	"net/http"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/encoder"

	"github.com/aminjam/martini-onemember/routes/onemember"
)

func init() {
	m := martini.Classic()
	m.Use(MapEncoder)

	m.Use(onemember.Handler)
	m.Group("/onemember", onemember.RouteHandler)

	m.Run()
}
func MapEncoder(c martini.Context, w http.ResponseWriter) {
	c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func main() {
}
