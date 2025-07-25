import { Controller, Post, Body, HttpException } from '@nestjs/common';
import axios from 'axios';

@Controller('users')
export class UsersController {
  private usersServiceBase = 'http://192.168.1.233:8889';

  @Post('login')
  async login(@Body() body: any) {
    try {
      const res = await axios.post(`${this.usersServiceBase}/api/users/login`, body);
      return res.data;
    } catch (err) {
      const status = err.response?.status || 500;
      const message = err.response?.data || 'Unknown error';
      throw new HttpException(message, status);
    }
  }
}
