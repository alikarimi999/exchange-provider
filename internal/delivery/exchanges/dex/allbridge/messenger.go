package allbridge

import "exchange-provider/internal/delivery/exchanges/dex/allbridge/types"

func chooseMessenger(tt types.TransferTime) AllBridgeMessengerProtocol {
	var messenger AllBridgeMessengerProtocol
	if tt.Wormhole == 0 {
		messenger = allbridge
	} else if tt.Allbridge == 0 {
		messenger = wormhole
	} else if tt.Allbridge <= tt.Wormhole {
		messenger = allbridge
	} else {
		messenger = wormhole
	}
	return messenger
}
