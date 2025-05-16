# AgentRunCreate

Payload for creating a run.


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `AgentID`                                                                         | **string*                                                                         | :heavy_minus_sign:                                                                | The ID of the agent to run.                                                       |
| `Input`                                                                           | [*components.AgentRunCreateInput](../../models/components/agentruncreateinput.md) | :heavy_minus_sign:                                                                | The input to the agent.                                                           |
| `Messages`                                                                        | [][components.Message](../../models/components/message.md)                        | :heavy_minus_sign:                                                                | The messages to pass an input to the agent.                                       |