# DlpSeverity

Severity levels for DLP findings and analyses.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpSeverityUnspecified

// Open enum: custom values can be created with a direct type cast
custom := components.DlpSeverity("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `DlpSeverityUnspecified` | UNSPECIFIED              |
| `DlpSeverityLow`         | LOW                      |
| `DlpSeverityMedium`      | MEDIUM                   |
| `DlpSeverityHigh`        | HIGH                     |