# AgentRun

Payload for creating a run.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `AgentID`                                                                           | **string*                                                                           | :heavy_minus_sign:                                                                  | The ID of the agent to run.                                                         |
| `Input`                                                                             | map[string]*any*                                                                    | :heavy_minus_sign:                                                                  | The input to the agent.                                                             |
| `Messages`                                                                          | [][components.Message](../../models/components/message.md)                          | :heavy_minus_sign:                                                                  | The messages to pass an input to the agent.                                         |
| `Status`                                                                            | [*components.AgentExecutionStatus](../../models/components/agentexecutionstatus.md) | :heavy_minus_sign:                                                                  | The status of the run. One of 'error', 'success'.                                   |