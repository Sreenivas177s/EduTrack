import { initializeRequestInfo, RequestInfo, HttpMethod, ApiResponse } from "@/types/apiTypes"
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// client - only method
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

