// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build windows

// Code generated by 'go generate'; DO NOT EDIT.

package dism

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modDismAPI = windows.NewLazySystemDLL("DismAPI.dll")

	procDismAddCapability    = modDismAPI.NewProc("DismAddCapability")
	procDismAddDriver        = modDismAPI.NewProc("DismAddDriver")
	procDismAddPackage       = modDismAPI.NewProc("DismAddPackage")
	procDismApplyUnattend    = modDismAPI.NewProc("DismApplyUnattend")
	procDismCloseSession     = modDismAPI.NewProc("DismCloseSession")
	procDismDisableFeature   = modDismAPI.NewProc("DismDisableFeature")
	procDismEnableFeature    = modDismAPI.NewProc("DismEnableFeature")
	procDismInitialize       = modDismAPI.NewProc("DismInitialize")
	procDismOpenSession      = modDismAPI.NewProc("DismOpenSession")
	procDismRemoveCapability = modDismAPI.NewProc("DismRemoveCapability")
	procDismRemoveDriver     = modDismAPI.NewProc("DismRemoveDriver")
	procDismRemovePackage    = modDismAPI.NewProc("DismRemovePackage")
	procDismShutdown         = modDismAPI.NewProc("DismShutdown")
)

func DismAddCapability(Session uint32, Name *uint16, LimitAccess bool, SourcePaths **uint16, SourcePathCount uint32, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	var _p0 uint32
	if LimitAccess {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall9(procDismAddCapability.Addr(), 8, uintptr(Session), uintptr(unsafe.Pointer(Name)), uintptr(_p0), uintptr(unsafe.Pointer(SourcePaths)), uintptr(SourcePathCount), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismAddDriver(Session uint32, DriverPath *uint16, ForceUnsigned bool) (e error) {
	var _p0 uint32
	if ForceUnsigned {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procDismAddDriver.Addr(), 3, uintptr(Session), uintptr(unsafe.Pointer(DriverPath)), uintptr(_p0))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismAddPackage(Session uint32, PackagePath *uint16, IgnoreCheck bool, PreventPending bool, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	var _p0 uint32
	if IgnoreCheck {
		_p0 = 1
	}
	var _p1 uint32
	if PreventPending {
		_p1 = 1
	}
	r0, _, _ := syscall.Syscall9(procDismAddPackage.Addr(), 7, uintptr(Session), uintptr(unsafe.Pointer(PackagePath)), uintptr(_p0), uintptr(_p1), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismApplyUnattend(Session uint32, UnattendFile *uint16, SingleSession bool) (e error) {
	var _p0 uint32
	if SingleSession {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall(procDismApplyUnattend.Addr(), 3, uintptr(Session), uintptr(unsafe.Pointer(UnattendFile)), uintptr(_p0))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismCloseSession(Session uint32) (e error) {
	r0, _, _ := syscall.Syscall(procDismCloseSession.Addr(), 1, uintptr(Session), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismDisableFeature(Session uint32, FeatureName *uint16, PackageName *uint16, RemovePayload bool, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	var _p0 uint32
	if RemovePayload {
		_p0 = 1
	}
	r0, _, _ := syscall.Syscall9(procDismDisableFeature.Addr(), 7, uintptr(Session), uintptr(unsafe.Pointer(FeatureName)), uintptr(unsafe.Pointer(PackageName)), uintptr(_p0), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismEnableFeature(Session uint32, FeatureName *uint16, Identifier *uint16, PackageIdentifier *DismPackageIdentifier, LimitAccess bool, SourcePaths *string, SourcePathCount uint32, EnableAll bool, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	var _p0 uint32
	if LimitAccess {
		_p0 = 1
	}
	var _p1 uint32
	if EnableAll {
		_p1 = 1
	}
	r0, _, _ := syscall.Syscall12(procDismEnableFeature.Addr(), 11, uintptr(Session), uintptr(unsafe.Pointer(FeatureName)), uintptr(unsafe.Pointer(Identifier)), uintptr(unsafe.Pointer(PackageIdentifier)), uintptr(_p0), uintptr(unsafe.Pointer(SourcePaths)), uintptr(SourcePathCount), uintptr(_p1), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismInitialize(LogLevel DismLogLevel, LogFilePath *uint16, ScratchDirectory *uint16) (e error) {
	r0, _, _ := syscall.Syscall(procDismInitialize.Addr(), 3, uintptr(LogLevel), uintptr(unsafe.Pointer(LogFilePath)), uintptr(unsafe.Pointer(ScratchDirectory)))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismOpenSession(ImagePath *uint16, WindowsDirectory *uint16, SystemDrive *uint16, Session *uint32) (e error) {
	r0, _, _ := syscall.Syscall6(procDismOpenSession.Addr(), 4, uintptr(unsafe.Pointer(ImagePath)), uintptr(unsafe.Pointer(WindowsDirectory)), uintptr(unsafe.Pointer(SystemDrive)), uintptr(unsafe.Pointer(Session)), 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismRemoveCapability(Session uint32, Name *uint16, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	r0, _, _ := syscall.Syscall6(procDismRemoveCapability.Addr(), 5, uintptr(Session), uintptr(unsafe.Pointer(Name)), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismRemoveDriver(Session uint32, DriverPath *uint16) (e error) {
	r0, _, _ := syscall.Syscall(procDismRemoveDriver.Addr(), 2, uintptr(Session), uintptr(unsafe.Pointer(DriverPath)), 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismRemovePackage(Session uint32, Identifier *uint16, PackageIdentifier *DismPackageIdentifier, CancelEvent *windows.Handle, Progress unsafe.Pointer, UserData unsafe.Pointer) (e error) {
	r0, _, _ := syscall.Syscall6(procDismRemovePackage.Addr(), 6, uintptr(Session), uintptr(unsafe.Pointer(Identifier)), uintptr(unsafe.Pointer(PackageIdentifier)), uintptr(unsafe.Pointer(CancelEvent)), uintptr(Progress), uintptr(UserData))
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}

func DismShutdown() (e error) {
	r0, _, _ := syscall.Syscall(procDismShutdown.Addr(), 0, 0, 0, 0)
	if r0 != 0 {
		e = syscall.Errno(r0)
	}
	return
}
