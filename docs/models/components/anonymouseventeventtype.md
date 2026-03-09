# AnonymousEventEventType

The nature of the event, for example "out of office".

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AnonymousEventEventTypeDefault

// Open enum: custom values can be created with a direct type cast
custom := components.AnonymousEventEventType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AnonymousEventEventTypeDefault`     | DEFAULT                              |
| `AnonymousEventEventTypeOutOfOffice` | OUT_OF_OFFICE                        |