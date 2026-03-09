# DatasourcesType

The types of datasource for which to run the report/policy.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.DatasourcesTypeAll

// Open enum: custom values can be created with a direct type cast
custom := components.DatasourcesType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `DatasourcesTypeAll`    | ALL                     |
| `DatasourcesTypeCustom` | CUSTOM                  |