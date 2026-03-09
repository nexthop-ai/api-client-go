# ToolParameterType

Parameter type (string, number, boolean, object, array)

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ToolParameterTypeString

// Open enum: custom values can be created with a direct type cast
custom := components.ToolParameterType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ToolParameterTypeString`  | string                     |
| `ToolParameterTypeNumber`  | number                     |
| `ToolParameterTypeBoolean` | boolean                    |
| `ToolParameterTypeObject`  | object                     |
| `ToolParameterTypeArray`   | array                      |