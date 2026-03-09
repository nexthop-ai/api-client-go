# BulkUploadHistoryEventStatus

The status of the upload, an enum of ACTIVE, SUCCESSFUL

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.BulkUploadHistoryEventStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.BulkUploadHistoryEventStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `BulkUploadHistoryEventStatusActive`     | ACTIVE                                   |
| `BulkUploadHistoryEventStatusSuccessful` | SUCCESSFUL                               |