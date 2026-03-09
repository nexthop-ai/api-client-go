# CustomDatasourceConfigConnectorType

The source from which document content was pulled, e.g. an API crawl or browser history

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CustomDatasourceConfigConnectorTypeAPICrawl

// Open enum: custom values can be created with a direct type cast
custom := components.CustomDatasourceConfigConnectorType("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CustomDatasourceConfigConnectorTypeAPICrawl`        | API_CRAWL                                            |
| `CustomDatasourceConfigConnectorTypeBrowserCrawl`    | BROWSER_CRAWL                                        |
| `CustomDatasourceConfigConnectorTypeBrowserHistory`  | BROWSER_HISTORY                                      |
| `CustomDatasourceConfigConnectorTypeBuiltin`         | BUILTIN                                              |
| `CustomDatasourceConfigConnectorTypeFederatedSearch` | FEDERATED_SEARCH                                     |
| `CustomDatasourceConfigConnectorTypePushAPI`         | PUSH_API                                             |
| `CustomDatasourceConfigConnectorTypeWebCrawl`        | WEB_CRAWL                                            |
| `CustomDatasourceConfigConnectorTypeNativeHistory`   | NATIVE_HISTORY                                       |