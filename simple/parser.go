package simple

import (
	"fmt"
	"io"
)

func extractField(bz []byte, field int32) ([]byte, error) {
	var err error
	var fieldNum int32
	for len(bz) > 0 {
		// parse the header fro field type
		bz, fieldNum, _, err = parseFieldHeader(bz)
		if err != nil {
			return nil, err
		}

		// we got it!
		if fieldNum == field {
			return bz, nil
		}

		// skip field
		skippy, err := skipSample(bz)
		if err != nil {
			return nil, err
		}
		if skippy < 0 {
			return nil, ErrInvalidLengthSample
		}
		if (skippy) > len(bz) {
			return nil, io.ErrUnexpectedEOF
		}
		bz = bz[skippy:]
	}
	return nil, fmt.Errorf("Desired field %d not found", field)
}

func parseFieldHeader(bz []byte) (rest []byte, fieldNum int32, wireType int, err error) {
	var wire uint64
	var offset int
	wire, offset, err = parseUint64(bz)
	if err != nil {
		return
	}
	wireType = int(wire & 0x7)
	fieldNum = int32(wire >> 3)
	if fieldNum <= 0 {
		err = fmt.Errorf("proto: Person: illegal tag %d (wire type %d)", fieldNum, wireType)
		return
	}
	rest = bz[offset:]
	return
}

func parseInt64(bz []byte) (wire int64, offset int, err error) {
	var uwire uint64
	uwire, offset, err = parseVarUint(bz, 64)
	wire = int64(uwire)
	return
}

func parseInt32(bz []byte) (wire int32, offset int, err error) {
	var uwire uint64
	uwire, offset, err = parseVarUint(bz, 32)
	wire = int32(uwire)
	return
}

func parseInt(bz []byte) (wire int, offset int, err error) {
	var uwire uint64
	uwire, offset, err = parseVarUint(bz, 64)
	wire = int(uwire)
	return
}

func parseUint64(bz []byte) (wire uint64, offset int, err error) {
	return parseVarUint(bz, 64)
}

func parseUint32(bz []byte) (wire uint32, offset int, err error) {
	var uwire uint64
	uwire, offset, err = parseVarUint(bz, 32)
	wire = uint32(uwire)
	return
}

func parseString(bz []byte) (string, error) {
	size, offset, err := parseInt(bz)
	if err != nil {
		return "", err
	}
	return string(bz[offset : offset+size]), nil
}

// parseVarUint is a helper and returns bytes as uint64
// to be converted by wrapper
func parseVarUint(bz []byte, maxShift uint) (wire uint64, offset int, err error) {
	l := len(bz)
	for shift := uint(0); ; shift += 7 {
		if shift >= maxShift {
			err = ErrIntOverflowSample
			return
		}
		if offset >= l {
			err = io.ErrUnexpectedEOF
			return
		}
		b := bz[offset]
		offset++
		wire |= (uint64(b) & 0x7F) << shift
		if b < 0x80 {
			break
		}
	}
	return wire, offset, nil
}
