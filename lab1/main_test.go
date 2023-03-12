package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func TestConvertIntToBigEndian(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Uint32()
		var bigEndian []byte
		bigEndian = convertIntToBigEndian(n)
		expected := convertIntToBigEndianLiblary(n)

		if bytes.Compare(bigEndian, expected) == 0 {
			t.Logf("\"convertIntToBigEndian\" SUCCESS, expected -> %v, got -> %v", expected, bigEndian)

		} else {
			t.Errorf("\"convertIntToBigEndian\" FAILED, expected -> %v, got -> %v", expected, bigEndian)
		}
	}
}

func TestConvertBigEndianToInt(t *testing.T) {
	for i := 0; i < 5; i++ {
		bytesArr := make([]byte, 4)
		rand.Read(bytesArr)
		var bigEndianInt uint32
		bigEndianInt = convertBigEndianToInt(bytesArr)
		expected := convertBigEndianToIntLiblary(bytesArr)

		if bigEndianInt == expected {
			t.Logf("\"convertIntToBigEndian\" SUCCESS, expected -> %v, got -> %v", expected, bigEndianInt)

		} else {
			t.Errorf("\"convertIntToBigEndian\" FAILED, expected -> %v, got -> %v", expected, bigEndianInt)
		}
	}
}

func TestConvertLittleEndianToInt(t *testing.T) {
	for i := 0; i < 5; i++ {
		n := rand.Uint32()
		var littleEndian []byte
		littleEndian = convertIntToLittleEndian(n)
		expected := convertIntToLittleEndianLiblary(n)

		if bytes.Compare(littleEndian, expected) == 0 {
			t.Logf("\"convertIntToLittleEndian\" SUCCESS, expected -> %v, got -> %v", expected, littleEndian)

		} else {
			t.Errorf("\"convertIntToBigEndian\" FAILED, expected -> %v, got -> %v", expected, littleEndian)
		}
	}
}

func TestConvertIntToLittleEndian(t *testing.T) {
	for i := 0; i < 5; i++ {
		bytesArr := make([]byte, 4)
		rand.Read(bytesArr)
		var littleEndianInt uint32
		littleEndianInt = convertLittleEndianToInt(bytesArr)
		expected := convertLittleEndianToIntLiblary(bytesArr)

		if littleEndianInt == expected {
			t.Logf("\"convertLittleEndianToInt\" SUCCESS, expected -> %v, got -> %v", expected, littleEndianInt)

		} else {
			t.Errorf("\"convertLittleEndianToInt\" FAILED, expected -> %v, got -> %v", expected, littleEndianInt)
		}
	}
}
