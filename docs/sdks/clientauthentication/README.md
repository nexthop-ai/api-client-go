# ClientAuthentication
(*Client.Authentication*)

## Overview

### Available Operations

* [CreateToken](#createtoken) - Create authentication token

## CreateToken

Creates an authentication token for the authenticated user. These are
specifically intended to be used with the [Web SDK](https://developers.glean.com/web).

Note: The tokens generated from this endpoint are **not** valid tokens
for use with the Client API (e.g. `/rest/api/v1/*`).


### Example Usage

<!-- UsageSnippet language="go" operationID="createauthtoken" method="post" path="/rest/api/v1/createauthtoken" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Authentication.CreateToken(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateAuthTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateauthtokenResponse](../../models/operations/createauthtokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |