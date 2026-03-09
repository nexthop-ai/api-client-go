# GrantType

The type of grant type being used.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.GrantTypeAuthCode

// Open enum: custom values can be created with a direct type cast
custom := components.GrantType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GrantTypeAuthCode`          | AUTH_CODE                    |
| `GrantTypeClientCredentials` | CLIENT_CREDENTIALS           |