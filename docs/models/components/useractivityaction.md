# UserActivityAction

The action for the activity

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.UserActivityActionAdd

// Open enum: custom values can be created with a direct type cast
custom := components.UserActivityAction("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `UserActivityActionAdd`         | ADD                             |
| `UserActivityActionAddReminder` | ADD_REMINDER                    |
| `UserActivityActionClick`       | CLICK                           |
| `UserActivityActionComment`     | COMMENT                         |
| `UserActivityActionDelete`      | DELETE                          |
| `UserActivityActionDismiss`     | DISMISS                         |
| `UserActivityActionEdit`        | EDIT                            |
| `UserActivityActionMention`     | MENTION                         |
| `UserActivityActionMove`        | MOVE                            |
| `UserActivityActionOther`       | OTHER                           |
| `UserActivityActionRestore`     | RESTORE                         |
| `UserActivityActionUnknown`     | UNKNOWN                         |
| `UserActivityActionVerify`      | VERIFY                          |
| `UserActivityActionView`        | VIEW                            |