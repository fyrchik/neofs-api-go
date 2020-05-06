package object

// SignedData returns marshaled payload of the Put request.
//
// If payload is nil, ErrHeaderNotFound returns.
func (m PutRequest) SignedData() ([]byte, error) {
	r := m.GetR()
	if r == nil {
		return nil, ErrHeaderNotFound
	}

	data := make([]byte, r.Size())

	if _, err := r.MarshalTo(data); err != nil {
		return nil, err
	}

	return data, nil
}

// ReadSignedData copies marshaled payload of the Put request to passed buffer.
//
// If payload is nil, ErrHeaderNotFound returns.
func (m PutRequest) ReadSignedData(p []byte) error {
	r := m.GetR()
	if r == nil {
		return ErrHeaderNotFound
	}

	_, err := r.MarshalTo(p)

	return err
}

// SignedDataSize returns the size of payload of the Put request.
//
// If payload is nil, -1 returns.
func (m PutRequest) SignedDataSize() int {
	r := m.GetR()
	if r == nil {
		return -1
	}

	return r.Size()
}
