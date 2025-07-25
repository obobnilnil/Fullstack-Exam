"use client"

import { useEffect, useState } from "react"
import { jwtDecode } from "jwt-decode"
import { Socket } from "socket.io-client"
import { connectSocket } from "@/utils/socket"

type Order = {
  id: string
  customer: string
  status: string
}

type DecodedToken = {
  email: string
  role_id: number
  iat: number
  exp: number
}

export default function DashboardPage() {
  const [orders, setOrders] = useState<Order[]>([])
  const [user, setUser] = useState<DecodedToken | null>(null)
  const [socket, setSocket] = useState<Socket | null>(null)
  const [connectionStatus, setConnectionStatus] = useState("disconnected")
  const [tokenError, setTokenError] = useState<string | null>(null)

  useEffect(() => {
    console.log("🔄 Dashboard component mounted")

    const token = localStorage.getItem("token")
    console.log("🔍 Token from localStorage:", token ? "Found" : "Not found")

    if (!token) {
      console.log("❌ No token found")
      setTokenError("Token not found")
      return
    }

    try {
      const decoded = jwtDecode<DecodedToken>(token)

      if (decoded.exp * 1000 < Date.now()) {
        console.warn("❌ Token expired")
        setTokenError("Token expired")
        return
      }

      console.log("✅ Token decoded successfully:", decoded)
      setUser(decoded)

      const socketInstance = connectSocket(token)
      setSocket(socketInstance)

      socketInstance.on("connect", () => {
        console.log("✅ Connected to WebSocket, Socket ID:", socketInstance.id)
        setConnectionStatus("connected")
      })

      socketInstance.on("connect_error", (err: Error) => {
        console.error("❌ WebSocket connection error:", err)
        setConnectionStatus("error")
      })

      socketInstance.on("disconnect", (reason: string) => {
        console.warn("⚠️ WebSocket disconnected:", reason)
        setConnectionStatus("disconnected")
      })

      socketInstance.on("orderCreated", (newOrder: Order) => {
        setOrders((prev) => [newOrder, ...prev])
      })

      socketInstance.on("orderUpdated", (updatedOrder: Order) => {
        setOrders((prev) =>
          prev.map((order) =>
            order.id === updatedOrder.id ? updatedOrder : order
          )
        )
      })

      socketInstance.on("orderAssigned", (assignedOrder: Order) => {
        setOrders((prev) =>
          prev.map((order) =>
            order.id === assignedOrder.id ? assignedOrder : order
          )
        )
      })

      socketInstance.on("orderCancelled", (cancelledOrder: Order) => {
        setOrders((prev) =>
          prev.map((order) =>
            order.id === cancelledOrder.id ? cancelledOrder : order
          )
        )
      })

      return () => {
        console.log("🔌 Cleaning up WebSocket connection")
        socketInstance.disconnect()
      }
    } catch (err) {
      console.error("❌ Invalid token format:", err)
      setTokenError("Invalid token")
    }
  }, [])

  const getRoleName = (role_id?: number): string => {
    if (!role_id) return "Unknown"
    return {
      1: "Admin",
      2: "Staff",
    }[role_id] || "Unknown"
  }

  return (
    <main className="p-6 max-w-2xl mx-auto relative min-h-screen">
      <h1 className="text-2xl font-bold mb-6 text-center">📦 Live Order Dashboard</h1>

      {tokenError && (
        <p className="text-red-600 text-center mb-4">
          ⚠️ {tokenError} – กรุณาเข้าสู่ระบบใหม่
        </p>
      )}

      {!tokenError && (
        <>
          {/* ✅ สถานะ WebSocket */}
          <div className="mb-4 p-3 rounded-lg text-center text-sm">
            <span
              className={`inline-block w-3 h-3 rounded-full mr-2 ${
                connectionStatus === "connected"
                  ? "bg-green-500"
                  : connectionStatus === "error"
                  ? "bg-red-500"
                  : "bg-yellow-500"
              }`}
            ></span>
            WebSocket: {connectionStatus}
            {socket?.id && (
              <span className="ml-2 text-gray-600">(ID: {socket.id})</span>
            )}
          </div>

          {/* ✅ แสดงข้อมูลผู้ใช้ */}
          {user && (
            <div className="mb-4 text-center text-gray-700 text-sm">
              <p>
                👤 Logged in as: <span className="font-medium">{user.email}</span>
              </p>
              <p>
                🛡 Role: <span className="font-medium">{getRoleName(user.role_id)}</span>
              </p>
            </div>
          )}

          {/* ✅ รายการออเดอร์ */}
          <div className="flex flex-col gap-4">
            {orders.length === 0 ? (
              <p className="text-center text-gray-400">ยังไม่มีออเดอร์</p>
            ) : (
              orders.map((order) => (
                <div
                  key={order.id}
                  className="bg-white shadow-md rounded-lg p-4 border flex justify-between items-center"
                >
                  <div>
                    <h2 className="text-lg font-semibold">{order.id}</h2>
                    <p className="text-sm text-gray-600">
                      👤 {order.customer}
                    </p>
                  </div>
                  <span
                    className={`text-sm font-medium px-3 py-1 rounded-full ${
                      order.status === "pending"
                        ? "bg-yellow-100 text-yellow-800"
                        : order.status === "assigned"
                        ? "bg-blue-100 text-blue-800"
                        : "bg-green-100 text-green-800"
                    }`}
                  >
                    {order.status}
                  </span>
                </div>
              ))
            )}
          </div>
        </>
      )}
    </main>
  )
}
