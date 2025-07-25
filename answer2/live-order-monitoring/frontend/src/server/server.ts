import { io, Socket } from "socket.io-client"

// สำหรับเชื่อมต่อ WebSocket ด้วย token
export const connectSocket = (token: string): Socket => {
  return io("ws://192.168.1.233:3000", {
    transports: ["websocket"],
    auth: { token },
    forceNew: true,
    reconnection: true,
    timeout: 5000,
  })
}