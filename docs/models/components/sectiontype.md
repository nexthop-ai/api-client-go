# SectionType

Type of the section. This defines how the section should be interpreted and rendered in the digest.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.SectionTypeChannel

// Open enum: custom values can be created with a direct type cast
custom := components.SectionType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `SectionTypeChannel`  | CHANNEL               |
| `SectionTypeMentions` | MENTIONS              |
| `SectionTypeTopic`    | TOPIC                 |