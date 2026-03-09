# PersonTeamRelationship

The team member's relationship to the team. This defaults to MEMBER if not set.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.PersonTeamRelationshipMember

// Open enum: custom values can be created with a direct type cast
custom := components.PersonTeamRelationship("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `PersonTeamRelationshipMember`         | MEMBER                                 |
| `PersonTeamRelationshipManager`        | MANAGER                                |
| `PersonTeamRelationshipLead`           | LEAD                                   |
| `PersonTeamRelationshipPointOfContact` | POINT_OF_CONTACT                       |
| `PersonTeamRelationshipOther`          | OTHER                                  |