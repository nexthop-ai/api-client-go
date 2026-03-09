# ToolType

Type of tool (READ, WRITE)

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ToolTypeRead

// Open enum: custom values can be created with a direct type cast
custom := components.ToolType("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `ToolTypeRead`  | READ            |
| `ToolTypeWrite` | WRITE           |