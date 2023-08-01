package evm

func (d *exchange) agent(fn string) string {
	return d.NID() + "." + fn
}
