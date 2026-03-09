# TimeRangeFilterTimePeriodType

The type of time period for which to filter findings.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.TimeRangeFilterTimePeriodTypePastDay

// Open enum: custom values can be created with a direct type cast
custom := components.TimeRangeFilterTimePeriodType("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `TimeRangeFilterTimePeriodTypePastDay`   | PAST_DAY                                 |
| `TimeRangeFilterTimePeriodTypePastWeek`  | PAST_WEEK                                |
| `TimeRangeFilterTimePeriodTypePastMonth` | PAST_MONTH                               |
| `TimeRangeFilterTimePeriodTypePastYear`  | PAST_YEAR                                |
| `TimeRangeFilterTimePeriodTypeCustom`    | CUSTOM                                   |