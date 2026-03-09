# ScopeType

Describes the scope for a ReadPermission, WritePermission, or GrantPermission object

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ScopeTypeGlobal

// Open enum: custom values can be created with a direct type cast
custom := components.ScopeType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `ScopeTypeGlobal` | GLOBAL            |
| `ScopeTypeOwn`    | OWN               |