import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { Reflector } from '@nestjs/core';
import { Request } from 'express';

@Injectable()
export class RolesGuard implements CanActivate {
  constructor(private reflector: Reflector) {}

  canActivate(context: ExecutionContext): boolean {
    const allowedRoles = this.reflector.get<number[]>('roles', context.getHandler());
    if (!allowedRoles) return true;

    const req = context.switchToHttp().getRequest<Request>();
    const user = req['user'];

    // console.log('[RolesGuard] Required roles:', allowedRoles);
    // console.log('[RolesGuard] From token role_id:', user?.role_id);

    return user && allowedRoles.includes(user.role_id);
  }
}


