package models

// Model supported by Groq Cloud. Full list available here: https://console.api.com/docs/models.
type Model string

const (
	// ModelDistilWhisperLargeV3EN available at https://huggingface.co/distil-whisper/distil-large-v3.
	ModelDistilWhisperLargeV3EN Model = "distil-whisper-large-v3-en"
	// ModelGemma29bIT available at https://huggingface.co/google/gemma-2-9b-it
	ModelGemma29bIT Model = "gemma2-9b-it"
	// ModelLlama3370BVersatile available at
	// https://github.com/meta-llama/llama-models/blob/main/models/llama3_3/MODEL_CARD.md
	ModelLlama3370BVersatile Model = "llama-3.3-70b-versatile"
	// ModelLLama318BInstant available at https://console.api.com/docs/model/llama-3.1-8b-instant
	ModelLLama318BInstant Model = "llama-3.1-8b-instant"
	// ModelLLamaGuard38B available at https://console.api.com/docs/model/llama-guard-3-8b
	ModelLLamaGuard38B Model = "llama-guard-3-8b"
	// ModelLLama70B8192 available at https://console.api.com/docs/model/llama3-70b-8192
	ModelLLama70B8192 Model = "llama3-70b-8192"
	// ModelLlama38b8192 available at https://huggingface.co/meta-llama/Meta-Llama-3-8B-Instruct
	ModelLlama38b8192 Model = "llama3-8b-8192"
	// ModelMixtral8x7B32768 available at https://huggingface.co/mistralai/Mixtral-8x7B-Instruct-v0.1
	ModelMixtral8x7B32768 Model = "mixtral-8x7b-32768"
	// ModelWhisperLargeV3 available at https://huggingface.co/openai/whisper-large-v3
	ModelWhisperLargeV3 Model = "whisper-large-v3"
	// ModelWhisperLargeV3Turbo available at https://huggingface.co/openai/whisper-large-v3-turbo
	ModelWhisperLargeV3Turbo Model = "whisper-large-v3-turbo"
)
