# InputOptionsTimePeriodType

Type of time period for which to run the report/policy. PAST_DAY is deprecated.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.InputOptionsTimePeriodTypeAllTime

// Open enum: custom values can be created with a direct type cast
custom := components.InputOptionsTimePeriodType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `InputOptionsTimePeriodTypeAllTime`   | ALL_TIME                              |
| `InputOptionsTimePeriodTypePastYear`  | PAST_YEAR                             |
| `InputOptionsTimePeriodTypePastDay`   | PAST_DAY                              |
| `InputOptionsTimePeriodTypeCustom`    | CUSTOM                                |
| `InputOptionsTimePeriodTypeLastNDays` | LAST_N_DAYS                           |