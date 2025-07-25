import { Injectable, OnModuleInit } from '@nestjs/common';
import { createClient } from 'redis';
import { OrderGateway } from '../ws/order.gateway';

@Injectable()
export class RedisSubscriber implements OnModuleInit {
  private client = createClient({ url: 'redis://192.168.1.233:6379' });

  constructor(private readonly gateway: OrderGateway) {}

  async onModuleInit() {
    console.log('[RedisSubscriber] Init started');

    try {
      await this.client.connect();
      // console.log('[RedisSubscriber] ✅ Redis connected');

      await this.client.subscribe('order_created', (msg) => {
        // console.log('[RedisSubscriber] 📦 Received order_created:', msg);
        const order = JSON.parse(msg);
        this.gateway.sendOrderCreated(order);
        console.log('[RedisSubscriber] 📥 order_created');
      });

      await this.client.subscribe('order_updated', (msg) => {
        // console.log('[RedisSubscriber] 📦 Received order_updated:', msg);
        const order = JSON.parse(msg);
        this.gateway.sendOrderUpdated(order);
        console.log('[RedisSubscriber] 📥 order_updated');
      });

      await this.client.subscribe('order_assigned', (msg) => {
        // console.log('[RedisSubscriber] 📦 Received order_assigned:', msg);
        const { orderId, staffId } = JSON.parse(msg);
        this.gateway.sendOrderAssigned(orderId, staffId);
        console.log('[RedisSubscriber] 📥 order_assigned');
      });

      await this.client.subscribe('order_cancelled', (msg) => {
        // console.log('[RedisSubscriber] 📦 Received order_cancelled:', msg);
        const { orderId } = JSON.parse(msg);
        this.gateway.sendOrderCancelled(orderId);
        console.log('[RedisSubscriber] 📥 order_cancelled');
      });
    } catch (err) {
      console.error('[RedisSubscriber] ❌ Error in onModuleInit:', err);
    }
  }
}
