# AgentEnum

Name of the agent.

## Example Usage

```go
import (
	"github.com/gleanwork/api-client-go/models/components"
)

value := components.AgentEnumDefault

// Open enum: custom values can be created with a direct type cast
custom := components.AgentEnum("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `AgentEnumDefault`   | DEFAULT              |
| `AgentEnumGpt`       | GPT                  |
| `AgentEnumUniversal` | UNIVERSAL            |
| `AgentEnumFast`      | FAST                 |
| `AgentEnumAdvanced`  | ADVANCED             |
| `AgentEnumAuto`      | AUTO                 |