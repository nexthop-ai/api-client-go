# UpdateType

Optional type classification for the update.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.UpdateTypeActionable

// Open enum: custom values can be created with a direct type cast
custom := components.UpdateType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `UpdateTypeActionable`  | ACTIONABLE              |
| `UpdateTypeInformative` | INFORMATIVE             |