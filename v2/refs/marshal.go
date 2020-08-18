package refs

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	ownerIDValField = 1

	containerIDValField = 1

	objectIDValField = 1

	addressContainerField = 1
	addressObjectField    = 2
)

func (o *OwnerID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(ownerIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *OwnerID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(ownerIDValField, o.val)
}

func (c *ContainerID) StableMarshal(buf []byte) ([]byte, error) {
	if c == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, c.StableSize())
	}

	_, err := proto.BytesMarshal(containerIDValField, buf, c.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *ContainerID) StableSize() int {
	if c == nil {
		return 0
	}

	return proto.BytesSize(containerIDValField, c.val)
}

func (o *ObjectID) StableMarshal(buf []byte) ([]byte, error) {
	if o == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, o.StableSize())
	}

	_, err := proto.BytesMarshal(objectIDValField, buf, o.val)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (o *ObjectID) StableSize() int {
	if o == nil {
		return 0
	}

	return proto.BytesSize(objectIDValField, o.val)
}

func (a *Address) StableMarshal(buf []byte) ([]byte, error) {
	if a == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, a.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.NestedStructureMarshal(addressContainerField, buf[offset:], a.cid)
	if err != nil {
		return nil, err
	}

	offset += n

	_, err = proto.NestedStructureMarshal(addressObjectField, buf[offset:], a.oid)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (a *Address) StableSize() (size int) {
	if a == nil {
		return 0
	}

	size += proto.NestedStructureSize(addressContainerField, a.cid)

	size += proto.NestedStructureSize(addressObjectField, a.oid)

	return size
}
