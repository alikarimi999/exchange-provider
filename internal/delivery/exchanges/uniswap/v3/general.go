package uniswapv3

import (
	"context"
	"exchange-provider/internal/delivery/exchanges/uniswap/v3/contracts"
	"exchange-provider/pkg/errors"
	"exchange-provider/pkg/wallet/eth"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func (u *UniSwapV3) generalSets() error {
	if err := u.setDefaultProvider(); err != nil {
		return err
	}
	if err := u.setChainId(); err != nil {
		return err
	}

	if err := u.setFactory(); err != nil {
		return err
	}

	if err := u.setWallet(); err != nil {
		return err
	}

	return nil
}

func (u *UniSwapV3) setDefaultProvider() error {
	agent := u.agent("setupProviders")

	client, err := ethclient.Dial(u.provider.URL)
	if err != nil {
		return err
	}

	u.provider.Client = client
	if err := u.provider.ping(); err != nil {
		if u.provider.counter > len(u.backupProvidersURL)*3 {
			return errors.Wrap(errors.NewMesssage("unable to connect to providers"))
		}
		u.provider.counter++
		u.l.Error(agent, errors.Wrap(err, u.provider.URL).Error())
		p := u.backupProvidersURL[0]
		u.backupProvidersURL = u.backupProvidersURL[1:]
		u.backupProvidersURL = append(u.backupProvidersURL, u.provider.URL)
		u.provider.URL = p
		time.Sleep(time.Second * 5)
		return u.setDefaultProvider()
	}
	u.provider.counter = 0
	return nil
}

func (u *UniSwapV3) setChainId() error {
	i, err := u.provider.ChainID(context.Background())
	if err != nil {
		return err
	}
	u.chainId = i
	return nil
}

func (u *UniSwapV3) setFactory() error {
	f, err := contracts.NewUniswapv3Factory(factory, u.provider)
	if err != nil {
		return err
	}
	u.factory = f
	return nil
}

func (u *UniSwapV3) setWallet() error {
	if u.cfg.AccountCount == 0 {
		u.cfg.AccountCount = 1
	}
	if u.cfg.Mnemonic == "" {
		u.cfg.Mnemonic, _ = eth.NewMnemonic(128)
	}
	w, err := eth.NewWallet(u.cfg.Mnemonic, u.provider.Client, u.cfg.AccountCount)
	if err != nil {
		return err
	}
	u.wallet = w
	u.cfg.Accounts, _ = w.AllAccounts()
	return nil
}
