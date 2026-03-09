# CustomSensitiveRuleLikelihoodThreshold

Likelihood threshold for BUILT_IN infotypes (e.g., LIKELY, VERY_LIKELY). Only applicable for BUILT_IN type.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.CustomSensitiveRuleLikelihoodThresholdLikely

// Open enum: custom values can be created with a direct type cast
custom := components.CustomSensitiveRuleLikelihoodThreshold("custom_value")
```


## Values

| Name                                                 | Value                                                |
| ---------------------------------------------------- | ---------------------------------------------------- |
| `CustomSensitiveRuleLikelihoodThresholdLikely`       | LIKELY                                               |
| `CustomSensitiveRuleLikelihoodThresholdVeryLikely`   | VERY_LIKELY                                          |
| `CustomSensitiveRuleLikelihoodThresholdPossible`     | POSSIBLE                                             |
| `CustomSensitiveRuleLikelihoodThresholdUnlikely`     | UNLIKELY                                             |
| `CustomSensitiveRuleLikelihoodThresholdVeryUnlikely` | VERY_UNLIKELY                                        |