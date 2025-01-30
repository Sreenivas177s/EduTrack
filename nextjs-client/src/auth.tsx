import NextAuth, { User } from "next-auth"
import Credentials from "next-auth/providers/credentials"
  
export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    Credentials({
      credentials: {
        email_id: {name: "email_id"},
        password: {name: "password",type: "password"},
      },
      authorize: async (credentials) => {
        try {
          const response = await loginUser(credentials)
          console.log(credentials,response.status)
          if (response.status === 200) {
            const data = await response.json()
            return {id: data.user.ID, name: data.user.first_name, email: data.user.email_id} as User
          } else {
            return null
          }
        } catch (error) {
          console.log("error :",error)
          return null
        }
      },
    }),
  ],
  pages: {
    signIn: "/login",
  },
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