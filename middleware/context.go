package middleware

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"net/http"
)

var sessionProperties []string

// Context martini context
type Context struct {
	render.Render
	C        martini.Context
	S        sessions.Session
	R        *http.Request
	W        http.ResponseWriter
	FormErr  binding.Errors
	Messages []string
	Errors   []string
	Response map[string]interface{}
	Session  map[string]interface{}
}

func (context *Context) init() {
	if context.Response == nil {
		context.Response = make(map[string]interface{})
	}
	if context.Session == nil {
		context.Session = make(map[string]interface{})
	}
}

// SessionGet martini context get session interface
func (context *Context) SessionGet(key string) interface{} {
	return context.S.Get(key)
}

// SessionSet martini context set session
func (context *Context) SessionSet(key string, val interface{}) {
	context.init()
	context.S.Set(key, val)
	context.Session[key] = val
	for _, val := range sessionProperties {
		if val == key {
			return
		}
	}
	sessionProperties = append(sessionProperties, key)
}

// SessionDelete martini context delete session
func (context *Context) SessionDelete(key string) {
	delete(context.Response, key)
	context.S.Delete(key)
}

// SessionClear martini context clear session
func (context *Context) SessionClear() {
	context.Clear()
	context.S.Clear()
}

// Get martini context get value by key
func (context *Context) Get(key string) interface{} {
	return context.Response[key]
}

// Set martini context set key
func (context *Context) Set(key string, val interface{}) {
	context.init()
	context.Response[key] = val
}

// Delete martini context delete key
func (context *Context) Delete(key string) {
	delete(context.Response, key)
}

// Clear martini context clear all key
func (context *Context) Clear() {
	for key := range context.Response {
		context.Delete(key)
	}
}

// AddMessage martini context add message
func (context *Context) AddMessage(message string) {
	context.Messages = append(context.Messages, message)
}

// ClearMessages martini context clear all message
func (context *Context) ClearMessages() {
	context.Messages = context.Messages[:0]
}

// HasMessage martini context check if has message
func (context *Context) HasMessage() bool {
	return (len(context.Messages) > 0)
}

// SetFormErrors martini context set form binding error
func (context *Context) SetFormErrors(err binding.Errors) {
	context.FormErr = err
}

// JoinFormErrors martini context init
func (context *Context) JoinFormErrors(err binding.Errors) {
	context.init()
}

// AddError martini context add error
func (context *Context) AddError(err string) {
	context.Errors = append(context.Errors, err)
}

// AddFieldError martini context add field error
func (context *Context) AddFieldError(field string, err string) {
}

// ClearError martini context clear all error
func (context *Context) ClearError() {
	context.Errors = context.Errors[:0]
}

// HasError martini context check if has error
func (context *Context) HasError() bool {
	return context.HasCommonError() || context.HasFieldError() || context.HasOverallError()
}

// HasCommonError martini context check if has common error
func (context *Context) HasCommonError() bool {
	return (len(context.Errors) > 0)
}

// HasFieldError check if has field error
func (context *Context) HasFieldError() bool {
	return false
}

// HasOverallError check if has overall error
func (context *Context) HasOverallError() bool {
	return false
}

// OverallErrors get OverallErrors
func (context *Context) OverallErrors() map[string]string {
	return nil
}

// FieldErrors get FieldErrors
func (context *Context) FieldErrors() map[string]string {
	return nil
}

// InitContext init martini context
func InitContext() martini.Handler {
	return func(c martini.Context, s sessions.Session, rnd render.Render, r *http.Request, w http.ResponseWriter) {
		ctx := &Context{
			Render: rnd,
			W:      w,
			R:      r,
			C:      c,
			S:      s,
		}
		c.Map(ctx)
	}
}
