# DlpReportStatus

The status of the policy/report. Only ACTIVE status will be picked for scans.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpReportStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.DlpReportStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `DlpReportStatusActive`    | ACTIVE                     |
| `DlpReportStatusInactive`  | INACTIVE                   |
| `DlpReportStatusCancelled` | CANCELLED                  |
| `DlpReportStatusNone`      | NONE                       |