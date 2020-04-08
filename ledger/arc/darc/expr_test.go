package darc

import (
	fmt "fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/fabric/ledger/arc"
	"golang.org/x/xerrors"
)

func TestExpression_Evolve(t *testing.T) {
	expr := newExpression()

	expr, err := expr.Evolve(nil)
	require.NoError(t, err)
	require.Len(t, expr.matches, 0)

	idents := []arc.Identity{
		fakeIdentity{buffer: []byte{0xaa}},
		fakeIdentity{buffer: []byte{0xbb}},
	}

	expr, err = expr.Evolve(idents)
	require.NoError(t, err)
	require.Len(t, expr.matches, 2)

	expr, err = expr.Evolve(idents)
	require.NoError(t, err)
	require.Len(t, expr.matches, 2)

	_, err = expr.Evolve([]arc.Identity{fakeIdentity{err: xerrors.New("oops")}})
	require.EqualError(t, err, "couldn't marshal identity: oops")
}

func TestExpression_Match(t *testing.T) {
	idents := []arc.Identity{
		fakeIdentity{buffer: []byte{0xaa}},
		fakeIdentity{buffer: []byte{0xbb}},
	}

	expr, err := newExpression().Evolve(idents)
	require.NoError(t, err)

	err = expr.Match(idents)
	require.NoError(t, err)

	err = expr.Match([]arc.Identity{fakeIdentity{buffer: []byte{0xcc}}})
	require.EqualError(t, err, "couldn't match identity '0xcc'")

	err = expr.Match([]arc.Identity{fakeIdentity{err: xerrors.New("oops")}})
	require.EqualError(t, err, "couldn't marshal identity: oops")
}

func TestExpression_Pack(t *testing.T) {
	idents := []arc.Identity{
		fakeIdentity{buffer: []byte{0xaa}},
		fakeIdentity{buffer: []byte{0xbb}},
	}

	expr, err := newExpression().Evolve(idents)
	require.NoError(t, err)

	pb, err := expr.Pack(nil)
	require.NoError(t, err)
	require.Len(t, pb.(*Expression).GetMatches(), 2)
}

// -----------------------------------------------------------------------------
// Utility functions

type fakeIdentity struct {
	arc.Identity
	buffer []byte
	err    error
}

func (i fakeIdentity) MarshalText() ([]byte, error) {
	return i.buffer, i.err
}

func (i fakeIdentity) String() string {
	return fmt.Sprintf("%#x", i.buffer)
}
