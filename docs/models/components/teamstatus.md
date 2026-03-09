# TeamStatus

whether this team is fully processed or there are still unprocessed operations that'll affect it

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.TeamStatusProcessed

// Open enum: custom values can be created with a direct type cast
custom := components.TeamStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `TeamStatusProcessed`         | PROCESSED                     |
| `TeamStatusQueuedForCreation` | QUEUED_FOR_CREATION           |
| `TeamStatusQueuedForDeletion` | QUEUED_FOR_DELETION           |