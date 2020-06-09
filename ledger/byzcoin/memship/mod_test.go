package memship

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/dela/consensus/viewchange"
	"go.dedis.ch/dela/consensus/viewchange/roster"
	"go.dedis.ch/dela/encoding"
	internal "go.dedis.ch/dela/internal/testing"
	"go.dedis.ch/dela/internal/testing/fake"
	types "go.dedis.ch/dela/ledger/byzcoin/memship/json"
	"go.dedis.ch/dela/ledger/inventory"
	"go.dedis.ch/dela/ledger/transactions/basic"
	"go.dedis.ch/dela/serde/json"
	"golang.org/x/xerrors"
)

func TestMessages(t *testing.T) {
	messages := []proto.Message{
		&Task{},
	}

	for _, m := range messages {
		internal.CoverProtoMessage(t, m)
	}
}

func TestClientTask_GetChangeSet(t *testing.T) {
	task := NewRemove([]uint32{1, 3, 1, 3, 4}).(clientTask)

	changeset := task.GetChangeSet()
	require.Equal(t, []uint32{4, 3, 1}, changeset.Remove)

	task = NewAdd(fake.NewAddress(0), fake.PublicKey{}).(clientTask)

	changeset = task.GetChangeSet()
	require.Len(t, changeset.Add, 1)
}

func TestClientTask_Pack(t *testing.T) {
	task := NewRemove([]uint32{1})

	taskpb, err := task.Pack(nil)
	require.NoError(t, err)
	require.Equal(t, []uint32{1}, taskpb.(*Task).GetRemove())

	task = NewAdd(fake.NewAddress(0), fake.PublicKey{})
	taskpb, err = task.Pack(encoding.NewProtoEncoder())
	require.NoError(t, err)
	require.NotNil(t, taskpb.(*Task).GetAddr())
	require.NotNil(t, taskpb.(*Task).GetPublicKey())

	task = NewAdd(fake.NewBadAddress(), nil)
	_, err = task.Pack(nil)
	require.EqualError(t, err, "couldn't marshal address: fake error")

	task = NewAdd(fake.NewAddress(0), fake.PublicKey{})
	_, err = task.Pack(fake.BadPackAnyEncoder{})
	require.EqualError(t, err, "couldn't pack public key: fake error")
}

func TestClientTask_VisitJSON(t *testing.T) {
	task := NewRemove([]uint32{2, 4})
	ser := json.NewSerializer()

	data, err := ser.Serialize(task)
	require.NoError(t, err)
	require.Equal(t, `{"Remove":[2,4],"Address":null,"PublicKey":null}`, string(data))

	task = NewAdd(fake.NewAddress(0), fake.PublicKey{})
	data, err = ser.Serialize(task)
	require.NoError(t, err)
	require.Equal(t, `{"Remove":null,"Address":"AAAAAA==","PublicKey":{}}`, string(data))

	task = clientTask{player: &viewchange.Player{Address: fake.NewBadAddress()}}
	_, err = task.VisitJSON(ser)
	require.EqualError(t, err, "couldn't marshal address: fake error")

	task = clientTask{player: &viewchange.Player{Address: fake.NewAddress(0)}}
	_, err = task.VisitJSON(fake.NewBadSerializer())
	require.EqualError(t, err, "couldn't serialize public key: fake error")
}

func TestClientTask_Fingerprint(t *testing.T) {
	// Client task with indices to remove.
	task := NewRemove([]uint32{0x02, 0x01, 0x03})

	buffer := new(bytes.Buffer)

	err := task.Fingerprint(buffer, nil)
	require.NoError(t, err)
	require.Equal(t, "\x02\x00\x00\x00\x01\x00\x00\x00\x03\x00\x00\x00", buffer.String())

	err = task.Fingerprint(fake.NewBadHash(), nil)
	require.EqualError(t, err, "couldn't write remove indices: fake error")

	// Client task with a player to add.
	task = NewAdd(fake.NewAddress(5), fake.PublicKey{})

	buffer = new(bytes.Buffer)

	err = task.Fingerprint(buffer, nil)
	require.NoError(t, err)
	require.Equal(t, "\x05\x00\x00\x00\xdf", buffer.String())

	err = task.Fingerprint(fake.NewBadHashWithDelay(1), nil)
	require.EqualError(t, err, "couldn't write address: fake error")

	err = task.Fingerprint(fake.NewBadHashWithDelay(2), nil)
	require.EqualError(t, err, "couldn't write public key: fake error")

	task = NewAdd(fake.NewBadAddress(), fake.PublicKey{})
	err = task.Fingerprint(buffer, nil)
	require.EqualError(t, err, "couldn't marshal address: fake error")

	task = NewAdd(fake.NewAddress(0), fake.NewBadPublicKey())
	err = task.Fingerprint(buffer, nil)
	require.EqualError(t, err, "couldn't marshal public key: fake error")
}

