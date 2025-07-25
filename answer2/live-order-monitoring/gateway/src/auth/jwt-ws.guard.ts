import {
  CanActivate,
  ExecutionContext,
  Injectable,
  Logger,
  UnauthorizedException,
} from '@nestjs/common';
import { Socket } from 'socket.io';
import * as jwt from 'jsonwebtoken';

@Injectable()
export class JwtWsGuard implements CanActivate {
  private readonly logger = new Logger('JwtWsGuard');

  canActivate(context: ExecutionContext): boolean {
    const client: Socket = context.switchToWs().getClient<Socket>();
    const token = client.handshake.auth?.token;

    if (!token) {
      this.logger.warn('❌ Missing token in WS handshake');
      throw new UnauthorizedException('Missing token');
    }

    try {
      const rawToken = token.replace('Bearer ', '');
      const payload = jwt.verify(rawToken, process.env.JWT_SECRET as string);

      this.logger.debug('✅ Token verified:', payload);
      client.data.user = payload;

      return true;
    } catch (error) {
      this.logger.error('❌ Invalid WS token:', error.message);
      throw new UnauthorizedException('Invalid token');
    }
  }
}
