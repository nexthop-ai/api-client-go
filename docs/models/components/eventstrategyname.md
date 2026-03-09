# EventStrategyName

The name of method used to surface relevant data for a given calendar event.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.EventStrategyNameCustomerCard

// Open enum: custom values can be created with a direct type cast
custom := components.EventStrategyName("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `EventStrategyNameCustomerCard`      | customerCard                         |
| `EventStrategyNameNews`              | news                                 |
| `EventStrategyNameCall`              | call                                 |
| `EventStrategyNameEmail`             | email                                |
| `EventStrategyNameMeetingNotes`      | meetingNotes                         |
| `EventStrategyNameLinkedIn`          | linkedIn                             |
| `EventStrategyNameRelevantDocuments` | relevantDocuments                    |
| `EventStrategyNameChatFollowUps`     | chatFollowUps                        |
| `EventStrategyNameConversations`     | conversations                        |