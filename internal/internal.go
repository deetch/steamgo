package internal

import (
	"encoding/binary"
	"io"
)

// EMsg helpers

const (
	protoMask uint32 = 0x80000000
	eMsgMask         = ^protoMask
)

func NewEMsg(e uint32) EMsg {
	return EMsg(e & eMsgMask)
}

func IsProto(e uint32) bool {
	return e&protoMask > 0
}

// Misc

type JobId uint64

type Serializer interface {
	Serialize(w io.Writer) error
}

type Deserializer interface {
	Deserialize(r io.Reader) error
}

type Serializable interface {
	Serializer
	Deserializer
}

type MessageBody interface {
	Serializable
	GetEMsg() EMsg
}

// Helpers for steam_language code

func ReadByte2Bool(r io.Reader) (bool, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c != 0, err
}

func ReadUint8(r io.Reader) (uint8, error) {
	var c uint8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint16(r io.Reader) (uint16, error) {
	var c uint16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint32(r io.Reader) (uint32, error) {
	var c uint32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadUint64(r io.Reader) (uint64, error) {
	var c uint64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt8(r io.Reader) (int8, error) {
	var c int8
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt16(r io.Reader) (int16, error) {
	var c int16
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt32(r io.Reader) (int32, error) {
	var c int32
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func ReadInt64(r io.Reader) (int64, error) {
	var c int64
	err := binary.Read(r, binary.LittleEndian, &c)
	return c, err
}

func WriteBool2Byte(w io.Writer, b bool) error {
	var err error
	if b {
		_, err = w.Write([]byte{1})
	} else {
		_, err = w.Write([]byte{0})
	}
	return err
}
