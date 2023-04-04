package evm

func (d *evmDex) agent(fn string) string {
	return d.NID() + "." + fn
}
