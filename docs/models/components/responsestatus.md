# ResponseStatus

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ResponseStatusAccepted

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ResponseStatusAccepted`   | ACCEPTED                   |
| `ResponseStatusDeclined`   | DECLINED                   |
| `ResponseStatusNoResponse` | NO_RESPONSE                |
| `ResponseStatusTentative`  | TENTATIVE                  |