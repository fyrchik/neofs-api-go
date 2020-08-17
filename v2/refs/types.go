package refs

type OwnerID struct {
	val []byte
}

type ContainerID struct {
	val []byte
}

type ObjectID struct {
	val []byte
}

type Address struct {
	cid *ContainerID

	oid *ObjectID
}

func (o *OwnerID) GetValue() []byte {
	if o != nil {
		return o.val
	}

	return nil
}

func (o *OwnerID) SetValue(v []byte) {
	if o != nil {
		o.val = v
	}
}

func (c *ContainerID) GetValue() []byte {
	if c != nil {
		return c.val
	}

	return nil
}

func (c *ContainerID) SetValue(v []byte) {
	if c != nil {
		c.val = v
	}
}

func (o *ObjectID) GetValue() []byte {
	if o != nil {
		return o.val
	}

	return nil
}

func (o *ObjectID) SetValue(v []byte) {
	if o != nil {
		o.val = v
	}
}

func (a *Address) GetContainerID() *ContainerID {
	if a != nil {
		return a.cid
	}

	return nil
}

func (a *Address) SetContainerID(v *ContainerID) {
	if a != nil {
		a.cid = v
	}
}

func (a *Address) GetObjectID() *ObjectID {
	if a != nil {
		return a.oid
	}

	return nil
}

func (a *Address) SetObjectID(v *ObjectID) {
	if a != nil {
		a.oid = v
	}
}
