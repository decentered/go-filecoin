// Package actor implements tooling to write and manipulate actors in go.
package actor

import (
	"fmt"
	"gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	cbor "gx/ipfs/QmRoARq3nkUb13HSKZGepCZSWe5GrVPwx7xURJGZ7KWv9V/go-ipld-cbor"
	"gx/ipfs/QmVmDhyTTUcQXFD1rRQ64fGLMSAoaQvNH3hwuaCFAPq2hy/errors"

	"github.com/filecoin-project/go-filecoin/types"
)

func init() {
	cbor.RegisterCborType(Actor{})
}

var (
	// ErrInvalidActorLength is returned when the actor length does not match the expected length.
	ErrInvalidActorLength = errors.New("invalid actor length")
)

// Actor is the central abstraction of entities in the system.
//
// Both individual accounts, as well as contracts (user & system level) are
// represented as actors. An actor has the following core functionality implemented on a system level:
// - track a Filecoin balance, using the `Balance` field
// - execute code stored in the `Code` field
// - read & write memory
// - replay protection, using the `Nonce` field
//
// Value sent to a non-existent address will be tracked as an empty actor that has a Balance but
// nil Code and Memory. You must nil check Code cids before comparing them.
//
// More specific capabilities for individual accounts or contract specific must be implemented
// inside the code.
//
// Not safe for concurrent access.
type Actor struct {
	Code    cid.Cid `refmt:",omitempty"`
	Head    cid.Cid `refmt:",omitempty"`
	Nonce   types.Uint64
	Balance *types.AttoFIL
}

// IncNonce increments the nonce of this actor by 1.
func (a *Actor) IncNonce() {
	a.Nonce = a.Nonce + 1
}

// Cid returns the canonical CID for the actor.
// TODO: can we avoid returning an error?
func (a *Actor) Cid() (cid.Cid, error) {
	obj, err := cbor.WrapObject(a, types.DefaultHashFunction, -1)
	if err != nil {
		return cid.Undef, errors.Wrap(err, "failed to marshal to cbor")
	}

	return obj.Cid(), nil
}

// NewActor constructs a new actor.
func NewActor(code cid.Cid, balance *types.AttoFIL) *Actor {
	return &Actor{
		Code:    code,
		Head:    cid.Undef,
		Nonce:   0,
		Balance: balance,
	}
}

// Unmarshal a actor from the given bytes.
func (a *Actor) Unmarshal(b []byte) error {
	return cbor.DecodeInto(b, a)
}

// Marshal the actor into bytes.
func (a *Actor) Marshal() ([]byte, error) {
	return cbor.DumpObject(a)
}

// Format implements fmt.Formatter.
func (a *Actor) Format(f fmt.State, c rune) {
	f.Write([]byte(fmt.Sprintf("<%s (%p); balance: %v; nonce: %d>", types.ActorCodeTypeName(a.Code), a, a.Balance, a.Nonce))) // nolint: errcheck
}
