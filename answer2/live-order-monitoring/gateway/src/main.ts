

import * as dotenv from 'dotenv';
dotenv.config();

import { NestFactory } from '@nestjs/core';
import { GatewayModule } from './gateway.module';

async function bootstrap() {
  const app = await NestFactory.create(GatewayModule);

  app.enableCors({
    origin: '*', 
    methods: ['GET', 'POST', 'PUT', 'PATCH', 'DELETE', 'OPTIONS'],
    allowedHeaders: ['Content-Type', 'Authorization', 'X-Auth-Token', 'Origin'],
    credentials: true, // ถ้า fetch แบบมี credentials ต้องใส่
  });

  app.use((req, res, next) => {
    console.log('[DEBUG] Incoming Request:', req.method, req.url);
    console.log('[DEBUG] Authorization:', req.headers.authorization);
    next();
  });

  await app.listen(process.env.PORT || 3000);
}
bootstrap();
