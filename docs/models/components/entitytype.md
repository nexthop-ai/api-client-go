# EntityType

The type of entity.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.EntityTypePerson

// Open enum: custom values can be created with a direct type cast
custom := components.EntityType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `EntityTypePerson`   | PERSON               |
| `EntityTypeProject`  | PROJECT              |
| `EntityTypeCustomer` | CUSTOMER             |