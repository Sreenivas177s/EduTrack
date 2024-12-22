import { initializeRequestInfo, RequestInfo, HttpMethod, ApiResponse } from "@/types/apiTypes"
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import { z } from "zod";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export async function invokeApi(endpoint: string, method: HttpMethod, queryParams?: Record<string, string | number>, body?: Record<string, unknown>): Promise<ApiResponse> {
  const rInfo: RequestInfo = initializeRequestInfo(endpoint, method, queryParams, body);
  
  const response = await fetch(rInfo.endpoint, {
    method: rInfo.method,
    headers: {
      'Content-Type': 'application/json',
      ...rInfo.headers
    },
    body: JSON.stringify(rInfo.body)
  });
  
  const data = await response.json();
  
  return {
    statusCode: response.status,
    message: response.statusText,
    data,
  };
}

export const SignInSchema = z.object({
  email_id: z.string({ required_error: "Email is required" })
    .min(1, "Email is required")
    .email("Invalid email"),
  password: z.string({ required_error: "Password is required" })
    .min(1, "Password is required")
    // .min(8, "Password must be more than 8 characters")
    .max(32, "Password must be less than 32 characters"),
})