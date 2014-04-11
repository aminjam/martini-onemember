package onemember

import (
	"fmt"
	"net/http"

	"github.com/aminjam/martini-onemember/utils"
	"github.com/aminjam/onemember"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/encoder"
)
func Create(w http.ResponseWriter, r *http.Request, enc encoder.Encoder, as onemember.AccountService) (int, string) {
	out, err := as.Create(
		r.FormValue("username"),
		r.FormValue("password"),
		r.FormValue("email"))
	if err != nil {
		return http.StatusBadRequest, string(encoder.Must(enc.Encode(utils.NewError(utils.ErrCodeNotExist, "create failed"))))
	}
	return http.StatusCreated, string(encoder.Must(enc.Encode(out)))
}

func Read(enc encoder.Encoder, as onemember.AccountService, parms martini.Params) (int, string) {
	out, err := as.GetByUsername(parms["username"])
	if err != nil || out == nil {
		return http.StatusNotFound, string(encoder.Must(enc.Encode(utils.NewError(utils.ErrCodeNotExist, fmt.Sprintf("the entry with id %s does not exist.", parms["username"])))))
	}
	return http.StatusOK, string(encoder.Must(enc.Encode(out)))
}

