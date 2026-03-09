# ~~Role~~

DEPRECATED - use permissions instead. Viewer's role on the specific document.

> :warning: **DEPRECATED**: Deprecated on 2026-02-05, removal scheduled for 2026-10-15: Use permissions instead.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.RoleAnswerModerator

// Open enum: custom values can be created with a direct type cast
custom := components.Role("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `RoleAnswerModerator` | ANSWER_MODERATOR      |
| `RoleOwner`           | OWNER                 |
| `RoleViewer`          | VIEWER                |