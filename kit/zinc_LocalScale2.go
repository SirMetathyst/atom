// Package kit ...
// Generated by the zinc tool.  DO NOT EDIT!
// Source: zinc_LocalScale2
package kit

import (
	"github.com/SirMetathyst/zinc"

)

// LocalScale2Key ...
const LocalScale2Key uint = 156111880

// LocalScale2Data ...
type LocalScale2Data struct {
	X float32
	Y float32	
}

// LocalScale2Component ...
type LocalScale2Component struct {
	ctx zinc.CTX
	data map[zinc.EntityID]LocalScale2Data
}

// NewLocalScale2Component ...
func NewLocalScale2Component() *LocalScale2Component {
	return &LocalScale2Component{
		data: make(map[zinc.EntityID]LocalScale2Data),
	}
}

// RegisterLocalScale2ComponentWith ...
func RegisterLocalScale2ComponentWith(e *zinc.EntityManager) {
	x := NewLocalScale2Component()
	ctx := e.RegisterComponent(LocalScale2Key, x)
	x.SetContext(ctx)
}

// RegisterLocalScale2Component ...
func RegisterLocalScale2Component() {
	x := NewLocalScale2Component()
	ctx := zinc.Default().RegisterComponent(LocalScale2Key, x)
	x.SetContext(ctx)
}

// SetContext ...
func (c *LocalScale2Component) SetContext(ctx zinc.CTX) {
	if c.ctx == nil {
		c.ctx = ctx
	}
}

func init() {
	RegisterLocalScale2Component()
}

// DeleteEntity ...
func (c *LocalScale2Component) DeleteEntity(id zinc.EntityID) {
	delete(c.data, id)
	c.ctx.ComponentDeleted(LocalScale2Key, id)
}

// HasEntity ...
func (c *LocalScale2Component) HasEntity(id zinc.EntityID) bool {
	_, ok := c.data[id]
	return ok
}

// SetLocalScale2 ...
func (c *LocalScale2Component) SetLocalScale2(id zinc.EntityID, localscale2 LocalScale2Data) {
	if c.ctx.HasEntity(id) {
		if c.HasEntity(id) {
			c.data[id] = localscale2
			c.ctx.ComponentUpdated(LocalScale2Key, id)
		} else {
			c.data[id] = localscale2
			c.ctx.ComponentAdded(LocalScale2Key, id)
		}
	}
}

// LocalScale2 ...
func (c *LocalScale2Component) LocalScale2(id zinc.EntityID) LocalScale2Data {
	return c.data[id]
}

// SetLocalScale2X ...
func SetLocalScale2X(e *zinc.EntityManager, id zinc.EntityID, localscale2 LocalScale2Data) {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	c.SetLocalScale2(id, localscale2)
}

// SetLocalScale2 ...
func SetLocalScale2(id zinc.EntityID, localscale2 LocalScale2Data) {
	SetLocalScale2X(zinc.Default(), id, localscale2)
}

// LocalScale2X ...
func LocalScale2X(e *zinc.EntityManager, id zinc.EntityID) LocalScale2Data {
	v, _ := e.Component(LocalScale2Key)
	c := v.(*LocalScale2Component)
	return c.LocalScale2(id)
}

// LocalScale2 ...
func LocalScale2(id zinc.EntityID) LocalScale2Data {
	return LocalScale2X(zinc.Default(), id)
}

// DeleteLocalScale2X ...
func DeleteLocalScale2X(e *zinc.EntityManager, id zinc.EntityID) {
	v, _ := e.Component(LocalPosition2Key)
	v.DeleteEntity(id)
}

// DeleteLocalScale2 ...
func DeleteLocalScale2(id zinc.EntityID) {
	DeleteLocalScale2X(zinc.Default(), id)
}

// HasLocalScale2X ...
func HasLocalScale2X(e *zinc.EntityManager, id zinc.EntityID) bool {
	v, _ := e.Component(LocalScale2Key)
	return v.HasEntity(id)
}

// HasLocalScale2 ...
func HasLocalScale2(id zinc.EntityID) bool {
	return HasLocalScale2X(zinc.Default(), id)
}