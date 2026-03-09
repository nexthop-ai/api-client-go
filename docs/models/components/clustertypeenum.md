# ClusterTypeEnum

The reason for inclusion of clusteredResults.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.ClusterTypeEnumSimilar

// Open enum: custom values can be created with a direct type cast
custom := components.ClusterTypeEnum("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ClusterTypeEnumSimilar`      | SIMILAR                       |
| `ClusterTypeEnumFreshness`    | FRESHNESS                     |
| `ClusterTypeEnumTitle`        | TITLE                         |
| `ClusterTypeEnumContent`      | CONTENT                       |
| `ClusterTypeEnumNone`         | NONE                          |
| `ClusterTypeEnumThreadReply`  | THREAD_REPLY                  |
| `ClusterTypeEnumThreadRoot`   | THREAD_ROOT                   |
| `ClusterTypeEnumPrefix`       | PREFIX                        |
| `ClusterTypeEnumSuffix`       | SUFFIX                        |
| `ClusterTypeEnumAuthorPrefix` | AUTHOR_PREFIX                 |
| `ClusterTypeEnumAuthorSuffix` | AUTHOR_SUFFIX                 |