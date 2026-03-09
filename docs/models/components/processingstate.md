# ProcessingState

The current state of the upload, an enum of UNAVAILABLE, UPLOAD STARTED, UPLOAD IN PROGRESS, UPLOAD COMPLETED, DELETION PAUSED, INDEXING COMPLETED

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ProcessingStateUnavailable

// Open enum: custom values can be created with a direct type cast
custom := components.ProcessingState("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ProcessingStateUnavailable`       | UNAVAILABLE                        |
| `ProcessingStateUploadStarted`     | UPLOAD STARTED                     |
| `ProcessingStateUploadInProgress`  | UPLOAD IN PROGRESS                 |
| `ProcessingStateUploadCompleted`   | UPLOAD COMPLETED                   |
| `ProcessingStateDeletionPaused`    | DELETION PAUSED                    |
| `ProcessingStateIndexingCompleted` | INDEXING COMPLETED                 |