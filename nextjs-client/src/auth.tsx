import NextAuth from "next-auth"
import Credentials from "next-auth/providers/credentials"
import { cookies } from "next/headers";
 
export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    Credentials({
      credentials: {
        email_id: {},
        password: {},
      },
      authorize: async (credentials) => {
        try {
          const response = await loginUser(credentials)
          if (response.status === 200) {
            const data = await response.json()
            console.log(data)
            const cookie = await cookies()
            cookie.set("Authorization", data.accessToken,{
              httpOnly: true,
              expires: data.expiresAt,
              sameSite: "strict",
              path: "/",
            })
            const userResponse = await fetch('http://localhost:3001/api/v1/users/me')
            const userData = await userResponse.json()
            console.log(userData)
            const user = userData["data"];
            return user
          } else {
            return null
          }
        } catch (error) {
          console.log(error.message)
          return null
        }
      },
    }),
  ],
})

async function loginUser(credentials : Partial<Record<"email_id" | "password", unknown>>) {
  const response = await fetch("http://localhost:3001/auth/login", {
    method: "POST",
    headers: {
      'Content-Type': 'application/json',
      origin: 'nextjs-client',
    },
    body: JSON.stringify(credentials)
  });
  return response;
}