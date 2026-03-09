# Event

The action the user took within a Glean client with respect to the object referred to by the given `trackingToken`.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.EventClick
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `EventClick`                    | CLICK                           |
| `EventContainerClick`           | CONTAINER_CLICK                 |
| `EventCopyLink`                 | COPY_LINK                       |
| `EventCreate`                   | CREATE                          |
| `EventDismiss`                  | DISMISS                         |
| `EventDownvote`                 | DOWNVOTE                        |
| `EventEmail`                    | EMAIL                           |
| `EventExecute`                  | EXECUTE                         |
| `EventFilter`                   | FILTER                          |
| `EventFirstToken`               | FIRST_TOKEN                     |
| `EventFocusIn`                  | FOCUS_IN                        |
| `EventLastToken`                | LAST_TOKEN                      |
| `EventManualFeedback`           | MANUAL_FEEDBACK                 |
| `EventManualFeedbackSideBySide` | MANUAL_FEEDBACK_SIDE_BY_SIDE    |
| `EventFeedbackTimeSaved`        | FEEDBACK_TIME_SAVED             |
| `EventMarkAsRead`               | MARK_AS_READ                    |
| `EventMessage`                  | MESSAGE                         |
| `EventMiddleClick`              | MIDDLE_CLICK                    |
| `EventPageBlur`                 | PAGE_BLUR                       |
| `EventPageFocus`                | PAGE_FOCUS                      |
| `EventPageLeave`                | PAGE_LEAVE                      |
| `EventPreview`                  | PREVIEW                         |
| `EventRelatedClick`             | RELATED_CLICK                   |
| `EventRightClick`               | RIGHT_CLICK                     |
| `EventSectionClick`             | SECTION_CLICK                   |
| `EventSeen`                     | SEEN                            |
| `EventSelect`                   | SELECT                          |
| `EventShare`                    | SHARE                           |
| `EventShowMore`                 | SHOW_MORE                       |
| `EventUpvote`                   | UPVOTE                          |
| `EventView`                     | VIEW                            |
| `EventVisible`                  | VISIBLE                         |