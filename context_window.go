// +build !web

package wua

type WindowContext struct {
	*contextBase
	message *Message
}

func newContext(msg *Message) Context {
	return &WindowContext{&contextBase{data: []byte(msg.Data)}, msg}
}

func (ctx *WindowContext) Done(ok bool, feedback interface{}) {
	send(callback, ctx.message.Id, ok, feedback)
}

// feedback should be a primary type, or implement the fmt.Stringer interface
// if not, convert your value to string first. e.g. string(bytes)
func (ctx *WindowContext) Success(feedback interface{}) {
	Logger.Info("success:", feedback)
	ctx.Done(true, feedback)
}

func (ctx *WindowContext) Error(feedback ...interface{}) {
	ctx.Done(false, feedback)
}
