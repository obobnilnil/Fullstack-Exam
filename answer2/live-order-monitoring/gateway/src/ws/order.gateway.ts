import {
  WebSocketGateway,
  WebSocketServer,
  OnGatewayInit,
  OnGatewayConnection,
  OnGatewayDisconnect,
} from '@nestjs/websockets';
import { Server, Socket } from 'socket.io';
import * as jwt from 'jsonwebtoken';

@WebSocketGateway({
  cors: { origin: '*' },
})
export class OrderGateway
  implements OnGatewayInit, OnGatewayConnection, OnGatewayDisconnect
{
  @WebSocketServer()
  server: Server;

  afterInit() {
    console.log('[WS] Gateway initialized');
  }

  handleConnection(client: Socket) {
    const token = client.handshake.auth?.token;

    if (!token) {
      console.log('[WS] ❌ Token not found in handshake');
      client.disconnect();
      return;
    }

    const secret = process.env.JWT_SECRET;
    if (!secret) {
      console.log('[WS] ❌ JWT_SECRET is not defined in environment');
      client.disconnect();
      return;
    }

    try {
      const payload = jwt.verify(token, secret) as any;

      const role = payload.role || payload.role_id;
      if (![1, 2].includes(role)) {
        console.log('[WS] ❌ Unauthorized role:', role);
        client.disconnect();
        return;
      }

      console.log('[WS] ✅ Socket connected:', payload.username || payload.sub);
      client.data.user = payload;
    } catch (err) {
      console.log('[WS] ❌ Token verification failed:', err.message);
      client.disconnect();
    }
  }

  handleDisconnect(client: Socket) {
    console.log('[WS] Client disconnected');
  }

  sendOrderCreated(order: any) {
    this.server.emit('orderCreated', order);
  }

  sendOrderUpdated(order: any) {
    this.server.emit('orderUpdated', order);
  }

  sendOrderAssigned(orderId: string, staffId: string) {
    this.server.emit('orderAssigned', { orderId, staffId });
  }

  sendOrderCancelled(orderId: string) {
    this.server.emit('orderCancelled', { orderId, status: 'cancelled' });
  }
}
