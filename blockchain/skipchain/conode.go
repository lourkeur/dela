package skipchain

import (
	"io"

	proto "github.com/golang/protobuf/proto"
	"go.dedis.ch/fabric/cosi"
	"go.dedis.ch/fabric/crypto"
	"go.dedis.ch/fabric/encoding"
	"go.dedis.ch/fabric/mino"
	"golang.org/x/xerrors"
)

// Conode is the type of participant for a skipchain. It contains an address
// and a public key that is part of the key pair used to sign blocks.
//
// - implements encoding.Packable
type Conode struct {
	addr      mino.Address
	publicKey crypto.PublicKey
}

// GetAddress returns the address of the conode.
func (c Conode) GetAddress() mino.Address {
	return c.addr
}

// GetPublicKey returns the public key of the conode.
func (c Conode) GetPublicKey() crypto.PublicKey {
	return c.publicKey
}

// Pack implements encoding.Packable. It returns the protobuf message for the
// conode.
func (c Conode) Pack() (proto.Message, error) {
	packed, err := c.publicKey.Pack()
	if err != nil {
		return nil, encoding.NewEncodingError("public key", err)
	}

	conode := &ConodeProto{}

	conode.Address, err = c.addr.MarshalText()
	if err != nil {
		return nil, encoding.NewEncodingError("address", err)
	}

	conode.PublicKey, err = protoenc.MarshalAny(packed)
	if err != nil {
		return nil, encoding.NewAnyEncodingError(packed, err)
	}

	return conode, nil
}

// iterator is a generic implementation of an iterator over a list of conodes.
type iterator struct {
	conodes []Conode
	index   int
}

func (i *iterator) HasNext() bool {
	if i.index+1 < len(i.conodes) {
		return true
	}
	return false
}

func (i *iterator) GetNext() *Conode {
	if i.HasNext() {
		i.index++
		return &i.conodes[i.index]
	}
	return nil
}

// addressIterator is an iterator for a list of addresses.
//
// - implements mino.AddressIterator
type addressIterator struct {
	*iterator
}

// GetNext implements mino.AddressIterator. It returns the next address.
func (i *addressIterator) GetNext() mino.Address {
	conode := i.iterator.GetNext()
	if conode != nil {
		return conode.GetAddress()
	}
	return nil
}

// publicKeyIterator is an iterator for a list of public keys.
//
// - implements crypto.PublicKeyIterator
type publicKeyIterator struct {
	*iterator
}

// GetNext implements crypto.PublicKeyIterator. It returns the next public key.
func (i *publicKeyIterator) GetNext() crypto.PublicKey {
	conode := i.iterator.GetNext()
	if conode != nil {
		return conode.GetPublicKey()
	}
	return nil
}

// Conodes is a list of conodes.
//
// - implements mino.Players
// - implements cosi.CollectiveAuthority
// - implements io.WriterTo
// - implements encoding.Packable
type Conodes []Conode

func newConodes(ca cosi.CollectiveAuthority) Conodes {
	conodes := make([]Conode, ca.Len())
	addrIter := ca.AddressIterator()
	publicKeyIter := ca.PublicKeyIterator()
	for i := range conodes {
		conodes[i] = Conode{
			addr:      addrIter.GetNext(),
			publicKey: publicKeyIter.GetNext(),
		}
	}

	return conodes
}

// Rotate takes the new leader and moves it to the beginning of the array while
// moving the old one to the end.
func (cc Conodes) Rotate(addr mino.Address) Conodes {
	index := 0
	for i, conode := range cc {
		if conode.GetAddress().Equal(addr) {
			index = i
		}
	}

	if index == 0 {
		return cc
	}

	newConodes := append(Conodes{cc[index]}, cc[1:index]...)
	newConodes = append(newConodes, cc[index+1:]...)
	newConodes = append(newConodes, cc[0])
	return newConodes
}

// Take implements mino.Players. It returns a subset of the conodes.
func (cc Conodes) Take(filters ...mino.FilterUpdater) mino.Players {
	f := mino.ApplyFilters(filters)
	conodes := make(Conodes, len(f.Indices))
	for i, k := range f.Indices {
		conodes[i] = cc[k]
	}
	return conodes
}

// Len implements mino.Players. It returns the length of the list of conodes.
func (cc Conodes) Len() int {
	return len(cc)
}

// AddressIterator implements mino.Players. It returns the address iterator.
func (cc Conodes) AddressIterator() mino.AddressIterator {
	return &addressIterator{
		iterator: &iterator{
			index:   -1,
			conodes: cc,
		},
	}
}

// PublicKeyIterator implements cosi.CollectiveAuthority. It returns the public
// key iterator.
func (cc Conodes) PublicKeyIterator() crypto.PublicKeyIterator {
	return &publicKeyIterator{
		iterator: &iterator{
			index:   -1,
			conodes: cc,
		},
	}
}

// Pack implements encoding.Packable. It converts the list of conodes to a list
// of protobuf messages.
func (cc Conodes) Pack() (proto.Message, error) {
	pb := &Roster{
		Conodes: make([]*ConodeProto, len(cc)),
	}

	for i, conode := range cc {
		packed, err := conode.Pack()
		if err != nil {
			return nil, encoding.NewEncodingError("conode", err)
		}

		pb.Conodes[i] = packed.(*ConodeProto)
	}

	return pb, nil
}

// WriteTo implements io.WriterTo. It writes the roster in the writer.
func (cc Conodes) WriteTo(w io.Writer) (int64, error) {
	sum := int64(0)

	for _, conode := range cc {
		buffer, err := conode.GetPublicKey().MarshalBinary()
		if err != nil {
			return sum, xerrors.Errorf("couldn't marshal public key: %v", err)
		}

		n, err := w.Write(buffer)
		sum += int64(n)
		if err != nil {
			return sum, xerrors.Errorf("couldn't write public key: %v", err)
		}

		n, err = w.Write([]byte(conode.GetAddress().String()))
		sum += int64(n)
		if err != nil {
			return sum, xerrors.Errorf("couldn't write address: %v", err)
		}
	}

	return sum, nil
}
