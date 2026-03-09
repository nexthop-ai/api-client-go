# DatasourceVisibility

The visibility of the datasource, an enum of VISIBLE_TO_ALL, VISIBLE_TO_TEST_GROUP, NOT_VISIBLE

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DatasourceVisibilityEnabledForAll

// Open enum: custom values can be created with a direct type cast
custom := components.DatasourceVisibility("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `DatasourceVisibilityEnabledForAll`       | ENABLED_FOR_ALL                           |
| `DatasourceVisibilityEnabledForTestGroup` | ENABLED_FOR_TEST_GROUP                    |
| `DatasourceVisibilityNotEnabled`          | NOT_ENABLED                               |