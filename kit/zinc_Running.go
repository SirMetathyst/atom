// Package kit ...
// generated by zinc/UniqueFlagComponentDataTemplate. DO NOT EDIT.
package kit

import (
	"github.com/SirMetathyst/zinc"
)

// ZRunning ...
var ZRunning uint = uint(343848780)

// RunningComponent ...
type RunningComponent struct {
	ctx  *zinc.ZContext
	data bool
}

// RegisterRunningComponentWith ...
func RegisterRunningComponentWith(e *zinc.ZEntityManager) {
	x := NewRunningComponent()
	ctx := e.RegisterComponent(ZRunning, x)
	x.SetContext(ctx)
}

// RegisterRunningComponent ...
func RegisterRunningComponent() {
	x := NewRunningComponent()
	ctx := zinc.Default().RegisterComponent(ZRunning, x)
	x.SetContext(ctx)
}

// NewRunningComponent ...
func NewRunningComponent() *RunningComponent {
	return &RunningComponent{data: false}
}

func init() {
	RegisterRunningComponent()
}

// SetContext ...
func (c *RunningComponent) SetContext(ctx *zinc.ZContext) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

// SetRunning ...
func (c *RunningComponent) SetRunning(v bool) {
	c.data = v
}

// IsRunning ...
func (c *RunningComponent) IsRunning() bool {
	return c.data
}

// HasEntity ...
func (c *RunningComponent) HasEntity(id zinc.ZEntityID) bool {
	return false
}

// DeleteEntity ...
func (c *RunningComponent) DeleteEntity(id zinc.ZEntityID) error {
	return nil
}

// SetRunningX ...
func SetRunningX(e *zinc.ZEntityManager) {
	v := e.Component(ZRunning)
	c := v.(*RunningComponent)
	c.SetRunning(true)
}

// SetRunning ...
func SetRunning() {
	SetRunningX(zinc.Default())
}

// IsRunningX ...
func IsRunningX(e *zinc.ZEntityManager) bool {
	v := e.Component(ZRunning)
	c := v.(*RunningComponent)
	return c.IsRunning()
}

// IsRunning ...
func IsRunning() bool {
	return IsRunningX(zinc.Default())
}

// NotRunningX ...
func NotRunningX(e *zinc.ZEntityManager) {
	v := e.Component(ZRunning)
	c := v.(*RunningComponent)
	c.SetRunning(false)
}

// NotRunning ...
func NotRunning() {
	NotRunningX(zinc.Default())
}