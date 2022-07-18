package context

import "context"

//NewCancelContext context with cancel created for example to respond to Kill signal from OS.
func NewCancelContext() (context.Context, context.CancelFunc) {
	// Create a new context
	ctx := context.Background()
	// Create a new context, with its cancellation function
	// from the original context
	ctx, cancel := context.WithCancel(ctx)

	return ctx, cancel
}
