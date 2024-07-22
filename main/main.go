package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

type Cases struct {
	Name      string
	Payload   []byte
	Result    string
	ResultRaw []byte
}

var dict []string

func init() {
	dict = []string{"SELECT", "ALTER", "select", "alter"}
}

func parser(payload []byte) ([]byte, error) {
	// param check
	if len(payload) < 8 {
		return nil, errors.New("payload is too short. ")
	}
	length := int(binary.BigEndian.Uint16(payload[0:2]))
	if length != len(payload) {
		return nil, errors.New("length check failed. ")
	}
	if payload[4] == 0 {
		return nil, errors.New("sql flag should not be 0. ")
	}

	// find start index for sql
	start := 8
	for start < len(payload) {
		fmt.Println(start)
		if !(32 <= payload[start] && payload[start] <= 127) {
			start++
		} else {
			break
		}
	}
	found := false
	for start < len(payload) {
		for _, key := range dict {
			if strings.HasPrefix(string(payload[start:]), key) {
				found = true
				break
			}
		}
		if found {
			break
		} else {
			start++
		}
	}

	// not found
	if !found {
		return nil, errors.New("no sql found. ")
	}

	// find end index
	end := start + 1
	for end < len(payload) {
		if 32 <= payload[end] && payload[end] <= 127 {
			end++
		} else {
			break
		}
	}
	return payload[start:end], nil
}
