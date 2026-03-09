# ActivityEnum

Activity e.g. search, home page visit or all.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ActivityEnumAll

// Open enum: custom values can be created with a direct type cast
custom := components.ActivityEnum("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ActivityEnumAll`    | ALL                  |
| `ActivityEnumSearch` | SEARCH               |