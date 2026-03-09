# Client.Verification

## Overview

### Available Operations

* [AddReminder](#addreminder) - Create verification
* [List](#list) - List verifications
* [Verify](#verify) - Update verification

## AddReminder

Creates a verification reminder for the document. Users can create verification reminders from different product surfaces.

### Example Usage

<!-- UsageSnippet language="go" operationID="addverificationreminder" method="post" path="/rest/api/v1/addverificationreminder" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Verification.AddReminder(ctx, components.ReminderRequest{
        DocumentID: "<id>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Verification != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `reminderRequest`                                                                                                                                                                                   | [components.ReminderRequest](../../models/components/reminderrequest.md)                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                  | Details about the reminder.                                                                                                                                                                         |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.AddverificationreminderResponse](../../models/operations/addverificationreminderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Returns the information to be rendered in verification dashboard. Includes information for each document owned by user regarding their verifications.

### Example Usage

<!-- UsageSnippet language="go" operationID="listverifications" method="post" path="/rest/api/v1/listverifications" -->
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

    res, err := s.Client.Verification.List(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.VerificationFeed != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `count`                                                                                                                                                                                             | `*int64`                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | Maximum number of documents to return                                                                                                                                                               |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.ListverificationsResponse](../../models/operations/listverificationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Verify

Verify documents to keep the knowledge up to date within customer corpus.

### Example Usage

<!-- UsageSnippet language="go" operationID="verify" method="post" path="/rest/api/v1/verify" -->
```go
package main

import(
	"context"
	"os"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(os.Getenv("GLEAN_API_TOKEN")),
    )

    res, err := s.Client.Verification.Verify(ctx, components.VerifyRequest{
        DocumentID: "<id>",
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Verification != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                           | Type                                                                                                                                                                                                | Required                                                                                                                                                                                            | Description                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                  | The context to use for the request.                                                                                                                                                                 |
| `verifyRequest`                                                                                                                                                                                     | [components.VerifyRequest](../../models/components/verifyrequest.md)                                                                                                                                | :heavy_check_mark:                                                                                                                                                                                  | Details about the verification request.                                                                                                                                                             |
| `locale`                                                                                                                                                                                            | `*string`                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                  | The client's preferred locale in rfc5646 format (e.g. `en`, `ja`, `pt-BR`). If omitted, the `Accept-Language` will be used. If not present or not supported, defaults to the closest match or `en`. |
| `opts`                                                                                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                  | The options for this request.                                                                                                                                                                       |

### Response

**[*operations.VerifyResponse](../../models/operations/verifyresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |