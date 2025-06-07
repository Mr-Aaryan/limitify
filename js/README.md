# limitify-sdk

A lightweight SDK for applying API rate limiting Nodejs apps.

## Installation

```bash
npm install limitify-sdk
```

## Usage

```bash
export function ratelimitMiddleware(apiKey) {
  const limiter = new RateLimiter(apiKey);

  return async function (req, res, next) {
    const [status, data] = await limiter.checkLimit(req);

    if (status !== 200) {
      return res.status(status).json({
        error: data?.detail || "Rate limit error",
      });
    }

    next();
  };
}

app.get(
  "/js",
  ratelimitMiddleware("SsgFMxviim38fhEgQQU8JibhBf5AiiskaCWjgf1VqX4"),
  (req, res) => {
    console.log("req jay nepal");
    return 'ho';
  }
);
```
