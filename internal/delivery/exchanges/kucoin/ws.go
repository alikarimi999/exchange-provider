package kucoin

// type webSocket struct {
// 	client *kucoin.WebSocketClient
// 	cache  *wsCache
// 	l      logger.Logger
// 	msg    <-chan *kucoin.WebSocketDownstreamMessage
// 	err    <-chan error
// }

// func (k *kucoinExchange) setupWebSocket(rc *redis.Client) {
// 	k.l.Debug("kucoin: setting up websocket...")
// 	resp, err := k.api.WebSocketPrivateToken()
// 	if err != nil {
// 		k.l.Fatal(fmt.Sprintf("kucoin: failed to get websocket token: %s", err.Error()))
// 	}

// 	tk := &kucoin.WebSocketTokenModel{}
// 	if err := resp.ReadData(tk); err != nil {
// 		k.l.Fatal(fmt.Sprintf("kucoin: failed to read websocket token: %s", err.Error()))
// 	}

// 	c := k.api.NewWebSocketClient(tk)
// 	mc, ec, err := c.Connect()
// 	if err != nil {
// 		k.l.Fatal(fmt.Sprintf("kucoin: failed to connect to websocket: %s", err.Error()))
// 	}

// 	chs := []*kucoin.WebSocketSubscribeMessage{}
// 	for _, topic := range k.cfg.WsTopics {
// 		chs = append(chs, kucoin.NewSubscribeMessage(topic, true))
// 	}

// 	err = c.Subscribe(chs...)
// 	if err != nil {
// 		k.l.Fatal(fmt.Sprintf("kucoin: failed to subscribe: %s", err.Error()))
// 	}

// 	k.ws = &webSocket{
// 		client: c,
// 		cache:  newCache(rc),
// 		msg:    mc,
// 		err:    ec,
// 	}
// 	k.l.Debug("kucoin: websocket setup done")
// 	return
// }

// run is a loop that listen to the websocket messages
// and process the received messages
// if the message is needed by the service cache it. just else ignore it.
// func (w *webSocket) run(wg *sync.WaitGroup) {
// defer wg.Done()
// w.l.Debug("kucoin: running websocket...")
// for {
// 	select {
// 	case m := <-w.msg:

// 		go func(msg *kucoin.WebSocketDownstreamMessage) {
// 			switch msg.Topic{
// 			case "/account/balance":

// 			}
// 		}(m)

// 	}
// }
// }
