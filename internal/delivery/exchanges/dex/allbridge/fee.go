package allbridge

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
)

type AllBridgeMessengerProtocol uint

const (
	none AllBridgeMessengerProtocol = iota
	allbridge
	wormhole
	layerZero
)

type bridgeFeeReq struct {
	DestinationChainId int `json:"destinationChainId"`
	SourceChainId      int `json:"sourceChainId"`
	Messenger          int `json:"messenger"`
}

type bridgeFeeRes struct {
	Fee string `json:"fee"`
}

func getBridgeFee(srcChain, dstChain int, messenger AllBridgeMessengerProtocol) (*big.Int, error) {
	url := "https://core.api.allbridgecoreapi.net/receive-fee"

	br := &bridgeFeeReq{DestinationChainId: dstChain, SourceChainId: srcChain, Messenger: int(messenger)}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(br)
	req, _ := http.NewRequest(http.MethodPost, url, &buf)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	bRes := &bridgeFeeRes{}
	if err := json.Unmarshal(b, bRes); err != nil {
		return nil, err
	}
	fee, _ := new(big.Int).SetString(bRes.Fee, 10)
	return fee, nil
}
