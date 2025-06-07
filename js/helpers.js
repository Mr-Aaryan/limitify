import maxmind from 'maxmind';
import path from 'path';
import { fileURLToPath } from 'url';

// For __dirname equivalent in ES Modules
const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const dbPath = path.join(__dirname, 'GeoLite2-Country.mmdb');

let countryReader = null;

// Open DB once and reuse
async function loadDb() {
  if (!countryReader) {
    countryReader = await maxmind.open(dbPath);
  }
  return countryReader;
}

export async function getCountry(ip) {
  try {
    const reader = await loadDb();
    const geo = reader.get(ip);
    return geo?.country?.name || 'NA';
  } catch {
    return 'NA';
  }
}

export function getClientIp(req) {
  const forwarded = req.headers['x-forwarded-for'];
  if (forwarded) return forwarded.split(',')[0].trim();

  const realIp = req.headers['x-real-ip'];
  if (realIp) return realIp.trim();

  return req.connection?.remoteAddress || req.socket?.remoteAddress || '127.0.0.1';
}

export function getRequestPath(req) {
  return req.path || new URL(req.url, `http://${req.headers.host}`).pathname || '/';
}

export function getRequestMethod(req) {
  return req.method || 'GET';
}
