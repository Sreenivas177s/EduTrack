import NextAuth from "next-auth"
import Credentials from "next-auth/providers/credentials"
 
export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    Credentials({
      credentials: {
        email_id: {},
        password: {},
      },
      authorize: async (credentials) => {
        let user = null
        try {
          const response = await fetch("http://localhost:3001/auth/login", {
            method: "POST",
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(credentials)
          });
          console.log(response)
          if (response.status === 200) {
            const data = await response.json();
            user = data;
          }
          if (!user) {
            // No user found, so this is their first attempt to login
            // Optionally, this is also the place you could do a user registration
            return null;
          }
          
          // return user object with their profile data
          return user
        } catch (_) {
          return null
        }
      },
    }),
  ],
})