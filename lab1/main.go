package main

import (
	"bytes"
	_ "bytes"
	"encoding/binary"
	"fmt"
)

func convertIntToBigEndian(n uint32) []byte {
	bigEndian := []byte{
		byte(n >> 24 & 0xFF),
		byte(n >> 16 & 0xFF),
		byte(n >> 8 & 0xFF),
		byte(n & 0xFF),
	}
	return bigEndian
}

func convertBigEndianToInt(n []byte) uint32 {
	var bigEndianInt uint32
	for _, b := range n {
		bigEndianInt = bigEndianInt<<8 | uint32(b)
	}
	return bigEndianInt
}

func convertIntToBigEndianLiblary(n uint32) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, uint32(n))
	bigEndianResult := buf.Bytes()

	return bigEndianResult
}
func convertBigEndianToIntLiblary(n []byte) uint32 {
	var bigEndianInt uint32
	_ = binary.Read(bytes.NewReader(n), binary.BigEndian, &bigEndianInt)
	return bigEndianInt
}
func convertIntToLittleEndian(n uint32) []byte {
	littleEndian := []byte{
		byte(n & 0xFF),
		byte(n >> 8 & 0xFF),
		byte(n >> 16 & 0xFF),
		byte(n >> 24 & 0xFF),
	}
	return littleEndian
}
func convertLittleEndianToInt(n []byte) uint32 {
	var littleEndianInt uint32
	for i := len(n) - 1; i >= 0; i-- {
		littleEndianInt <<= 8
		littleEndianInt |= uint32(n[i])

	}
	return littleEndianInt
}
func convertIntToLittleEndianLiblary(n uint32) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.LittleEndian, uint32(n))
	littleEndianResult := buf.Bytes()

	return littleEndianResult
}
func convertLittleEndianToIntLiblary(n []byte) uint32 {
	var littleEndianInt uint32
	_ = binary.Read(bytes.NewReader(n), binary.LittleEndian, &littleEndianInt)
	return littleEndianInt
}
func main() {
	n := uint32(0x12345678)
	var bigEndian, bigEndianLib []byte
	var bigEndianToInt, bigEndianToIntLib uint32

	bigEndian = convertIntToBigEndian(n)
	fmt.Printf("Big endian %X\n", bigEndian)

	bigEndianToInt = convertBigEndianToInt(bigEndian)
	fmt.Printf("Big endian to int %d\n", bigEndianToInt)

	bigEndianLib = convertIntToBigEndianLiblary(n)
	fmt.Printf("Big endian lib %X\n", bigEndianLib)

	bigEndianToIntLib = convertBigEndianToIntLiblary(bigEndianLib)
	fmt.Printf("Big endian to int lib %d\n\n", bigEndianToIntLib)

	var littleEndian, littleEndianLib []byte
	var littleEndianToInt, littleEndianToIntLib uint32

	littleEndian = convertIntToLittleEndian(n)
	fmt.Printf("Little endian %X\n", littleEndian)

	littleEndianToInt = convertLittleEndianToInt(littleEndian)
	fmt.Printf("Little endian to int %d\n", littleEndianToInt)

	littleEndianLib = convertIntToLittleEndianLiblary(n)
	fmt.Printf("Little endian lib %X\n", littleEndianLib)

	littleEndianToIntLib = convertLittleEndianToIntLiblary(littleEndianLib)
	fmt.Printf("Little endian to int lib %d\n", littleEndianToIntLib)
}
