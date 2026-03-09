# DocumentVisibilityOverrideOverride

The visibility-override state of the document.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DocumentVisibilityOverrideOverrideNone

// Open enum: custom values can be created with a direct type cast
custom := components.DocumentVisibilityOverrideOverride("custom_value")
```


## Values

| Name                                                       | Value                                                      |
| ---------------------------------------------------------- | ---------------------------------------------------------- |
| `DocumentVisibilityOverrideOverrideNone`                   | NONE                                                       |
| `DocumentVisibilityOverrideOverrideHideFromAll`            | HIDE_FROM_ALL                                              |
| `DocumentVisibilityOverrideOverrideHideFromGroups`         | HIDE_FROM_GROUPS                                           |
| `DocumentVisibilityOverrideOverrideHideFromAllExceptOwner` | HIDE_FROM_ALL_EXCEPT_OWNER                                 |