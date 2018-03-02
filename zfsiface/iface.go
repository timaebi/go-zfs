package zfsiface

import (
	"io"
	"time"
)

type Dataset interface {
	GetNativeProperties() *NativeProperties
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

type NativeProperties struct {
	Name          string
	Origin        string
	Used          uint64
	Avail         uint64
	Mountpoint    string
	Compression   string
	Type          string
	Written       uint64
	Volsize       uint64
	Logicalused   uint64
	Usedbydataset uint64
	Quota         uint64
	Referenced    uint64
	Creation       time.Time
}
