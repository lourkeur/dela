package byzcoin

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.dedis.ch/fabric/encoding"
	"go.dedis.ch/fabric/internal/testing/fake"
	"go.dedis.ch/fabric/mino"
)

func TestIterator_HasNext(t *testing.T) {
	iter := &iterator{
		roster: &roster{addrs: make([]mino.Address, 3)},
	}

	require.True(t, iter.HasNext())

	iter.index = 1
	require.True(t, iter.HasNext())

	iter.index = 2
	require.True(t, iter.HasNext())

	iter.index = 3
	require.False(t, iter.HasNext())

	iter.index = 10
	require.False(t, iter.HasNext())
}

func TestIterator_GetNext(t *testing.T) {
	iter := &iterator{
		roster: &roster{addrs: make([]mino.Address, 3)},
	}

	for i := 0; i < 3; i++ {
		c := iter.GetNext()
		require.NotNil(t, c)
	}

	require.Equal(t, 3, iter.GetNext())
}

func TestAddressIterator_GetNext(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))
	iter := &addressIterator{
		iterator: &iterator{
			roster: &roster,
		},
	}

	for _, target := range roster.addrs {
		addr := iter.GetNext()
		require.Equal(t, target, addr)
	}

	require.Nil(t, iter.GetNext())
}

func TestPublicKeyIterator_GetNext(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))
	iter := &publicKeyIterator{
		iterator: &iterator{
			roster: &roster,
		},
	}

	for _, target := range roster.pubkeys {
		pubkey := iter.GetNext()
		require.Equal(t, target, pubkey)
	}

	require.Nil(t, iter.GetNext())
}

func TestRoster_Take(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))

	roster2 := roster.Take(mino.RangeFilter(1, 2))
	require.Equal(t, 1, roster2.Len())

	roster2 = roster.Take(mino.RangeFilter(1, 3))
	require.Equal(t, 2, roster2.Len())
}

func TestRoster_Len(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))
	require.Equal(t, 3, roster.Len())
}

func TestRoster_GetPublicKey(t *testing.T) {
	authority := fake.NewAuthority(3, fake.NewSigner)
	roster := rosterFactory{}.New(authority)

	iter := authority.AddressIterator()
	i := 0
	for iter.HasNext() {
		pubkey, index := roster.GetPublicKey(iter.GetNext())
		require.Equal(t, authority.GetSigner(i).GetPublicKey(), pubkey)
		require.Equal(t, i, index)
		i++
	}

	pubkey, index := roster.GetPublicKey(fake.NewAddress(999))
	require.Equal(t, -1, index)
	require.Nil(t, pubkey)
}

func TestRoster_Pack(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))

	rosterpb, err := roster.Pack(encoding.NewProtoEncoder())
	require.NoError(t, err)
	require.NotNil(t, rosterpb)

	roster.addrs[1] = fake.NewBadAddress()
	_, err = roster.Pack(encoding.NewProtoEncoder())
	require.EqualError(t, err, "couldn't marshal address: fake error")

	_, err = roster.Pack(fake.BadPackAnyEncoder{})
	require.EqualError(t, err, "couldn't pack public key: fake error")
}

func TestRosterFactory_FromProto(t *testing.T) {
	roster := rosterFactory{}.New(fake.NewAuthority(3, fake.NewSigner))
	rosterpb, err := roster.Pack(encoding.NewProtoEncoder())
	require.NoError(t, err)

	factory := newRosterFactory(fake.AddressFactory{}, fake.PublicKeyFactory{})

	decoded, err := factory.FromProto(rosterpb)
	require.NoError(t, err)
	require.Equal(t, roster.Len(), decoded.Len())

	_, err = factory.FromProto(nil)
	require.EqualError(t, err, "invalid message type '<nil>'")

	_, err = factory.FromProto(&Roster{Addresses: [][]byte{{}}})
	require.EqualError(t, err, "mismatch array length 1 != 0")

	factory.pubkeyFactory = fake.NewBadPublicKeyFactory()
	_, err = factory.FromProto(rosterpb)
	require.EqualError(t, err, "couldn't decode public key: fake error")
}
