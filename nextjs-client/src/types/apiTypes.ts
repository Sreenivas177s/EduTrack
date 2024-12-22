'use client';
import Bowser from 'bowser';

export type BrowserInfo = {
    name: string;
    version: string;
};

export type HttpMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH';

export type RequestInfo = {
    endpoint: string;
    method: HttpMethod;
    headers?: Record<string, string>;
    queryParams?: Record<string, string | number>;
    body?: Record<string, unknown>;
    browserInfo?: BrowserInfo;
};

export type ApiResponse = {
    statusCode: number;
    message?: string;
    data?: object;
    entity?: string;
};

export function initializeRequestInfo(
    endpoint: string,
    method: HttpMethod,
    queryParams?: Record<string, string | number>,
    body?: Record<string, unknown>
): RequestInfo {
    const browser = Bowser.getParser(window.navigator.userAgent);

    const browserInfo: BrowserInfo = {
        name: browser.getBrowserName() || 'Unknown Browser',
        version: browser.getBrowserVersion() || 'Unknown Version'
    };

    return {
        endpoint,
        method,
        queryParams,
        body,
        browserInfo
    };
}