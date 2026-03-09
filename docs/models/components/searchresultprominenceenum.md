# SearchResultProminenceEnum

The level of visual distinction that should be given to a result.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.SearchResultProminenceEnumHero

// Open enum: custom values can be created with a direct type cast
custom := components.SearchResultProminenceEnum("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `SearchResultProminenceEnumHero`     | HERO                                 |
| `SearchResultProminenceEnumPromoted` | PROMOTED                             |
| `SearchResultProminenceEnumStandard` | STANDARD                             |