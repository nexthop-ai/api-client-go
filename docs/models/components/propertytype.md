# PropertyType

The type of custom property - this governs the search and faceting behavior. Note that MULTIPICKLIST is not yet supported.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PropertyTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.PropertyType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PropertyTypeText`          | TEXT                        |
| `PropertyTypeDate`          | DATE                        |
| `PropertyTypeInt`           | INT                         |
| `PropertyTypeUserid`        | USERID                      |
| `PropertyTypePicklist`      | PICKLIST                    |
| `PropertyTypeTextlist`      | TEXTLIST                    |
| `PropertyTypeMultipicklist` | MULTIPICKLIST               |