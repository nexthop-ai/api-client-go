# WriteActionType

Valid only for write actions. Represents the type of write action. REDIRECT - The client renders the URL which contains information for carrying out the action. EXECUTION - Send a request to an external server and execute the action. MCP - Send a tools/call request to an MCP server to execute the action.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.WriteActionTypeRedirect

// Open enum: custom values can be created with a direct type cast
custom := components.WriteActionType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `WriteActionTypeRedirect`  | REDIRECT                   |
| `WriteActionTypeExecution` | EXECUTION                  |
| `WriteActionTypeMcp`       | MCP                        |