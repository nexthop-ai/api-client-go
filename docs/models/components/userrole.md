# UserRole

A user's role with respect to a specific document.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.UserRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.UserRole("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `UserRoleOwner`           | OWNER                     |
| `UserRoleViewer`          | VIEWER                    |
| `UserRoleAnswerModerator` | ANSWER_MODERATOR          |
| `UserRoleEditor`          | EDITOR                    |
| `UserRoleVerifier`        | VERIFIER                  |