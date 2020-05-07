package threshold

import (
	"go.dedis.ch/fabric/cosi"
	"go.dedis.ch/fabric/crypto"
	"go.dedis.ch/fabric/encoding"
	"go.dedis.ch/fabric/mino"
	"golang.org/x/xerrors"
)

//go:generate protoc -I ./ --go_out=./ ./messages.proto

// Threshold is a function that returns the threshold to reach for a given n.
type Threshold func(int) int

func defaultThreshold(n int) int {
	return n
}

// CoSi is an implementation of the cosi.CollectiveSigning interface that is
// using streams to parallelize the work.
type CoSi struct {
	encoder encoding.ProtoMarshaler
	mino    mino.Mino
	signer  crypto.AggregateSigner

	Threshold Threshold
}

// NewCoSi returns a new instance.
func NewCoSi(m mino.Mino, signer crypto.AggregateSigner) *CoSi {
	return &CoSi{
		encoder:   encoding.NewProtoEncoder(),
		mino:      m,
		signer:    signer,
		Threshold: defaultThreshold,
	}
}

// GetPublicKeyFactory implements cosi.CollectiveSigning. It returns the public
// key factory.
func (c *CoSi) GetPublicKeyFactory() crypto.PublicKeyFactory {
	return c.signer.GetPublicKeyFactory()
}

// GetSignatureFactory implements cosi.CollectiveSigning. It returns the
// signature factory.
func (c *CoSi) GetSignatureFactory() crypto.SignatureFactory {
	return signatureFactory{
		encoder:    c.encoder,
		sigFactory: c.signer.GetSignatureFactory(),
	}
}

// GetVerifierFactory implements cosi.CollectiveSigning. It returns the verifier
// factory.
func (c *CoSi) GetVerifierFactory() crypto.VerifierFactory {
	return verifierFactory{factory: c.signer.GetVerifierFactory()}
}

// Listen implements cosi.CollectiveSigning.
func (c *CoSi) Listen(h cosi.Hashable) (cosi.Actor, error) {
	rpc, err := c.mino.MakeRPC("cosi", newHandler(c, h))
	if err != nil {
		return nil, xerrors.Errorf("couldn't make rpc: %v", err)
	}

	return newActor(c, rpc), nil
}
