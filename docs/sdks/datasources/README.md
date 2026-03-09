# Datasources

## Overview

Manage datasources.

### Available Operations

* [GetDatasourceInstanceConfiguration](#getdatasourceinstanceconfiguration) - Get datasource instance configuration
* [UpdateDatasourceInstanceConfiguration](#updatedatasourceinstanceconfiguration) - Update datasource instance configuration

## GetDatasourceInstanceConfiguration

Gets the greenlisted configuration values for a datasource instance. Returns only configuration keys that are exposed via the public API greenlist.


### Example Usage

<!-- UsageSnippet language="go" operationID="getDatasourceInstanceConfiguration" method="get" path="/rest/api/v1/configure/datasources/{datasourceId}/instances/{instanceId}" -->
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

    res, err := s.Datasources.GetDatasourceInstanceConfiguration(ctx, "o365sharepoint", "o365sharepoint_abc123")
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `datasourceID`                                           | `string`                                                 | :heavy_check_mark:                                       | The datasource type identifier (e.g. o365sharepoint)     | o365sharepoint                                           |
| `instanceID`                                             | `string`                                                 | :heavy_check_mark:                                       | The datasource instance identifier                       | o365sharepoint_abc123                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetDatasourceInstanceConfigurationResponse](../../models/operations/getdatasourceinstanceconfigurationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateDatasourceInstanceConfiguration

Updates the greenlisted configuration values for a datasource instance. Only configuration keys that are exposed via the public API greenlist may be set. Returns the full greenlisted configuration after the update is applied.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateDatasourceInstanceConfiguration" method="patch" path="/rest/api/v1/configure/datasources/{datasourceId}/instances/{instanceId}" -->
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

    res, err := s.Datasources.UpdateDatasourceInstanceConfiguration(ctx, "o365sharepoint", "o365sharepoint_abc123", components.UpdateDatasourceConfigurationRequest{
        Configuration: components.DatasourceInstanceConfiguration{
            Values: map[string]components.ConfigurationValue{

            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DatasourceConfigurationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        | Example                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |                                                                                                                    |
| `datasourceID`                                                                                                     | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The datasource type identifier (e.g. o365sharepoint)                                                               | o365sharepoint                                                                                                     |
| `instanceID`                                                                                                       | `string`                                                                                                           | :heavy_check_mark:                                                                                                 | The datasource instance identifier                                                                                 | o365sharepoint_abc123                                                                                              |
| `updateDatasourceConfigurationRequest`                                                                             | [components.UpdateDatasourceConfigurationRequest](../../models/components/updatedatasourceconfigurationrequest.md) | :heavy_check_mark:                                                                                                 | N/A                                                                                                                |                                                                                                                    |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |                                                                                                                    |

### Response

**[*operations.UpdateDatasourceInstanceConfigurationResponse](../../models/operations/updatedatasourceinstanceconfigurationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 403, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |