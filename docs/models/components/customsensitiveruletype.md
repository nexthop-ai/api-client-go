# CustomSensitiveRuleType

Type of the custom sensitive rule.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CustomSensitiveRuleTypeRegex

// Open enum: custom values can be created with a direct type cast
custom := components.CustomSensitiveRuleType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CustomSensitiveRuleTypeRegex`    | REGEX                             |
| `CustomSensitiveRuleTypeTerm`     | TERM                              |
| `CustomSensitiveRuleTypeInfoType` | INFO_TYPE                         |