# DocumentVisibilityUpdateResultOverride

The visibility-override state of the document.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DocumentVisibilityUpdateResultOverrideNone

// Open enum: custom values can be created with a direct type cast
custom := components.DocumentVisibilityUpdateResultOverride("custom_value")
```


## Values

| Name                                                           | Value                                                          |
| -------------------------------------------------------------- | -------------------------------------------------------------- |
| `DocumentVisibilityUpdateResultOverrideNone`                   | NONE                                                           |
| `DocumentVisibilityUpdateResultOverrideHideFromAll`            | HIDE_FROM_ALL                                                  |
| `DocumentVisibilityUpdateResultOverrideHideFromGroups`         | HIDE_FROM_GROUPS                                               |
| `DocumentVisibilityUpdateResultOverrideHideFromAllExceptOwner` | HIDE_FROM_ALL_EXCEPT_OWNER                                     |