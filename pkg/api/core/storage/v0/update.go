package v0

import "github.com/vmmgr/imacon/pkg/api/core/storage"

func update(input, replace *storage.Storage) {
	if replace.Type != input.Type {
		replace.Type = input.Type
	}
	if replace.GroupID != input.GroupID {
		replace.GroupID = input.GroupID
	}
	if replace.Path != input.Path {
		replace.Path = input.Path
	}

	if replace.CloudInit != input.CloudInit {
		replace.CloudInit = input.CloudInit
	}
	if replace.MinCPU != input.MinCPU {
		replace.MinCPU = input.MinCPU
	}
	if replace.MinMem != input.MinMem {
		replace.MinMem = input.MinMem
	}
	if replace.OS != input.OS {
		replace.OS = input.OS
	}
	if replace.Admin != input.Admin {
		replace.Admin = input.Admin
	}
	if replace.Lock != input.Lock {
		replace.Lock = input.Lock
	}
}
