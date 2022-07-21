package kucoin

// type wsCache struct {
// 	redis *redis.Client
// 	ctx   context.Context
// }

// func newCache(rc *redis.Client) *wsCache {

// 	c := &wsCache{
// 		redis: rc,
// 		ctx:   context.Background(),
// 	}

// 	return c
// }

// func (c *wsCache) recordOrder(order *dto.OrderRecord) error {
// 	key := fmt.Sprintf("kucoin:order:%s", order.OrderId)
// 	return c.redis.Set(c.ctx, key, order, 0).Err()
// }

// func (c *wsCache) getOrder(orderId string) (*dto.OrderRecord, error) {
// 	key := fmt.Sprintf("kucoin:order:%s", orderId)
// 	order := &dto.OrderRecord{}
// 	err := c.redis.Get(c.ctx, key).Scan(order)
// 	return order, err
// }

// func (c *wsCache) deleteOrder(orderId string) error {
// 	key := fmt.Sprintf("kucoin:order:%s", orderId)
// 	return c.redis.Del(c.ctx, key).Err()
// }
