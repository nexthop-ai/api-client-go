# MessageType

Semantically groups content of a certain type. It can be used for purposes such as differential UI treatment. USER authored messages should be of type CONTENT and do not need `messageType` specified.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.MessageTypeUpdate

// Open enum: custom values can be created with a direct type cast
custom := components.MessageType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `MessageTypeUpdate`         | UPDATE                      |
| `MessageTypeContent`        | CONTENT                     |
| `MessageTypeContext`        | CONTEXT                     |
| `MessageTypeControl`        | CONTROL                     |
| `MessageTypeControlStart`   | CONTROL_START               |
| `MessageTypeControlFinish`  | CONTROL_FINISH              |
| `MessageTypeControlCancel`  | CONTROL_CANCEL              |
| `MessageTypeControlRetry`   | CONTROL_RETRY               |
| `MessageTypeControlUnknown` | CONTROL_UNKNOWN             |
| `MessageTypeDebug`          | DEBUG                       |
| `MessageTypeDebugExternal`  | DEBUG_EXTERNAL              |
| `MessageTypeError`          | ERROR                       |
| `MessageTypeHeading`        | HEADING                     |
| `MessageTypeWarning`        | WARNING                     |
| `MessageTypeServerTool`     | SERVER_TOOL                 |