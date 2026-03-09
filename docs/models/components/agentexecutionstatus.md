# AgentExecutionStatus

The status of the run. One of 'error', 'success'.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AgentExecutionStatusError

// Open enum: custom values can be created with a direct type cast
custom := components.AgentExecutionStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `AgentExecutionStatusError`   | error                         |
| `AgentExecutionStatusSuccess` | success                       |