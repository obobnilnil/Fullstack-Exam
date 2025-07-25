import { Module, MiddlewareConsumer, NestModule, RequestMethod } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt'; // ✅ ต้อง import เพิ่ม
import { APP_GUARD } from '@nestjs/core';

import { OrderController } from './order.controller';
import { UsersController } from './users.controller';
import { OrderGateway } from './ws/order.gateway';
import { RedisSubscriber } from './redis/redis.subscriber';
import { JwtMiddleware } from './auth/jwt.middleware';
import { RolesGuard } from './auth/roles.guard';


@Module({
  imports: [
    JwtModule.register({
      secret: process.env.JWT_SECRET,
      signOptions: { expiresIn: '1h' },
    }),
  ],
  controllers: [OrderController, UsersController],
  providers: [
    OrderGateway,
    RedisSubscriber,
    {
      provide: APP_GUARD,
      useClass: RolesGuard,
    },
  ],
})
export class GatewayModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer
      .apply(JwtMiddleware)
      .exclude(
        { path: 'users/login', method: RequestMethod.POST },
      )
      .forRoutes('*');
  }
}

