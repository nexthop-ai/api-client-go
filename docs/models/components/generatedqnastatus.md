# GeneratedQnaStatus

Status of backend generating the answer

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.GeneratedQnaStatusComputing

// Open enum: custom values can be created with a direct type cast
custom := components.GeneratedQnaStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `GeneratedQnaStatusComputing` | COMPUTING                     |
| `GeneratedQnaStatusDisabled`  | DISABLED                      |
| `GeneratedQnaStatusFailed`    | FAILED                        |
| `GeneratedQnaStatusNoAnswer`  | NO_ANSWER                     |
| `GeneratedQnaStatusSkipped`   | SKIPPED                       |
| `GeneratedQnaStatusStreaming` | STREAMING                     |
| `GeneratedQnaStatusSucceeded` | SUCCEEDED                     |
| `GeneratedQnaStatusTimeout`   | TIMEOUT                       |