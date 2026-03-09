# AnnouncementPostType

This determines whether this is an external-link post or a regular announcement post. TEXT - Regular announcement that can contain rich text. LINK - Announcement that is linked to an external site.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AnnouncementPostTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.AnnouncementPostType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `AnnouncementPostTypeText` | TEXT                       |
| `AnnouncementPostTypeLink` | LINK                       |