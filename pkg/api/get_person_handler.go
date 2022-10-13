package api

import (
	"encoding/json"
	"fmt"
)

// personstorage "github.com/darchlabs/api-example/pkg/storage/person"

func getPersonHandler(ctx handlerCtx) handlerRes {
	// set headers
	ctx.w.Header().Set("Content-Type", "application/json")

	id := ctx.r.FormValue("id")

	pp, err := ctx.ps.GetPersonById(id)
	if err != nil {
		return handlerRes{"", 500, err}
	}

	ppString := fmt.Sprint(json.Marshal(pp))

	return handlerRes{ppString, 200, nil}
}
