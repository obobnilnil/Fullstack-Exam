// สำหรับ login ปกติ HTTP POST
export async function login(email: string, password: string) {
  try {
    const res = await fetch("http://192.168.1.233:3000/users/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    })

    if (!res.ok) {
      throw new Error("Login failed")
    }

    const data = await res.json()
    return data
  } catch (err) {
    console.error(err)
    throw err
  }
}
