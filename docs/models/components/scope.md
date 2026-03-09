# Scope

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ScopeAppCard

// Open enum: custom values can be created with a direct type cast
custom := components.Scope("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ScopeAppCard`                | APP_CARD                      |
| `ScopeAutocompleteExactMatch` | AUTOCOMPLETE_EXACT_MATCH      |
| `ScopeAutocompleteFuzzyMatch` | AUTOCOMPLETE_FUZZY_MATCH      |
| `ScopeAutocompleteZeroQuery`  | AUTOCOMPLETE_ZERO_QUERY       |
| `ScopeNewTabPage`             | NEW_TAB_PAGE                  |