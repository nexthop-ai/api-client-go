# AnnouncementChannel

This determines whether this is a Social Feed post or a regular announcement.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AnnouncementChannelMain

// Open enum: custom values can be created with a direct type cast
custom := components.AnnouncementChannel("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `AnnouncementChannelMain`       | MAIN                            |
| `AnnouncementChannelSocialFeed` | SOCIAL_FEED                     |