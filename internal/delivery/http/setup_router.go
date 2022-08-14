package http

func (r *Router) setup() {
	r.auth = newAuthService(r.v)
	r.NewLimiters()
	r.orderSrvGrpV0()
}
