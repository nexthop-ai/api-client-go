# AgentCapabilities

Describes features that the agent supports. example: {
  "ap.io.messages": true,
  "ap.io.streaming": true
}


## Fields

| Field                                                                                                                         | Type                                                                                                                          | Required                                                                                                                      | Description                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `ApIoMessages`                                                                                                                | **bool*                                                                                                                       | :heavy_minus_sign:                                                                                                            | Whether the agent supports messages as an input. If true, you'll pass `messages` as an input when running the agent.          |
| `ApIoStreaming`                                                                                                               | **bool*                                                                                                                       | :heavy_minus_sign:                                                                                                            | Whether the agent supports streaming output. If true, you you can stream agent ouput. All agents currently support streaming. |
| `AdditionalProperties`                                                                                                        | map[string]*any*                                                                                                              | :heavy_minus_sign:                                                                                                            | N/A                                                                                                                           |