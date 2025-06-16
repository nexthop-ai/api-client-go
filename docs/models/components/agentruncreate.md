# AgentRunCreate

Payload for creating a run.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `AgentID`                                                  | *string*                                                   | :heavy_check_mark:                                         | The ID of the agent to run.                                |
| `Input`                                                    | map[string]*any*                                           | :heavy_minus_sign:                                         | The input to the agent.                                    |
| `Messages`                                                 | [][components.Message](../../models/components/message.md) | :heavy_minus_sign:                                         | The messages to pass an input to the agent.                |
| `Metadata`                                                 | map[string]*any*                                           | :heavy_minus_sign:                                         | The metadata to pass to the agent.                         |