package dto

// func Coin(sym, chain string) (entity.Coin, error) {
// 	switch sym {
// 	case "BTC":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainBTC,
// 		}, nil

// 	case "ADA":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainADA,
// 		}, nil

// 	case "SOL":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainSOL,
// 		}, nil

// 	case "BCH":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainBCH,
// 		}, nil

// 	case "LTC":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainLTC,
// 		}, nil

// 	case "TRX", "":
// 		return entity.Coin{
// 			Symbol: sym,
// 			Chain:  entity.ChainTRC20,
// 		}, nil

// 	case "USDT":
// 		switch chain {
// 		case "TRC20":
// 			return entity.Coin{
// 				Symbol: sym,
// 				Chain:  entity.ChainTRC20,
// 			}, nil
// 		case "ERC20":
// 			return entity.Coin{
// 				Symbol: sym,
// 				Chain:  entity.ChainERC20,
// 			}, nil
// 		default:
// 			return entity.Coin{}, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("unsupported chain: %s for coin: %s", chain, sym)))
// 		}

// 	case "BTT":
// 		switch chain {
// 		case "TRC20":
// 			return entity.Coin{
// 				Symbol: sym,
// 				Chain:  entity.ChainTRC20,
// 			}, nil
// 		default:
// 			return entity.Coin{}, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("unsupported chain: %s for coin: %s", chain, sym)))
// 		}
// 	default:
// 		return entity.Coin{}, errors.Wrap(errors.ErrBadRequest, errors.NewMesssage(fmt.Sprintf("unsupported coin: %s", sym)))

// 	}
// }
