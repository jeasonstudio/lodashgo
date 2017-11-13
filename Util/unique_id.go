package Util

import (
	"strconv"
)

// UniqueID Generates a unique ID. If prefix is given, the ID is appended to it.
func UniqueID(prefix string) func() string {
	idMap := make(map[string]int)
	if prefix == "" { prefix = "LODASHGO_DEFAULT_" }
	return func () string {
		idMap[prefix] = idMap[prefix] + 1
		return prefix + strconv.Itoa(idMap[prefix])
	}
}