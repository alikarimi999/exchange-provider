package http

func (r *Router) setup() {
	r.NewLimiters()
	r.orderSrvGrpV0()
}
