# ConnectorType

The source from which document content was pulled, e.g. an API crawl or browser history

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ConnectorTypeAPICrawl

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ConnectorTypeAPICrawl`        | API_CRAWL                      |
| `ConnectorTypeBrowserCrawl`    | BROWSER_CRAWL                  |
| `ConnectorTypeBrowserHistory`  | BROWSER_HISTORY                |
| `ConnectorTypeBuiltin`         | BUILTIN                        |
| `ConnectorTypeFederatedSearch` | FEDERATED_SEARCH               |
| `ConnectorTypePushAPI`         | PUSH_API                       |
| `ConnectorTypeWebCrawl`        | WEB_CRAWL                      |
| `ConnectorTypeNativeHistory`   | NATIVE_HISTORY                 |