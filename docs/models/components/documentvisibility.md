# DocumentVisibility

The level of visibility of the document as understood by our system.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DocumentVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.DocumentVisibility("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `DocumentVisibilityPrivate`                 | PRIVATE                                     |
| `DocumentVisibilitySpecificPeopleAndGroups` | SPECIFIC_PEOPLE_AND_GROUPS                  |
| `DocumentVisibilityDomainLink`              | DOMAIN_LINK                                 |
| `DocumentVisibilityDomainVisible`           | DOMAIN_VISIBLE                              |
| `DocumentVisibilityPublicLink`              | PUBLIC_LINK                                 |
| `DocumentVisibilityPublicVisible`           | PUBLIC_VISIBLE                              |