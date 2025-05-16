# AgentSchemas

Defines the structure and properties of an agent.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `AgentID`                                                          | *string*                                                           | :heavy_check_mark:                                                 | The ID of the agent.                                               |
| `InputSchema`                                                      | [components.InputSchema](../../models/components/inputschema.md)   | :heavy_check_mark:                                                 | The schema for the agent input. In JSON Schema format.             |
| `OutputSchema`                                                     | [components.OutputSchema](../../models/components/outputschema.md) | :heavy_check_mark:                                                 | The schema for the agent output. In JSON Schema format.            |