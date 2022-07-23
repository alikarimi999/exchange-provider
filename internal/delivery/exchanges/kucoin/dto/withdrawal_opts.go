package dto

// func Options(coin entity.Coin) (map[string]string, error) {
// 	opts := make(map[string]string)

// 	switch coin.Symbol {
// 	case "BTC", "ADA", "SOL", "BCH", "LTC", "TRX", "BTT":
// 		opts["chain"] = ""
// 		return opts, nil

// 	case "USDT":
// 		switch coin.Chain {
// 		case entity.ChainTRC20:
// 			opts["chain"] = "TRC20"
// 			return opts, nil

// 		case entity.ChainERC20:
// 			opts["chain"] = "ERC20"
// 			return opts, nil
// 		default:
// 			return nil, errors.New("invalid chain")
// 		}
// 	default:
// 		return nil, errors.New("invalid symbol")
// 	}

// }
