"use client"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { login } from "@/utils/api"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"

export default function LoginPage() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [error, setError] = useState("")
  const router = useRouter()

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()

    try {
      const result = await login(email, password)

      if (result?.token) {
        console.log("✅ Login success, saving token:", result.token)
        localStorage.setItem("token", result.token)
        router.push("/dashboard")
      } else {
        setError("❌ Login failed: Invalid credentials")
      }
    } catch (err) {
      setError("❌ Login error occurred")
      console.error(err)
    }
  }

  return (
    <div className="max-w-md mx-auto mt-20">
      <h1 className="text-2xl font-bold mb-4">🔐 Login</h1>

      <form onSubmit={handleLogin} className="space-y-4">
        <div>
          <Label htmlFor="email">Email</Label>
          <Input
            id="email"
            type="email"
            value={email}
            onChange={e => setEmail(e.target.value)}
            required
          />
        </div>

        <div>
          <Label htmlFor="password">Password</Label>
          <Input
            id="password"
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            required
          />
        </div>

        {error && <p className="text-red-500">{error}</p>}

        <Button type="submit" className="w-full">
          เข้าสู่ระบบ
        </Button>
      </form>
    </div>
  )
}
