# CalendarEventEventType

The nature of the event, for example "out of office".

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CalendarEventEventTypeDefault

// Open enum: custom values can be created with a direct type cast
custom := components.CalendarEventEventType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CalendarEventEventTypeDefault`     | DEFAULT                             |
| `CalendarEventEventTypeOutOfOffice` | OUT_OF_OFFICE                       |