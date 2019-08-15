// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbsemantic

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Prop struct {
	_tab flatbuffers.Table
}

func GetRootAsProp(buf []byte, offset flatbuffers.UOffsetT) *Prop {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Prop{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Prop) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Prop) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Prop) K() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Prop) VType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Prop) MutateVType(n byte) bool {
	return rcv._tab.MutateByteSlot(6, n)
}

func (rcv *Prop) V(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func PropStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PropAddK(builder *flatbuffers.Builder, k flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(k), 0)
}
func PropAddVType(builder *flatbuffers.Builder, vType byte) {
	builder.PrependByteSlot(1, vType, 0)
}
func PropAddV(builder *flatbuffers.Builder, v flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(v), 0)
}
func PropEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}