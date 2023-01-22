package evm

func (d *EvmDex) agent(fn string) string {
	return d.Id() + "." + fn
}
