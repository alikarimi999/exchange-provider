package multichain

func (m *Multichain) Name() string {
	return m.cfg.Id
}

func (m *Multichain) AccountId() string {
	return m.cfg.Id
}

func (m *Multichain) NID() string {
	return m.cfg.Id
}
