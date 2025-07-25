package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	redisinfra "live-order-monitoring/services/orders/infra/redis"
	"live-order-monitoring/services/orders/model"

	"github.com/gin-gonic/gin"
)

var rdb = redisinfra.GetRedisClient()

func publish(channel string, payload any) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[REDIS PUB] ❌ Failed to marshal payload for %s: %v", channel, err)
		return
	}

	err = rdb.Publish(ctx, channel, data).Err()
	if err != nil {
		log.Printf("[REDIS PUB] ❌ Failed to publish to %s: %v", channel, err)
	} else {
		log.Printf("[REDIS PUB] ✅ Published to %s", channel)
	}
}

func PublishOrderCreated(order model.Order) {
	publish("order_created", order)
}

func PublishOrderUpdated(order model.Order) {
	publish("order_updated", order)
}

func PublishOrderAssigned(orderID, staffID string) {
	publish("order_assigned", gin.H{
		"orderId": orderID,
		"staffId": staffID,
	})
}

func PublishOrderCancelled(orderID string) {
	publish("order_cancelled", gin.H{
		"orderId": orderID,
	})
}
