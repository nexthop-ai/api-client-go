# Authentication

## Overview

Manage indexing API tokens.

### Available Operations

* [Checkdatasourceauth](#checkdatasourceauth) - Check datasource authorization

## Checkdatasourceauth

Returns all datasource instances that require per-user OAuth authorization
for the authenticated user, along with a transient auth token that can be
appended to auth URLs to complete OAuth flows.

Clients construct the full OAuth URL by combining the backend base URL,
the `authUrlRelativePath` from each instance, and the transient auth token:
`<backend>/<authUrlRelativePath>?transient_auth_token=<token>`.


### Example Usage

<!-- UsageSnippet language="go" operationID="checkdatasourceauth" method="post" path="/rest/api/v1/checkdatasourceauth" -->
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

    res, err := s.Authentication.Checkdatasourceauth(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.CheckDatasourceAuthResponse != nil {
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

**[*operations.CheckdatasourceauthResponse](../../models/operations/checkdatasourceauthresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |