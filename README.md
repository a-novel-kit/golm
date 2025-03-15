# Go-LM

Go bindings and API for popular LLM services.

```bash
go get github.com/a-novel-kit/golm
```

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/a-novel-kit/golm/main.yaml)
[![codecov](https://codecov.io/gh/a-novel-kit/golm/graph/badge.svg?token=FlJyQ4XbHn)](https://codecov.io/gh/a-novel-kit/golm)
[![Go Report Card](https://goreportcard.com/badge/github.com/a-novel-kit/golm)](https://goreportcard.com/report/github.com/a-novel-kit/golm)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/a-novel-kit/golm)
![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/a-novel-kit/golm)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/a-novel-kit/golm)

![Coverage graph](https://codecov.io/gh/a-novel-kit/golm/graphs/sunburst.svg?token=FlJyQ4XbHn)

# Chat completion

Available bindings:

- [Groq cloud](https://console.api.com/docs/api-reference)

## Configure a binding

You can use bindings available from this repository to configure your Groq client.

```go
package main

import (
	"os"

	"github.com/a-novel-kit/golm/bindings/groq"
	"github.com/a-novel-kit/golm/bindings/groq/models"
)

func main() {
	binding := groq.New(os.Getenv("GROQ_TOKEN"), models.ModelLlama3370BVersatile)
}
```

Once you have a binding configured, you can use it to initialize a new chat instance.

```go
package main

import (
	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)
}
```

## Make a simple completion request

```go
package main

import (
	"log"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	message := golm.NewUserMessage("Count from 1 to ten.")

	resp, _ := chat.Completion(context.Background(), message, golm.CompletionParams{})

	// [Assistant]:
	// Here is a list of numbers from 1 to 10:
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	log.Println(resp)
}
```

You can set a system message to bias the behavior of the completion.

```go
package main

import (
	"log"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	chat.SetSystem(golm.NewSystemMessage(
		"You are a counter. Just count. Return your answer as a list of " +
			"numbers separated by commas. Example: 1,2,3",
    ))

	message := golm.NewUserMessage("Count from 1 to ten.")

	resp, _ := chat.Completion(context.Background(), message, golm.CompletionParams{})

	// [Assistant]:
	// 1,2,3,4,5,6,7,8,9,10
	log.Println(resp)
}
```

## JSON response

You can instruct your model to return JSON, and let the binding parse it for you.

> Smaller models might have trouble returning proper JSON, and it might not be supported well by all APIs.

```go
package main

import (
	"context"
	"log"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	// Most models will require you to specify the JSON format,
	// somewhere in the chat history.
	chat.SetSystem(golm.NewSystemMessage(
		"You are a counter. Just count. Return your answer as a JSON array of " +
			"numbers, separated by commas. Wrap your array in the key \"count\". " +
			"Example: { \"count\": [1,2,3] }",
	))

	message := golm.NewUserMessage("Count from 1 to ten.")

	var jsonOutput struct {
		Count int `json:"count"`
	}

	_ = chat.CompletionJSON(context.Background(), message, golm.CompletionParams{}, &jsonOutput)

	// { "count": [1,2,3,4,5,6,7,8,9,10] }
	log.Println(jsonOutput)
}
```

## Streaming

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	chat.SetSystem(golm.NewSystemMessage(
		"You are a counter. Just count. Return your answer as a list of " +
			"numbers separated by commas. Example: 1,2,3",
	))

	message := golm.NewUserMessage("Count from 1 to ten.")

	// Make sure to have some security to prevent the request from
	// hanging. The chat implementation should support cancelable
	// contexts.
	timedCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	outC, terminate := chat.CompletionStream(timedCtx, message, golm.CompletionParams{})

	var output string

	for elem := range outC {
		output += elem
	}

	if err := terminate(); err != nil {
        panic(err)
    }

	// 1,2,3,4,5,6,7,8,9,10
	log.Println(output)
}
```

## Manipulating history

When you call the completion API, the message history is automatically updated.

```go
package main

import (
	"log"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	message := golm.NewUserMessage("Count from 1 to ten.")

	_, _ = chat.Completion(context.Background(), message, golm.CompletionParams{})

	// This query is aware of the previous request / answer.
	message = golm.NewUserMessage("Keep counting up to 20")
	resp, _ := chat.Completion(context.Background(), message, golm.CompletionParams{})

	// [Assistant]:
	// 11,12,13,14,15,16,17,18,19,20
	log.Println(resp)
}
```

You can change this history / add messages manually.

```go
package main

import (
	"context"
	"log"

	"github.com/a-novel-kit/golm"
)

func main() {
	chat := golm.NewChat(binding)

	chat.PushHistory(
		golm.NewUserMessage("Count from 1 to ten."),
		golm.NewAssistantMessage("one, two, three, four, five, six, seven, eight, nine, ten"),
	)

	message := golm.NewUserMessage("Keep counting up to 20")
	resp, _ := chat.Completion(context.Background(), message, golm.CompletionParams{})

	// [Assistant]:
	// eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen, eighteen, nineteen, twenty
	log.Println(resp)
}
```
