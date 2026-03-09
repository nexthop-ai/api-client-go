# DlpFrequency

Interval between scans. DAILY is deprecated.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DlpFrequencyOnce

// Open enum: custom values can be created with a direct type cast
custom := components.DlpFrequency("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `DlpFrequencyOnce`       | ONCE                     |
| `DlpFrequencyDaily`      | DAILY                    |
| `DlpFrequencyWeekly`     | WEEKLY                   |
| `DlpFrequencyContinuous` | CONTINUOUS               |
| `DlpFrequencyNone`       | NONE                     |