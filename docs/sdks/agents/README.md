# Agents
(*Client.Agents*)

## Overview

### Available Operations

* [Retrieve](#retrieve) - Get Agent
* [RetrieveSchemas](#retrieveschemas) - Get Agent Schemas
* [List](#list) - Search Agents
* [RunStream](#runstream) - Create Run, Stream Output
* [Run](#run) - Create Run, Wait for Output

## Retrieve

Get an agent by ID. This endpoint implements the LangChain Agent Protocol, specifically part of the Agents stage (https://langchain-ai.github.io/agent-protocol/api.html#tag/agents/GET/agents/{agent_id}). It adheres to the standard contract defined for agent interoperability and can be used by agent runtimes that support the Agent Protocol.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Agents.Retrieve(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `agentID`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The ID of the agent.                                                                                       |
| `timezoneOffset`                                                                                           | **int64*                                                                                                   | :heavy_minus_sign:                                                                                         | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC. |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAgentResponse](../../models/operations/getagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveSchemas

Get an agent's schemas by ID. This endpoint implements the LangChain Agent Protocol, specifically part of the Agents stage (https://langchain-ai.github.io/agent-protocol/api.html#tag/agents/GET/agents/{agent_id}/schemas). It adheres to the standard contract defined for agent interoperability and can be used by agent runtimes that support the Agent Protocol.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Agents.RetrieveSchemas(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentSchemas != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `agentID`                                                                                                  | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The ID of the agent.                                                                                       |
| `timezoneOffset`                                                                                           | **int64*                                                                                                   | :heavy_minus_sign:                                                                                         | The offset of the client's timezone in minutes from UTC. e.g. PDT is -420 because it's 7 hours behind UTC. |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GetAgentSchemasResponse](../../models/operations/getagentschemasresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

List Agents available in this service. This endpoint implements the LangChain Agent Protocol, specifically part of the Agents stage (https://langchain-ai.github.io/agent-protocol/api.html#tag/agents/POST/agents/search). It adheres to the standard contract defined for agent interoperability and can be used by agent runtimes that support the Agent Protocol.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Agents.List(ctx, components.SearchAgentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchAgentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.SearchAgentsRequest](../../models/components/searchagentsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.SearchAgentsResponse](../../models/operations/searchagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RunStream

Creates and triggers a run of an agent. Streams the output in SSE format. This endpoint implements the LangChain Agent Protocol, specifically part of the Runs stage (https://langchain-ai.github.io/agent-protocol/api.html#tag/runs/POST/runs/stream). It adheres to the standard contract defined for agent interoperability and can be used by agent runtimes that support the Agent Protocol. Note that running agents that reference third party platform write actions is unsupported as it requires user confirmation.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Agents.RunStream(ctx, components.AgentRunCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.AgentRunCreate](../../models/components/agentruncreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateAndStreamRunResponse](../../models/operations/createandstreamrunresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Run

Creates and triggers a run of an agent. Waits for final output and then returns it. This endpoint implements the LangChain Agent Protocol, specifically part of the Runs stage (https://langchain-ai.github.io/agent-protocol/api.html#tag/runs/POST/runs/wait). It adheres to the standard contract defined for agent interoperability and can be used by agent runtimes that support the Agent Protocol. Note that running agents that reference third party platform write actions is unsupported as it requires user confirmation.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            ActAsBearerToken: apiclientgo.String(os.Getenv("GLEAN_ACT_AS_BEARER_TOKEN")),
        }),
    )

    res, err := s.Client.Agents.Run(ctx, components.AgentRunCreate{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentRunWaitResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.AgentRunCreate](../../models/components/agentruncreate.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.CreateAndWaitRunResponse](../../models/operations/createandwaitrunresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |