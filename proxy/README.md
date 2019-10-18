# Proxy

This module will provide a valid response for your AWS Lambda, ready
to pass to the API Gateway.

# Example

> assert.Equal is only used as a helper for the reader

```go
import (
	"github.com/brugnara/lambda-utils/proxy"
)

// simple body:
body := "body"

// json? No problems
js, _ := json.Marshal(yourObject)
body := string(js)

// simple
response := NewResponse(nil, body)
assert.Equal(response.StatusCode, http.StatusOK, ":(") // 500
assert.Equal(response.Body, body, ":(")

// with CORS Header
response := NewResponse(nil, body, true)
assert.Equal(response.StatusCode, http.StatusOK, ":(") // 500
assert.Equal(response.Headers["Access-Control-Allow-Origin"], "*", ":(")
assert.Equal(response.Body, body, ":(")

// handle errors
response := NewResponse(errors.New("A generic error"), body)
assert.Equal(response.StatusCode, http.StatusInternalServerError, ":(") // 500
assert.Equal(response.Body, body, ":(")
```
