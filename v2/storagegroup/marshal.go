package storagegroup

import (
	"github.com/nspcc-dev/neofs-api-go/util/proto"
)

const (
	sizeField       = 1
	hashField       = 2
	expirationField = 3
	objectIDsField  = 4
)

// StableMarshal marshals unified storage group structure in a protobuf
// compatible way without field order shuffle.
func (s *StorageGroup) StableMarshal(buf []byte) ([]byte, error) {
	if s == nil {
		return []byte{}, nil
	}

	if buf == nil {
		buf = make([]byte, s.StableSize())
	}

	var (
		offset, n int
		err       error
	)

	n, err = proto.UInt64Marshal(sizeField, buf[offset:], s.size)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.BytesMarshal(hashField, buf[offset:], s.hash)
	if err != nil {
		return nil, err
	}

	offset += n

	n, err = proto.UInt64Marshal(expirationField, buf[offset:], s.exp)
	if err != nil {
		return nil, err
	}

	offset += n

	for i := range s.members {
		n, err = proto.NestedStructureMarshal(objectIDsField, buf[offset:], s.members[i])
		if err != nil {
			return nil, err
		}

		offset += n
	}

	return buf, nil
}

// StableSize of storage group structure marshalled by StableMarshal function.
func (s *StorageGroup) StableSize() (size int) {
	if s == nil {
		return 0
	}

	size += proto.UInt64Size(sizeField, s.size)
	size += proto.BytesSize(hashField, s.hash)
	size += proto.UInt64Size(expirationField, s.exp)

	for i := range s.members {
		size += proto.NestedStructureSize(objectIDsField, s.members[i])
	}

	return size
}
