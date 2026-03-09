# ReportStatusResponseStatus

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ReportStatusResponseStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.ReportStatusResponseStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ReportStatusResponseStatusPending`    | PENDING                                |
| `ReportStatusResponseStatusSuccess`    | SUCCESS                                |
| `ReportStatusResponseStatusFailure`    | FAILURE                                |
| `ReportStatusResponseStatusCancelled`  | CANCELLED                              |
| `ReportStatusResponseStatusCancelling` | CANCELLING                             |
| `ReportStatusResponseStatusActive`     | ACTIVE                                 |