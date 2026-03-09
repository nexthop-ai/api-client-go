# KnowledgeType

Indicates the kind of knowledge a tool would access or modify.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.KnowledgeTypeNeutralKnowledge

// Open enum: custom values can be created with a direct type cast
custom := components.KnowledgeType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `KnowledgeTypeNeutralKnowledge` | NEUTRAL_KNOWLEDGE               |
| `KnowledgeTypeCompanyKnowledge` | COMPANY_KNOWLEDGE               |
| `KnowledgeTypeWorldKnowledge`   | WORLD_KNOWLEDGE                 |