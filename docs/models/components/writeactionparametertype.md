# WriteActionParameterType

The type of the value (e.g., integer, string, boolean, etc.)

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.WriteActionParameterTypeUnknown

// Open enum: custom values can be created with a direct type cast
custom := components.WriteActionParameterType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `WriteActionParameterTypeUnknown` | UNKNOWN                           |
| `WriteActionParameterTypeInteger` | INTEGER                           |
| `WriteActionParameterTypeString`  | STRING                            |
| `WriteActionParameterTypeBoolean` | BOOLEAN                           |