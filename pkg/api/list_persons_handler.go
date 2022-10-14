package api

type ListPersonsHandler struct {
	storage PersonStorage
}

func NewListPersonsHandler(ps PersonStorage) *ListPersonsHandler {
	return &ListPersonsHandler{storage: ps}
}

func (*ListPersonsHandler) Invoke(ctx *handlerCtx) *handlerRes {
	// set headers
	ctx.w.Header().Set("Content-Type", "application/json")

	pp, err := ctx.ps.List()
	if err != nil {
		return &handlerRes{err.Error(), 500, err}
	}

	return &handlerRes{pp, 200, nil}
}