func TestServerTask_Consume(t *testing.T) {
	r := roster.New(fake.NewAuthority(3, fake.NewSigner))
	rosterpb, err := r.Pack(encoding.NewProtoEncoder())
	require.NoError(t, err)

	task := serverTask{
		clientTask:    clientTask{remove: []uint32{2}},
		rosterFactory: roster.NewRosterFactory(fake.AddressFactory{}, fake.PublicKeyFactory{}),
		encoder:       encoding.NewProtoEncoder(),
		inventory:     fakeInventory{value: rosterpb},
	}

	err = task.Consume(nil, fakePage{values: map[string]proto.Message{}})
	require.NoError(t, err)

	task.inventory = fakeInventory{err: xerrors.New("oops")}
	err = task.Consume(nil, fakePage{})
	require.EqualError(t, err, "couldn't get previous page: oops")

	task.inventory = fakeInventory{errPage: xerrors.New("oops")}
	err = task.Consume(nil, fakePage{})
	require.EqualError(t, err, "couldn't read roster: oops")

	task.inventory = fakeInventory{value: nil}
	err = task.Consume(nil, fakePage{})
	require.EqualError(t, err, "couldn't decode roster: invalid message type '<nil>'")

	task.inventory = fakeInventory{value: rosterpb}
	task.encoder = fake.BadPackEncoder{}
	err = task.Consume(nil, fakePage{})
	require.EqualError(t, err, "couldn't encode roster: fake error")

	task.encoder = encoding.NewProtoEncoder()
	err = task.Consume(nil, fakePage{errWrite: xerrors.New("oops")})
	require.EqualError(t, err, "couldn't write roster: oops")
}

func TestTaskManager_GetAuthority(t *testing.T) {
	manager := TaskManager{
		inventory:     fakeInventory{},
		rosterFactory: fakeRosterFactory{},
	}

	authority, err := manager.GetAuthority(0)
	require.NoError(t, err)
	require.Equal(t, 3, authority.Len())

	manager.inventory = fakeInventory{err: xerrors.New("oops")}
	_, err = manager.GetAuthority(1)
	require.EqualError(t, err, "couldn't read page: oops")

	manager.inventory = fakeInventory{errPage: xerrors.New("oops")}
	_, err = manager.GetAuthority(1)
	require.EqualError(t, err, "couldn't read entry: oops")

	manager.inventory = fakeInventory{}
	manager.rosterFactory = fakeRosterFactory{err: xerrors.New("oops")}
	_, err = manager.GetAuthority(1)
	require.EqualError(t, err, "couldn't decode roster: oops")
}

func TestTaskManager_Wait(t *testing.T) {
	manager := TaskManager{
		me:            fake.NewAddress(0),
		inventory:     fakeInventory{},
		rosterFactory: fakeRosterFactory{},
	}

	allowed := manager.Wait()
	require.True(t, allowed)

	manager.me = fake.NewAddress(1)
	allowed = manager.Wait()
	require.False(t, allowed)

	manager.inventory = fakeInventory{err: xerrors.New("oops")}
	allowed = manager.Wait()
	require.False(t, allowed)
}

func TestTaskManager_Verify(t *testing.T) {
	manager := TaskManager{
		inventory:     fakeInventory{},
		rosterFactory: fakeRosterFactory{},
	}

	authority, err := manager.Verify(fake.NewAddress(0), 0)
	require.NoError(t, err)
	require.Equal(t, 3, authority.Len())

	_, err = manager.Verify(fake.NewAddress(1), 0)
	require.EqualError(t, err, "<fake.Address[1]> is not the leader")

	manager.inventory = fakeInventory{err: xerrors.New("oops")}
	_, err = manager.Verify(fake.NewAddress(0), 0)
	require.EqualError(t, err, "couldn't get authority: couldn't read page: oops")
}

