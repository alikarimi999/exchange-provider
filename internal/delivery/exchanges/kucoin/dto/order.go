package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"order_service/internal/entity"
	"strings"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
)

type CreateOrderRequestModel kucoin.CreateOrderModel

func CreateOrderRequest(from, to entity.Coin, vol string) (*CreateOrderRequestModel, error) {
	o := &CreateOrderRequestModel{
		ClientOid: uuid.New().String(),
	}
	switch from.Symbol {
	case "BTC":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "BTC-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s", from.Symbol, to.Symbol))
		}

	case "ADA":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "ADA-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s ", from.Symbol, to.Symbol))
		}

	case "SOL":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "SOL-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s ", from.Symbol, to.Symbol))
		}

	case "BCH":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "BCH-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s ", from.Symbol, to.Symbol))
		}

	case "LTC":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "LTC-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s ", from.Symbol, to.Symbol))
		}

	case "TRX":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "TRX-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s", from.Symbol, to.Symbol))
		}

	case "BTT":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "BTT-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s", from.Symbol, to.Symbol))
		}

	case "USDT":
		switch to.Symbol {
		case "BTC":
			o.Symbol = "BTC-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "ADA":
			o.Symbol = "ADA-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "SOL":
			o.Symbol = "SOL-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "BCH":
			o.Symbol = "BCH-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "LTC":
			o.Symbol = "LTC-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "TRX":
			o.Symbol = "TRX-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		case "BTT":
			o.Symbol = "BTT-USDT"
			o.Type = "market"
			o.Side = "buy"
			o.Funds = vol

		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s", from.Symbol, to.Symbol))

		}

	default:
		return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s", from.Symbol, to.Symbol))
	}
	return o.fixVolPrecision(), nil
}

// kucoin is sensitive to the precision of the volume, so we need to fix it
// TODO: this is a hack, we should use a proper way to do this
func (o *CreateOrderRequestModel) fixVolPrecision() *CreateOrderRequestModel {

	switch o.Symbol {

	case "BTC-USDT": // limit  8 precision
		if o.Side == "sell" {
			o.Size = trim(o.Size, 8)
		} else {
			o.Funds = trim(o.Funds, 6)
		}

	case "LTC-USDT": // limit 6 precision
		if o.Side == "sell" {
			o.Size = trim(o.Size, 6)
		} else {
			o.Funds = trim(o.Funds, 6)
		}

	case "ADA-USDT", "SOL-USDT", "BCH-USDT", "BTT-USDT", "TRX-USDT": // limit 4 precision
		if o.Side == "sell" {
			o.Size = trim(o.Size, 4)
		} else {
			o.Funds = trim(o.Funds, 6)
		}
	}

	return o
}

func trim(s string, l int) string {
	if s == "" || l == 0 {
		return s
	}

	ss := strings.Split(s, ".")
	var result string

	if len(ss) == 2 {
		if len(ss[1]) > l {
			result = ss[0] + "." + ss[1][:l]
		} else {
			result = s
		}

	} else {
		result = ss[0] + ".0"
	}
	return result
}

type OrderRecord struct {
	OrderId string
	Symbol  string
	Volume  string
}

// implement encoding.BinaryMarshaler
func (o *OrderRecord) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
