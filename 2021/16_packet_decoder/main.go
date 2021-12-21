package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"
)

type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
	count  int
}

func NewBitReader(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0, 0}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	r.count++
	return bit, nil
}

func (r *BitReader) ReadUint(nbits int) (uint64, error) {
	var result uint64
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit {
			result |= 1 << uint(i)
		}
	}
	return result, nil
}

type header struct {
	ver uint64
	t   uint64
}

func readHeader(r *BitReader) header {
	hdr := header{}
	var err error
	hdr.ver, err = r.ReadUint(3)
	if err != nil {
		panic(err)
	}
	hdr.t, err = r.ReadUint(3)
	if err != nil {
		panic(err)
	}
	return hdr
}

func readLiteralContent(r *BitReader) uint64 {
	var result uint64
	cont := true
	i := 1
	for ; cont; i += 5 {
		var err error
		cont, err = r.ReadBit()
		if err != nil {
			panic(err)
		}
		for j := 0; j < 4; j++ {
			bit, err := r.ReadBit()
			if err != nil {
				panic(err)
			}
			result <<= 1
			if bit {
				result++
			}
		}
	}
	return result
}

func readOperatorContent(r *BitReader) (uint64, []uint64) {
	subpktCountMode, err := r.ReadBit()
	if err != nil {
		panic(err)
	}
	var sumVer uint64
	var lit []uint64
	if subpktCountMode { // total length in count
		count, err := r.ReadUint(11)
		if err != nil {
			panic(err)
		}
		for i := 0; i < int(count); i++ {
			ver, val := readPacket(r)
			sumVer += ver
			lit = append(lit, val)
		}
	} else { // total length in bits
		bitLength, err := r.ReadUint(15)
		if err != nil {
			panic(err)
		}
		start := r.count
		for r.count+1-start < int(bitLength) {
			ver, val := readPacket(r)
			sumVer += ver
			lit = append(lit, val)
		}
	}
	return sumVer, lit
}

func readPacket(r *BitReader) (uint64, uint64) {
	hdr := readHeader(r)
	sumVer := hdr.ver
	var val uint64
	switch hdr.t {
	case 4:
		val = readLiteralContent(r)
	case 0:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		for _, v := range lit {
			val += v
		}
	case 1:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		val = 1
		for _, v := range lit {
			val *= v
		}
	case 2:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		min := uint64(math.MaxUint64)
		for _, v := range lit {
			if v < min {
				min = v
			}
		}
		val = min
	case 3:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		max := uint64(0)
		for _, v := range lit {
			if v > max {
				max = v
			}
		}
		val = max
	case 5:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		if lit[0] > lit[1] {
			val = 1
		}
	case 6:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		if lit[0] < lit[1] {
			val = 1
		}
	case 7:
		ver, lit := readOperatorContent(r)
		sumVer += ver
		if lit[0] == lit[1] {
			val = 1
		}
	}
	return sumVer, val
}

func solve(input string) (string, string) {
	hexdec, err := hex.DecodeString(input)
	r := NewBitReader(bytes.NewReader(hexdec))
	if err != nil {
		panic(err)
	}
	sumVer, val := readPacket(r)
	return strconv.Itoa(int(sumVer)), strconv.Itoa(int(val))
}

func main() {
	start := time.Now()
	part1, part2 := solve(puzzle)
	elapsed := time.Since(start)

	fmt.Printf("Part 1: %s\nPart 2: %s\n", part1, part2)
	fmt.Printf("Program took %s", elapsed)
}
