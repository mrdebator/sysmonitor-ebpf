// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64be || armbe || mips || mips64 || mips64p32 || ppc64 || s390 || s390x || sparc || sparc64
// +build arm64be armbe mips mips64 mips64p32 ppc64 s390 s390x sparc sparc64

package syscalls_old

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadSyscalls_old returns the embedded CollectionSpec for syscalls_old.
func loadSyscalls_old() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Syscalls_oldBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load syscalls_old: %w", err)
	}

	return spec, err
}

// loadSyscalls_oldObjects loads syscalls_old and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*syscalls_oldObjects
//	*syscalls_oldPrograms
//	*syscalls_oldMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadSyscalls_oldObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadSyscalls_old()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// syscalls_oldSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type syscalls_oldSpecs struct {
	syscalls_oldProgramSpecs
	syscalls_oldMapSpecs
}

// syscalls_oldSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type syscalls_oldProgramSpecs struct {
	SysExit *ebpf.ProgramSpec `ebpf:"sys_exit"`
}

// syscalls_oldMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type syscalls_oldMapSpecs struct {
	NamespaceTable *ebpf.MapSpec `ebpf:"namespace_table"`
	SyscallTable   *ebpf.MapSpec `ebpf:"syscall_table"`
}

// syscalls_oldObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadSyscalls_oldObjects or ebpf.CollectionSpec.LoadAndAssign.
type syscalls_oldObjects struct {
	syscalls_oldPrograms
	syscalls_oldMaps
}

func (o *syscalls_oldObjects) Close() error {
	return _Syscalls_oldClose(
		&o.syscalls_oldPrograms,
		&o.syscalls_oldMaps,
	)
}

// syscalls_oldMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadSyscalls_oldObjects or ebpf.CollectionSpec.LoadAndAssign.
type syscalls_oldMaps struct {
	NamespaceTable *ebpf.Map `ebpf:"namespace_table"`
	SyscallTable   *ebpf.Map `ebpf:"syscall_table"`
}

func (m *syscalls_oldMaps) Close() error {
	return _Syscalls_oldClose(
		m.NamespaceTable,
		m.SyscallTable,
	)
}

// syscalls_oldPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadSyscalls_oldObjects or ebpf.CollectionSpec.LoadAndAssign.
type syscalls_oldPrograms struct {
	SysExit *ebpf.Program `ebpf:"sys_exit"`
}

func (p *syscalls_oldPrograms) Close() error {
	return _Syscalls_oldClose(
		p.SysExit,
	)
}

func _Syscalls_oldClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed syscalls_old_bpfeb.o
var _Syscalls_oldBytes []byte
