package api

type GetPersonHandler struct {
	storage PersonStorage
}

func NewGetPersonHandler(ps PersonStorage) *GetPersonHandler {
	return &GetPersonHandler{
		storage: ps,
	}
}

func (*GetPersonHandler) Invoke(ctx *handlerCtx) *handlerRes {
	// set headers
	ctx.w.Header().Set("Content-Type", "application/json")

	id := ctx.r.FormValue("id")

	pp, err := ctx.ps.GetPersonById(id)
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	return &handlerRes{pp, 200, nil}
}
