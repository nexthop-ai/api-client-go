# Policies
(*Client.Governance.Data.Policies*)

## Overview

### Available Operations

* [Retrieve](#retrieve) - Gets specified Policy.
* [Update](#update) - Updates an existing policy.
* [List](#list) - Lists policies.
* [Create](#create) - Creates new policy.
* [Download](#download) - Downloads violations CSV for policy.

## Retrieve

Fetches the specified policy version, or the latest if no version is provided.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Data.Policies.Retrieve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDlpReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                  | Type                                                                                                                                                       | Required                                                                                                                                                   | Description                                                                                                                                                |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                                      | :heavy_check_mark:                                                                                                                                         | The context to use for the request.                                                                                                                        |
| `id`                                                                                                                                                       | *string*                                                                                                                                                   | :heavy_check_mark:                                                                                                                                         | The id of the policy to fetch.                                                                                                                             |
| `version`                                                                                                                                                  | **int64*                                                                                                                                                   | :heavy_minus_sign:                                                                                                                                         | The version of the policy to fetch. Each time a policy is updated, the older version is still stored. If this is left empty, the latest policy is fetched. |
| `opts`                                                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                                                   | :heavy_minus_sign:                                                                                                                                         | The options for this request.                                                                                                                              |

### Response

**[*operations.GetpolicyResponse](../../models/operations/getpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Updates an existing policy.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Data.Policies.Update(ctx, "<id>", components.UpdateDlpReportRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateDlpReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `id`                                                                                   | *string*                                                                               | :heavy_check_mark:                                                                     | The id of the policy to fetch.                                                         |
| `updateDlpReportRequest`                                                               | [components.UpdateDlpReportRequest](../../models/components/updatedlpreportrequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdatepolicyResponse](../../models/operations/updatepolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Lists policies with filtering.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Data.Policies.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.ListDlpReportsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `autoHide`                                                | **bool*                                                   | :heavy_minus_sign:                                        | Filter to return reports with a given value of auto-hide. |
| `frequency`                                               | **string*                                                 | :heavy_minus_sign:                                        | Filter to return reports with a given frequency.          |
| `opts`                                                    | [][operations.Option](../../models/operations/option.md)  | :heavy_minus_sign:                                        | The options for this request.                             |

### Response

**[*operations.GetpoliciesResponse](../../models/operations/getpoliciesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Creates a new policy with specified specifications and returns its id.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Data.Policies.Create(ctx, components.CreateDlpReportRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDlpReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.CreateDlpReportRequest](../../models/components/createdlpreportrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CreatepolicyResponse](../../models/operations/createpolicyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Download

Downloads CSV violations report for a specific policy id. This does not support continuous policies.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Data.Policies.Download(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | The id of the policy to download violations for.         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DownloadpolicycsvResponse](../../models/operations/downloadpolicycsvresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |