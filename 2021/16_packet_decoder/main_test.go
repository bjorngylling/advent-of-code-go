package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_readLiteralPacket(t *testing.T) {
	hexdec, _ := hex.DecodeString("D2FE28")
	r := NewBitReader(bytes.NewReader(hexdec))
	got := readHeader(r)
	if got.ver != 6 {
		t.Errorf("header ver = %d, want 6", got.ver)
	}
	if got.t != 4 {
		t.Errorf("header t = %d, want 4", got.t)
	}
	content := readLiteralContent(r)
	if content != 2021 {
		t.Errorf("content = %d, want 2021", content)
	}
}

func Test_readOperatorPacket_bitLengthMode(t *testing.T) {
	hexdec, _ := hex.DecodeString("38006F45291200")
	r := NewBitReader(bytes.NewReader(hexdec))
	got := readHeader(r)
	if got.ver != 1 {
		t.Errorf("header ver = %d, want 1", got.ver)
	}
	if got.t != 6 {
		t.Errorf("header t = %d, want 6", got.t)
	}
	sumVer, _ := readOperatorContent(r)
	sumVer += got.ver
	if sumVer != 9 {
		t.Errorf("sum of versions = %d, want 9", sumVer)
	}
}

func Test_readOperatorPacket_subpktCountMode(t *testing.T) {
	hexdec, _ := hex.DecodeString("EE00D40C823060")
	r := NewBitReader(bytes.NewReader(hexdec))
	got := readHeader(r)
	if got.ver != 7 {
		t.Errorf("header ver = %d, want 7", got.ver)
	}
	if got.t != 3 {
		t.Errorf("header t = %d, want 3", got.t)
	}
	sumVer, _ := readOperatorContent(r)
	sumVer += got.ver
	if sumVer != 14 {
		t.Errorf("sum of versions = %d, want 14", sumVer)
	}
}

func Test_readPacket(t *testing.T) {
	hexdec, _ := hex.DecodeString("D2FE28")
	r := NewBitReader(bytes.NewReader(hexdec))
	got, _ := readPacket(r)
	if got != 6 {
		t.Errorf("readPacket = %d, want 6", got)
	}
}

func Test_solve(t *testing.T) {
	for i, tt := range testPuzzles {
		t.Run(fmt.Sprint("example ", i), func(t *testing.T) {
			part1, part2 := solve(tt.Puzzle)
			if tt.Solution1 != "" && part1 != tt.Solution1 {
				t.Errorf("solve() part1 = %v, want %v", part1, tt.Solution1)
			}
			if tt.Solution2 != "" && part2 != tt.Solution2 {
				t.Errorf("solve() part2 = %v, want %v", part2, tt.Solution2)
			}
		})
	}
}
