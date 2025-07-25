// // src/utils/socket.ts
// import { io, Socket } from "socket.io-client"

// let socket: Socket | null = null

// export const connectSocket = (token: string): Socket => {
//   if (!socket) {
//     socket = io("ws://192.168.1.233:3000", {
//       transports: ["websocket"],
//       auth: {
//         token,
//       },
//       autoConnect: false,
//     })
//   }

//   socket.auth = { token } // 👈 safe assign
//   socket.connect()
//   return socket
// }


import { io, Socket } from "socket.io-client"

let socket: Socket | null = null

export const connectSocket = (token: string): Socket => {
  // disconnect ตัวเก่าก่อน (เพื่อไม่ให้ค้าง)
  if (socket) {
    socket.disconnect()
    socket = null
  }

  console.log("🔑 Token before socket connect:", token)

  // สร้าง socket ใหม่พร้อมส่ง token (ไม่ต้อง prefix "Bearer")
  socket = io("ws://192.168.1.233:3000", {
    transports: ["websocket"],
    auth: {
      token: token, // ✅ ส่ง token เปล่า ไม่ต้องใส่ Bearer
    },
    autoConnect: true, // ✅ ปล่อยให้มัน connect ทันที
  })
    // socket = io("ws://192.168.1.233:3000", {
    // transports: ["websocket"],
    // auth: {
    //     token: `Bearer ${token}` // ✅ แก้ตรงนี้
    // },
    // autoConnect: true,
    // })

  return socket
}
