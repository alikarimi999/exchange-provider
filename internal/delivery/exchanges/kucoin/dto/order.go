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
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s by exchange: kucoin", from.Symbol, to.Symbol))
		}

	case "TRX":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "TRX-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s by exchange: kucoin", from.Symbol, to.Symbol))
		}

	case "BTT":
		switch to.Symbol {
		case "USDT":
			o.Symbol = "BTT-USDT"
			o.Type = "market"
			o.Side = "sell"
			o.Size = vol
		default:
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s by exchange: kucoin", from.Symbol, to.Symbol))
		}

	case "USDT":
		switch to.Symbol {
		case "BTC":
			o.Symbol = "BTC-USDT"
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
			return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s by exchange: kucoin", from.Symbol, to.Symbol))

		}

	default:
		return nil, errors.New(fmt.Sprintf("unsupported coin pair: %s-%s by exchange: kucoin", from.Symbol, to.Symbol))
	}
	return o.fixVolPrecision(), nil
}

// kucoin is sensitive to the precision of the volume, so we need to fix it
// TODO: this is a hack, we should use a proper way to do this
func (o *CreateOrderRequestModel) fixVolPrecision() *CreateOrderRequestModel {

	switch o.Symbol {
	case "BTT-USDT", "TRX-USDT":
		if o.Side == "sell" {
			ss := strings.Split(o.Size, ".")
			if len(ss) == 2 {
				o.Size = ss[0] + "." + ss[1][:4]
			} else {
				o.Size = ss[0] + ".0"
			}
		}
	}

	return o
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
