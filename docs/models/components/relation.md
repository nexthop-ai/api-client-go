# Relation

How this document relates to the including entity.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.RelationAttachment

// Open enum: custom values can be created with a direct type cast
custom := components.Relation("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `RelationAttachment`           | ATTACHMENT                     |
| `RelationCanonical`            | CANONICAL                      |
| `RelationCase`                 | CASE                           |
| `RelationContactLower`         | contact                        |
| `RelationContactUpper`         | CONTACT                        |
| `RelationConversationMessages` | CONVERSATION_MESSAGES          |
| `RelationExpert`               | EXPERT                         |
| `RelationFrom`                 | FROM                           |
| `RelationHighlight`            | HIGHLIGHT                      |
| `RelationOpportunityLower`     | opportunity                    |
| `RelationOpportunityUpper`     | OPPORTUNITY                    |
| `RelationRecent`               | RECENT                         |
| `RelationSource`               | SOURCE                         |
| `RelationTicket`               | TICKET                         |
| `RelationTranscript`           | TRANSCRIPT                     |
| `RelationWith`                 | WITH                           |