package golm

import "context"

type chatContext struct{}

// WithContext associates a new Chat instance with the current context, from the provided binding.
func WithContext[Req, Res, Stream any](ctx context.Context, binding ChatBinding[Req, Res, Stream]) context.Context {
	chat := NewChat(binding)

	return context.WithValue(ctx, chatContext{}, chat)
}

func ContextWithRaw[Req, Res any](ctx context.Context) Chat[Req, Res] {
	chat := ctx.Value(chatContext{}).(Chat[Req, Res])

	return chat
}

func Context(ctx context.Context) ChatBase {
	chat := ctx.Value(chatContext{})

	if chat == nil {
		return nil
	}

	return chat.(ChatBase)
}
