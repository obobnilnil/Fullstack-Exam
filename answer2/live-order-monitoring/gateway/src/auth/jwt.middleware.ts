import { Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import * as jwt from 'jsonwebtoken';

@Injectable()
export class JwtMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    // console.log(`[JWT Middleware] 🔸 ${req.method} ${req.path}`);

    // ✅ Bypass for POST /orders
    // if (req.method === 'POST' && req.path === '/orders') {
    if (req.method === 'POST' && req.originalUrl === '/orders') {
      // console.log('[JWT Middleware] ✅ Bypassing auth for POST /orders');
      return next();
    }

    const authHeader = req.headers.authorization;
    if (!authHeader || !authHeader.startsWith('Bearer ')) {
      // console.log('[JWT Middleware] ❌ No or malformed Authorization header');
      return res.status(401).json({ message: 'Unauthorized - token missing' });
    }

    const token = authHeader.split(' ')[1];

    try {
      const secret = process.env.JWT_SECRET;
      if (!secret) {
        console.error('[JWT Middleware] ❌ JWT_SECRET not found in environment variables');
        return res.status(500).json({ message: 'JWT secret not configured on server' });
      }

      const decoded = jwt.verify(token, secret);
      // console.log('✅ JWT decoded:', decoded);

      req['user'] = decoded;
      next();
    } catch (err) {
      // console.error('[JWT Middleware] ❌ Token verification failed:', err.message);
      return res.status(401).json({ message: 'Invalid token' });
    }
  }
}