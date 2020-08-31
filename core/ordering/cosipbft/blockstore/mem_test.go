package blockstore

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/core/ordering/cosipbft/types"
	"go.dedis.ch/dela/core/store"
	"go.dedis.ch/dela/core/validation/simple"
)

func TestInMemory_Len(t *testing.T) {
	store := NewInMemory()
	require.Equal(t, uint64(0), store.Len())

	store.blocks = []types.BlockLink{makeLink(t, types.Digest{}), makeLink(t, types.Digest{})}
	require.Equal(t, uint64(2), store.Len())
}

func TestInMemory_Store(t *testing.T) {
	store := NewInMemory()

	err := store.Store(makeLink(t, types.Digest{}))
	require.NoError(t, err)

	err = store.Store(makeLink(t, store.blocks[0].GetTo().GetHash()))
	require.NoError(t, err)

	err = store.Store(makeLink(t, types.Digest{}))
	require.EqualError(t, err, "mismatch link '00000000' != '2c34ce1d'")
}

func TestInMemory_Get(t *testing.T) {
	store := NewInMemory()

	store.blocks = []types.BlockLink{makeLink(t, types.Digest{})}

	block, err := store.Get(store.blocks[0].GetTo().GetHash())
	require.NoError(t, err)
	require.Equal(t, store.blocks[0], block)

	_, err = store.Get(types.Digest{})
	require.EqualError(t, err, "block not found: no block")
}

func TestInMemory_GetByIndex(t *testing.T) {
	store := NewInMemory()

	store.blocks = []types.BlockLink{
		makeLink(t, types.Digest{}, types.WithIndex(0)),
		makeLink(t, types.Digest{}, types.WithIndex(1)),
		makeLink(t, types.Digest{}, types.WithIndex(2)),
	}

	block, err := store.GetByIndex(1)
	require.NoError(t, err)
	require.Equal(t, uint64(1), block.GetTo().GetIndex())

	block, err = store.GetByIndex(2)
	require.NoError(t, err)
	require.Equal(t, uint64(2), block.GetTo().GetIndex())

	_, err = store.GetByIndex(3)
	require.EqualError(t, err, "block not found: no block")
}

func TestInMemory_Last(t *testing.T) {
	store := NewInMemory()

	_, err := store.Last()
	require.EqualError(t, err, "store empty: no block")

	store.blocks = []types.BlockLink{makeLink(t, types.Digest{})}
	block, err := store.Last()
	require.NoError(t, err)
	require.Equal(t, store.blocks[0], block)
}

func TestInMemory_Range(t *testing.T) {
	store := NewInMemory()

	count := 0
	store.Range(func(types.BlockLink) bool {
		count++
		return true
	})
	require.Equal(t, 0, count)

	store.blocks = []types.BlockLink{makeLink(t, types.Digest{}), makeLink(t, types.Digest{})}
	store.Range(func(types.BlockLink) bool {
		count++
		return false
	})
	require.Equal(t, 1, count)

	store.Range(func(types.BlockLink) bool {
		count++
		return true
	})
	require.Equal(t, 3, count)
}

func TestInMemory_Watch(t *testing.T) {
	store := NewInMemory()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := store.Watch(ctx)

	store.Store(makeLink(t, types.Digest{}))

	link := <-ch
	require.Equal(t, types.Digest{}, link.GetFrom())

	cancel()
	_, more := <-ch
	require.False(t, more)
}

func TestInMemory_WithTx(t *testing.T) {
	store := NewInMemory()

	tx := &fakeTx{}
	txstore := store.WithTx(tx)

	err := txstore.Store(makeLink(t, types.Digest{}))
	require.NoError(t, err)
	require.Len(t, store.blocks, 0)
	require.Len(t, txstore.(*InMemory).blocks, 1)

	tx.fn()
	require.Len(t, store.blocks, 1)
}

// Utility functions -----------------------------------------------------------

func makeLink(t *testing.T, from types.Digest, opts ...types.BlockOption) types.BlockLink {
	to, err := types.NewBlock(simple.NewData(nil), opts...)
	require.NoError(t, err)

	link, err := types.NewBlockLink(from, to)
	require.NoError(t, err)

	return link
}

type fakeTx struct {
	store.Transaction

	fn func()
}

func (tx *fakeTx) OnCommit(fn func()) {
	tx.fn = fn
}
