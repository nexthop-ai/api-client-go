# RelationType

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.RelationTypeEquals

// Open enum: custom values can be created with a direct type cast
custom := components.RelationType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `RelationTypeEquals`    | EQUALS                  |
| `RelationTypeIDEquals`  | ID_EQUALS               |
| `RelationTypeLt`        | LT                      |
| `RelationTypeGt`        | GT                      |
| `RelationTypeNotEquals` | NOT_EQUALS              |