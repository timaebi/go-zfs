package zfsiface

import (
	"io"
)

type Dataset interface {
	//GetName() string
	//GetOrigin() string
	//GetUsed() uint64
	//GetAvail() uint64
	//GetMountPoint()string
	//GetCompression() string
	//GetType() string
	//GetWritten() uint64
	//GetVolsize() uint64
	//GetLogicalused() uint64
	//GetUsedbydataset() uint64
	//GetQuota() uint64
	//GetReferenced() uint64

	Clone(dest string, properties map[string]string) (Dataset, error)
	Unmount(force bool) (Dataset, error)
	Mount(overlay bool, options []string) (Dataset, error)
	SendSnapshot(output io.Writer) error
	Destroy(flags DestroyFlag) error
	SetProperty(key, val string) error
	GetProperty(key string) (string, error)
	Rename(name string, createParent bool, recursiveRenameSnapshots bool) (Dataset, error)
	Snapshots() ([]Dataset, error)
	Snapshot(name string, recursive bool) (Dataset, error)
	Rollback(destroyMoreRecent bool) error
	Children(depth uint64) ([]Dataset, error)
	Diff(snapshot string) ([]*InodeChange, error)
}

// InodeType is the type of inode as reported by Diff
type InodeType int

// Types of Inodes
const (
	_                         = iota // 0 == unknown type
	BlockDevice     InodeType = iota
	CharacterDevice
	Directory
	Door
	NamedPipe
	SymbolicLink
	EventPort
	Socket
	File
)

// ChangeType is the type of inode change as reported by Diff
type ChangeType int

// Types of Changes
const (
	_                   = iota // 0 == unknown type
	Removed  ChangeType = iota
	Created
	Modified
	Renamed
)

// DestroyFlag is the options flag passed to Destroy
type DestroyFlag int

// Valid destroy options
const (
	DestroyDefault         DestroyFlag = 1 << iota
	DestroyRecursive                   = 1 << iota
	DestroyRecursiveClones             = 1 << iota
	DestroyDeferDeletion               = 1 << iota
	DestroyForceUmount                 = 1 << iota
)

// InodeChange represents a change as reported by Diff
type InodeChange struct {
	Change               ChangeType
	Type                 InodeType
	Path                 string
	NewPath              string
	ReferenceCountChange int
}
