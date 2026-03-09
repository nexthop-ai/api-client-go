# PersonToTeamRelationshipRelationship

The team member's relationship to the team. This defaults to MEMBER if not set.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PersonToTeamRelationshipRelationshipMember

// Open enum: custom values can be created with a direct type cast
custom := components.PersonToTeamRelationshipRelationship("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `PersonToTeamRelationshipRelationshipMember`         | MEMBER                                               |
| `PersonToTeamRelationshipRelationshipManager`        | MANAGER                                              |
| `PersonToTeamRelationshipRelationshipLead`           | LEAD                                                 |
| `PersonToTeamRelationshipRelationshipPointOfContact` | POINT_OF_CONTACT                                     |
| `PersonToTeamRelationshipRelationshipOther`          | OTHER                                                |