func TestTaskManager_FromProto(t *testing.T) {
	manager := NewTaskManager(fakeInventory{value: &empty.Empty{}}, fake.Mino{}, fake.Signer{})

	task, err := manager.FromProto(&Task{Addr: []byte{0x1}})
	require.NoError(t, err)
	require.NotNil(t, task)

	task, err = manager.FromProto(&Task{Addr: []byte{}, PublicKey: &any.Any{}})
	require.NoError(t, err)
	require.NotNil(t, task.(serverTask).player)

	taskAny, err := ptypes.MarshalAny(&Task{})
	require.NoError(t, err)

	task, err = manager.FromProto(taskAny)
	require.NoError(t, err)
	require.NotNil(t, task)

	_, err = manager.FromProto(&empty.Empty{})
	require.EqualError(t, err, "invalid message type '*empty.Empty'")

	manager.encoder = fake.BadUnmarshalAnyEncoder{}
	_, err = manager.FromProto(taskAny)
	require.EqualError(t, err, "couldn't unmarshal message: fake error")

	manager.rosterFactory = roster.NewRosterFactory(fake.AddressFactory{}, fake.NewBadPublicKeyFactory())
	_, err = manager.FromProto(&Task{Addr: []byte{}, PublicKey: &any.Any{}})
	require.EqualError(t, err,
		"couldn't unpack player: couldn't decode public key: fake error")
}

func TestTaskManager_VisitJSON(t *testing.T) {
	manager := NewTaskManager(nil, fake.Mino{}, fake.NewSigner())
	ser := json.NewSerializer()

	var task serverTask
	err := ser.Deserialize([]byte(`{"Address":"AAAAAA==","PublicKey":{}}`), manager, &task)
	require.NoError(t, err)
	require.NotNil(t, task.player)

	_, err = manager.VisitJSON(fake.NewBadFactoryInput())
	require.EqualError(t, err, "couldn't deserialize task: fake error")

	input := fake.FactoryInput{
		Message: types.Task{Address: []byte{}, PublicKey: []byte{}},
		Serde:   fake.NewBadSerializer(),
	}
	_, err = manager.VisitJSON(input)
	require.EqualError(t, err, "couldn't deserialize public key: fake error")
}

func TestRegister(t *testing.T) {
	factory := basic.NewTransactionFactory(fake.NewSigner())
	Register(factory, NewTaskManager(nil, fake.Mino{}, fake.NewSigner()))
}

// -----------------------------------------------------------------------------
// Utility functions

type fakePage struct {
	inventory.WritablePage
	values   map[string]proto.Message
	errRead  error
	errWrite error
	counter  *fake.Counter
}

func (p fakePage) GetIndex() uint64 {
	return 5
}

func (p fakePage) Read(key []byte) (proto.Message, error) {
	if p.errRead != nil {
		defer p.counter.Decrease()
		if p.counter.Done() {
			return nil, p.errRead
		}
	}

	return p.values[string(key)], nil
}

func (p fakePage) Write(key []byte, value proto.Message) error {
	if p.errWrite != nil {
		defer p.counter.Decrease()
		if p.counter.Done() {
			return p.errWrite
		}
	}

	p.values[string(key)] = value
	return nil
}

func (p fakePage) Defer(fn func([]byte)) {
	fn([]byte{0x12})
}

type fakeInventory struct {
	inventory.Inventory
	value   proto.Message
	err     error
	errPage error
}

func (i fakeInventory) Len() uint64 {
	return 1
}

func (i fakeInventory) GetPage(uint64) (inventory.Page, error) {
	values := map[string]proto.Message{
		RosterValueKey: i.value,
	}
	return fakePage{values: values, errRead: i.errPage}, i.err
}

type fakeRosterFactory struct {
	roster.Factory

	err error
}

func (f fakeRosterFactory) FromProto(proto.Message) (viewchange.Authority, error) {
	return roster.New(fake.NewAuthority(3, fake.NewSigner)), f.err
}
