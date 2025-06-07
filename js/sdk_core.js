import { getCountry, getClientIp, getRequestMethod, getRequestPath } from './helpers.js';

export class RateLimiter {
  constructor(apiKey, serverUrl = 'http://localhost:5000/rate-limit') {
    this.apiKey = apiKey;
    this.serverUrl = serverUrl;
  }

  async checkLimit(request) {
    const path = getRequestPath(request);
    const method = getRequestMethod(request);
    const ip = getClientIp(request) || '0.0.0.0';
    const timestamp = new Date().toISOString().replace('T', ' ').split('.')[0];
    const countryCode = await getCountry(ip);

    try {
      const response = await fetch(this.serverUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          api_key: this.apiKey,
          path,
          timestamp,
          method,
          ip,
          country_code: countryCode
        })
      });

      const data = await response.json();
      return [response.status, data];
    } catch (err) {
      return [500, { error: err.message }];
    }
  }
}
