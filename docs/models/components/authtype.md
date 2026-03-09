# AuthType

The type of authentication being used.
Use 'OAUTH_*' when Glean calls an external API (e.g., Jira) on behalf of a user to obtain an OAuth token.
'OAUTH_ADMIN' utilizes an admin token for external API calls on behalf all users.
'OAUTH_USER' uses individual user tokens for external API calls.
'DWD' refers to domain wide delegation.


## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AuthTypeNone

// Open enum: custom values can be created with a direct type cast
custom := components.AuthType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `AuthTypeNone`       | NONE                 |
| `AuthTypeOauthUser`  | OAUTH_USER           |
| `AuthTypeOauthAdmin` | OAUTH_ADMIN          |
| `AuthTypeAPIKey`     | API_KEY              |
| `AuthTypeBasicAuth`  | BASIC_AUTH           |
| `AuthTypeDwd`        | DWD                  |