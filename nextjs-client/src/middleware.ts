import { NextRequest, NextResponse } from 'next/server'
import { cookies } from 'next/headers'
 
export default async function middleware(req: NextRequest) {
  // 2. Check if the current route is protected or public
  const path = req.nextUrl.pathname
  const isProtectedRoute = path.startsWith('/ui/')
  const isPublicRoute = path.startsWith('/login') || path.startsWith('/signup')
  const rootRoute = path === '/'
 
  // 3. Decrypt the session from the cookie
  const cookie = (await cookies())
  const token = cookie.get('authjs.session-token')

  // 4. Redirect to /login if the user is not authenticated
  if ((isProtectedRoute || rootRoute) && !token) {
    return NextResponse.redirect(new URL('/login', req.nextUrl))
  }
 
  // 5. Redirect to /dashboard if the user is authenticated
  if ((isPublicRoute || rootRoute) && token) {
    return NextResponse.redirect(new URL('/ui/home', req.nextUrl))
  }
 
  return NextResponse.next()
}
 
// Routes Middleware should not run on
export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
}