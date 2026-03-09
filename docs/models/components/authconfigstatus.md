# AuthConfigStatus

Auth status of the tool.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AuthConfigStatusAwaitingAuth

// Open enum: custom values can be created with a direct type cast
custom := components.AuthConfigStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AuthConfigStatusAwaitingAuth` | AWAITING_AUTH                  |
| `AuthConfigStatusAuthorized`   | AUTHORIZED                     |
| `AuthConfigStatusAuthDisabled` | AUTH_DISABLED                  |