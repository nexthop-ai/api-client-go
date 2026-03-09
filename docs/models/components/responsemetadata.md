# ResponseMetadata

Metadata about the response (e.g., latency, token count).


## Fields

| Field                                                | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `LatencyMs`                                          | `*int64`                                             | :heavy_minus_sign:                                   | Time taken to generate the response in milliseconds. |
| `TokenCount`                                         | `*int64`                                             | :heavy_minus_sign:                                   | Number of tokens in the response.                    |
| `ModelUsed`                                          | `*string`                                            | :heavy_minus_sign:                                   | The specific model version used.                     |