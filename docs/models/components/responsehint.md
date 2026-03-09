# ResponseHint

Hints for the response content.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ResponseHintAllResultCounts

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseHint("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ResponseHintAllResultCounts`    | ALL_RESULT_COUNTS                |
| `ResponseHintFacetResults`       | FACET_RESULTS                    |
| `ResponseHintQueryMetadata`      | QUERY_METADATA                   |
| `ResponseHintResults`            | RESULTS                          |
| `ResponseHintSpellcheckMetadata` | SPELLCHECK_METADATA              |