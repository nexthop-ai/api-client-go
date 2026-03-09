# ExportInfoStatus

The status of the export

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ExportInfoStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ExportInfoStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ExportInfoStatusPending`   | PENDING                     |
| `ExportInfoStatusCompleted` | COMPLETED                   |
| `ExportInfoStatusFailed`    | FAILED                      |