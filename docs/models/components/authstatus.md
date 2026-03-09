# AuthStatus

The per-user authorization status for a datasource.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AuthStatusDisabled

// Open enum: custom values can be created with a direct type cast
custom := components.AuthStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AuthStatusDisabled`     | DISABLED                 |
| `AuthStatusAwaitingAuth` | AWAITING_AUTH            |
| `AuthStatusAuthorized`   | AUTHORIZED               |
| `AuthStatusStaleOauth`   | STALE_OAUTH              |
| `AuthStatusSegMigration` | SEG_MIGRATION            |