# Visibilityoverrides
(*Client.Governance.Documents.Visibilityoverrides*)

## Overview

### Available Operations

* [List](#list) - Fetches documents visibility.
* [Create](#create) - Hide/Un-hide docs.

## List

Fetches the visibility override status of the documents passed.

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
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Documents.Visibilityoverrides.List(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDocumentVisibilityOverridesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `docIds`                                                   | []*string*                                                 | :heavy_minus_sign:                                         | List of doc-ids which will have their hide status fetched. |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.GetdocvisibilityResponse](../../models/operations/getdocvisibilityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Sets the visibility-override state of the documents specified, effectively hiding or un-hiding documents.

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
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Governance.Documents.Visibilityoverrides.Create(ctx, components.UpdateDocumentVisibilityOverridesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateDocumentVisibilityOverridesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [components.UpdateDocumentVisibilityOverridesRequest](../../models/components/updatedocumentvisibilityoverridesrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |
| `opts`                                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                                   | :heavy_minus_sign:                                                                                                         | The options for this request.                                                                                              |

### Response

**[*operations.SetdocvisibilityResponse](../../models/operations/setdocvisibilityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |