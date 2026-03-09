# StructuredResultSource

Source context for this result. Possible values depend on the result type.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.StructuredResultSourceExpertDetection

// Open enum: custom values can be created with a direct type cast
custom := components.StructuredResultSource("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `StructuredResultSourceExpertDetection` | EXPERT_DETECTION                        |
| `StructuredResultSourceEntityNlq`       | ENTITY_NLQ                              |
| `StructuredResultSourceCalendarEvent`   | CALENDAR_EVENT                          |
| `StructuredResultSourceAgent`           | AGENT                                   |