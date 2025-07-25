
import {
  Controller,
  Get,
  Post,
  Put,
  Param,
  Body,
  Req,
  UseGuards,
} from '@nestjs/common';
import { Roles } from './auth/roles.decorator';
import { RolesGuard } from './auth/roles.guard';
import { Request } from 'express';
import axios from 'axios';

@Controller('orders')
@UseGuards(RolesGuard)
// export class OrderGateway {
export class OrderController {
  // private orderServiceBase = 'http://orders-service:8888/api/orders';
  private orderServiceBase = 'http://192.168.1.233:8888/api/orders';

  @Get()
  @Roles(1, 2) // admin, staff
  async getAllOrders(@Req() req: Request) {
    const res = await axios.get(this.orderServiceBase, {
      headers: { Authorization: req.headers.authorization },
    });
    return res.data;
  }

  // @Post()
  // @Roles(3) // customer
  // async createOrder(@Body() body, @Req() req: Request) {
  //   const res = await axios.post(this.orderServiceBase, body, {
  //     headers: { Authorization: req.headers.authorization },
  //   });
  //   return res.data;
  // }

  @Post()
  async createOrder(@Body() body, @Req() req: Request) {
    const res = await axios.post(this.orderServiceBase, body, {
      headers: { Authorization: req.headers.authorization },
   });
    return res.data;
  }


  @Put(':id')
  @Roles(1) // admin
  async updateOrder(@Param('id') id: string, @Body() body, @Req() req: Request) {
    const res = await axios.put(`${this.orderServiceBase}/${id}`, body, {
      headers: { Authorization: req.headers.authorization },
    });
    return res.data;
  }

  @Put(':id/assign')
  @Roles(1) // admin
  async assignOrder(@Param('id') id: string, @Body() body, @Req() req: Request) {
    const res = await axios.put(`${this.orderServiceBase}/${id}/assign`, body, {
      headers: { Authorization: req.headers.authorization },
    });
    return res.data;
  }

  @Put(':id/cancel')
  @Roles(1) // admin
  async cancelOrder(@Param('id') id: string, @Req() req: Request) {
    const res = await axios.put(`${this.orderServiceBase}/${id}/cancel`, {}, {
      headers: { Authorization: req.headers.authorization },
    });
    return res.data;
  }
}
