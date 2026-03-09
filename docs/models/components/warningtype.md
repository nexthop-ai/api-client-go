# WarningType

The type of the warning.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.WarningTypeLongQuery

// Open enum: custom values can be created with a direct type cast
custom := components.WarningType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `WarningTypeLongQuery`               | LONG_QUERY                           |
| `WarningTypeQuotedPunctuation`       | QUOTED_PUNCTUATION                   |
| `WarningTypePunctuationOnly`         | PUNCTUATION_ONLY                     |
| `WarningTypeCopypastedQuotes`        | COPYPASTED_QUOTES                    |
| `WarningTypeInvalidOperator`         | INVALID_OPERATOR                     |
| `WarningTypeMaybeInvalidFacetQuery`  | MAYBE_INVALID_FACET_QUERY            |
| `WarningTypeTooManyDatasourceGroups` | TOO_MANY_DATASOURCE_GROUPS           |