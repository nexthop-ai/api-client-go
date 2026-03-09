# State

The verification state for the document.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.StateUnverified

// Open enum: custom values can be created with a direct type cast
custom := components.State("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `StateUnverified` | UNVERIFIED        |
| `StateVerified`   | VERIFIED          |
| `StateDeprecated` | DEPRECATED        |