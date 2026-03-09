# LastScanStatus

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.LastScanStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.LastScanStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `LastScanStatusPending`    | PENDING                    |
| `LastScanStatusSuccess`    | SUCCESS                    |
| `LastScanStatusFailure`    | FAILURE                    |
| `LastScanStatusCancelled`  | CANCELLED                  |
| `LastScanStatusCancelling` | CANCELLING                 |
| `LastScanStatusActive`     | ACTIVE                     |