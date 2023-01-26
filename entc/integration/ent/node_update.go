// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/node"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks     []Hook
	mutation  *NodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetValue sets the "value" field.
func (nu *NodeUpdate) SetValue(i int) *NodeUpdate {
	nu.mutation.ResetValue()
	nu.mutation.SetValue(i)
	return nu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableValue(i *int) *NodeUpdate {
	if i != nil {
		nu.SetValue(*i)
	}
	return nu
}

// AddValue adds i to the "value" field.
func (nu *NodeUpdate) AddValue(i int) *NodeUpdate {
	nu.mutation.AddValue(i)
	return nu
}

// ClearValue clears the value of the "value" field.
func (nu *NodeUpdate) ClearValue() *NodeUpdate {
	nu.mutation.ClearValue()
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nu *NodeUpdate) ClearUpdatedAt() *NodeUpdate {
	nu.mutation.ClearUpdatedAt()
	return nu
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nu *NodeUpdate) SetPrevID(id int) *NodeUpdate {
	nu.mutation.SetPrevID(id)
	return nu
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nu *NodeUpdate) SetNillablePrevID(id *int) *NodeUpdate {
	if id != nil {
		nu = nu.SetPrevID(*id)
	}
	return nu
}

// SetPrev sets the "prev" edge to the Node entity.
func (nu *NodeUpdate) SetPrev(n *Node) *NodeUpdate {
	return nu.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nu *NodeUpdate) SetNextID(id int) *NodeUpdate {
	nu.mutation.SetNextID(id)
	return nu
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nu *NodeUpdate) SetNillableNextID(id *int) *NodeUpdate {
	if id != nil {
		nu = nu.SetNextID(*id)
	}
	return nu
}

// SetNext sets the "next" edge to the Node entity.
func (nu *NodeUpdate) SetNext(n *Node) *NodeUpdate {
	return nu.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (nu *NodeUpdate) ClearPrev() *NodeUpdate {
	nu.mutation.ClearPrev()
	return nu
}

// ClearNext clears the "next" edge to the Node entity.
func (nu *NodeUpdate) ClearNext() *NodeUpdate {
	nu.mutation.ClearNext()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks[int, NodeMutation](ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok && !nu.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nu *NodeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NodeUpdate {
	nu.modifiers = append(nu.modifiers, modifiers...)
	return nu
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Value(); ok {
		_spec.SetField(node.FieldValue, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedValue(); ok {
		_spec.AddField(node.FieldValue, field.TypeInt, value)
	}
	if nu.mutation.ValueCleared() {
		_spec.ClearField(node.FieldValue, field.TypeInt)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	if nu.mutation.UpdatedAtCleared() {
		_spec.ClearField(node.FieldUpdatedAt, field.TypeTime)
	}
	if nu.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(nu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *NodeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetValue sets the "value" field.
func (nuo *NodeUpdateOne) SetValue(i int) *NodeUpdateOne {
	nuo.mutation.ResetValue()
	nuo.mutation.SetValue(i)
	return nuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableValue(i *int) *NodeUpdateOne {
	if i != nil {
		nuo.SetValue(*i)
	}
	return nuo
}

// AddValue adds i to the "value" field.
func (nuo *NodeUpdateOne) AddValue(i int) *NodeUpdateOne {
	nuo.mutation.AddValue(i)
	return nuo
}

// ClearValue clears the value of the "value" field.
func (nuo *NodeUpdateOne) ClearValue() *NodeUpdateOne {
	nuo.mutation.ClearValue()
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nuo *NodeUpdateOne) ClearUpdatedAt() *NodeUpdateOne {
	nuo.mutation.ClearUpdatedAt()
	return nuo
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nuo *NodeUpdateOne) SetPrevID(id int) *NodeUpdateOne {
	nuo.mutation.SetPrevID(id)
	return nuo
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillablePrevID(id *int) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetPrevID(*id)
	}
	return nuo
}

// SetPrev sets the "prev" edge to the Node entity.
func (nuo *NodeUpdateOne) SetPrev(n *Node) *NodeUpdateOne {
	return nuo.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nuo *NodeUpdateOne) SetNextID(id int) *NodeUpdateOne {
	nuo.mutation.SetNextID(id)
	return nuo
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableNextID(id *int) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetNextID(*id)
	}
	return nuo
}

// SetNext sets the "next" edge to the Node entity.
func (nuo *NodeUpdateOne) SetNext(n *Node) *NodeUpdateOne {
	return nuo.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (nuo *NodeUpdateOne) ClearPrev() *NodeUpdateOne {
	nuo.mutation.ClearPrev()
	return nuo
}

// ClearNext clears the "next" edge to the Node entity.
func (nuo *NodeUpdateOne) ClearNext() *NodeUpdateOne {
	nuo.mutation.ClearNext()
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	nuo.defaults()
	return withHooks[*Node, NodeMutation](ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok && !nuo.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (nuo *NodeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *NodeUpdateOne {
	nuo.modifiers = append(nuo.modifiers, modifiers...)
	return nuo
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   node.Table,
			Columns: node.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Value(); ok {
		_spec.SetField(node.FieldValue, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedValue(); ok {
		_spec.AddField(node.FieldValue, field.TypeInt, value)
	}
	if nuo.mutation.ValueCleared() {
		_spec.ClearField(node.FieldValue, field.TypeInt)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	if nuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(node.FieldUpdatedAt, field.TypeTime)
	}
	if nuo.mutation.PrevCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.PrevIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.NextCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.NextIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(nuo.modifiers...)
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
