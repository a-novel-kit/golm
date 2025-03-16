package models

import (
	"github.com/samber/lo"
)

// ChatCompletionRequest creates a model response for the given chat conversation.
//
// https://console.api.com/docs/api-reference#chat
type ChatCompletionRequest struct {
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the
	// text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *FrequencyPenalty `json:"frequency_penalty,omitempty"`

	// This is not yet supported by any of our  Whether to return log probabilities of the output tokens or not.
	// If true, returns the log probabilities of each output token returned in the content of message.
	LogProbs *LogProbs `json:"logprobs,omitempty"`

	// The maximum number of tokens that can be generated in the chat completion. The total length of input tokens and
	// generated tokens is limited by the model's context length.
	MaxCompletionTokens *MaxCompletionTokens `json:"max_completion_tokens,omitempty"`

	// A list of messages comprising the conversation so far.
	Messages Messages `json:"messages"`

	// ID of the model to use. For details on which models are compatible with the Chat API, see available models
	// https://console.api.com/docs/
	Model Model `json:"model"`

	// How many chat completion choices to generate for each input message. Note that the current moment, only n=1 is
	// supported. Other values will result in a 400 response.
	//
	// Defaults to 1.
	SuggestionsCount *SuggestionsCount `json:"n,omitempty"`

	// Whether to enable parallel function calling during tool use.
	// Defaults to true.
	ParallelToolCalls *ParallelToolCalls `json:"parallel_tool_calls,omitempty"`

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so
	// far, increasing the model's likelihood to talk about new topics.
	PresencePenalty *PresencePenalty `json:"presence_penalty,omitempty"`

	// Specifies how to output reasoning tokens.
	ReasoningFormat ReasoningFormat `json:"reasoning_format,omitempty"`

	// An object specifying the format that the model must output.
	//
	// Setting to { "type": "json_object" } enables JSON mode, which guarantees the message the model generates is
	// valid JSON.
	//
	// Important: when using JSON mode, you must also instruct the model to produce JSON yourself via a system or user
	// message.
	ResponseFormat ResponseFormat `json:"response_format,omitempty"`

	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with
	// the same seed and parameters should return the same result. Determinism is not guaranteed, and you should refer
	// to the system_fingerprint response parameter to monitor changes in the backend.
	Seed *Seed `json:"seed,omitempty"`

	// The service tier to use for the request. Defaults to ServiceTierOnDemand.
	//
	// ServiceTierAuto will automatically select the highest tier available within the rate limits of your organization.
	// ServiceTierFlex uses the flex tier, which will succeed or fail quickly.
	ServiceTier ServiceTier `json:"service_tier,omitempty"`

	// Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the
	// stop sequence.
	Stop *Stop `json:"stop,omitempty"`

	// If set, partial message deltas will be sent. Tokens will be sent as data-only server-sent events as they become
	// available, with the stream terminated by a data: [DONE] message.
	Stream *Stream `json:"stream,omitempty"`

	// Options for streaming response. Only set this when you set Stream: true.
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`

	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random,
	// while lower values like 0.2 will make it more focused and deterministic. We generally recommend altering this
	// or top_p but not both.
	Temperature *Temperature `json:"temperature,omitempty"`

	// Controls which (if any) tool is called by the model. ToolChoiceStaticNone means the model will not call
	// any tool and instead generates a message. ToolChoiceStaticAuto means the model can pick between generating a
	// message or calling one or more tools. Required means the model must call one or more tools. Specifying a
	// particular tool via {"type": "function", "function": {"name": "my_function"}} forces the model to call that tool.
	//
	// ToolChoiceStaticNone is the default when no tools are present. auto is the default if tools are present.\
	ToolChoice *ToolChoice `json:"tool_choice,omitempty"`

	// A list of tools the model may call. Currently, only functions are supported as a tool. Use this to provide a
	// list of functions the model may generate JSON inputs for. A max of 128 functions are supported.
	Tools []Tool `json:"tools,omitempty"`

	// This is not yet supported by any of our  An integer between 0 and 20 specifying the number of most
	// likely tokens to return at each token position, each with an associated log probability. logprobs must be set
	// to true if this parameter is used.
	TopLogProbs *TopLogProbs `json:"top_logprobs,omitempty"`

	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of
	// the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass
	// are considered. We generally recommend altering this or temperature but not both.
	TopP *TopP `json:"top_p,omitempty"`

	// A unique identifier representing your end-user, which can help us monitor and detect abuse.
	User string `json:"user,omitempty"`
}

// FrequencyPenalty is a number between -2.0 and 2.0. Positive values penalize new tokens based on their existing
// frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
type FrequencyPenalty float64

func NewFrequencyPenalty(f float64) *FrequencyPenalty {
	return lo.ToPtr(FrequencyPenalty(f))
}

// LogProbs determines whether to return log probabilities of the output tokens or not.
// If true, returns the log probabilities of each output token returned in the content of message.
type LogProbs bool

func NewLogProbs(l bool) *LogProbs {
	return lo.ToPtr(LogProbs(l))
}

// MaxCompletionTokens is the maximum number of tokens that can be generated in the chat completion. The total
// length of input tokens and generated tokens is limited by the model's context length.
type MaxCompletionTokens int

func NewMaxCompletionTokens(m int) *MaxCompletionTokens {
	return lo.ToPtr(MaxCompletionTokens(m))
}

// SuggestionsCount sets how many chat completion choices to generate for each input message. Note that the current
// moment, only n=1 is supported. Other values will result in a 400 response.
type SuggestionsCount int

func NewSuggestionsCount(s int) *SuggestionsCount {
	return lo.ToPtr(SuggestionsCount(s))
}

// ParallelToolCalls enables parallel function calling during tool use.
type ParallelToolCalls bool

func NewParallelToolCalls(p bool) *ParallelToolCalls {
	return lo.ToPtr(ParallelToolCalls(p))
}

// PresencePenalty is a number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear
// in the text so far, increasing the model's likelihood to talk about new topics.
type PresencePenalty float64

func NewPresencePenalty(p float64) *PresencePenalty {
	return lo.ToPtr(PresencePenalty(p))
}

// ReasoningFormat specifies how to output reasoning tokens.
type ReasoningFormat string

// Seed of the generation. If specified, our system will make a best effort to sample deterministically, such that
// repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed,
// band you should refer to the system_fingerprint response parameter to monitor changes in the backend.
type Seed int

func NewSeed(s int) *Seed {
	return lo.ToPtr(Seed(s))
}

// Temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values
// like 0.2 will make it more focused and deterministic. We generally recommend altering this or top_p but not both.
type Temperature float64

func NewTemperature(t float64) *Temperature {
	return lo.ToPtr(Temperature(t))
}

// TopLogProbs is an integer between 0 and 20 specifying the number of most likely tokens to return at each token
// position, each with an associated log probability. logprobs must be set to true if this parameter is used.
type TopLogProbs int

func NewTopLogProbs(t int) *TopLogProbs {
	return lo.ToPtr(TopLogProbs(t))
}

// TopP is an alternative to sampling with temperature, called nucleus sampling, where the model considers the
// results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability
// mass are considered. We generally recommend altering this or temperature but not both.
type TopP float64

func NewTopP(t float64) *TopP {
	return lo.ToPtr(TopP(t))
}
