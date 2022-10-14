package api

type DelPersonHandler struct {
	storage PersonStorage
}

func NewDelPersonHandler(ps PersonStorage) *DelPersonHandler {
	return &DelPersonHandler{storage: ps}
}

func (*DelPersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	id := ctx.r.FormValue("id")

	pp, err := ctx.ps.DeletePersonById(id)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	return &handlerRes{pp, 200, nil}
}
