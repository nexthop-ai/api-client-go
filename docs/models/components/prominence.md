# Prominence

The level of visual distinction that should be given to a result.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ProminenceHero

// Open enum: custom values can be created with a direct type cast
custom := components.Prominence("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ProminenceHero`     | HERO                 |
| `ProminencePromoted` | PROMOTED             |
| `ProminenceStandard` | STANDARD             |