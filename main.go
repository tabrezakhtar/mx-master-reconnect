package main

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Windows API constants for Bluetooth
const (
	ERROR_SUCCESS           = 0
	ERROR_NO_MORE_ITEMS     = 259
	BLUETOOTH_MAX_NAME_SIZE = 248
)

// BLUETOOTH_DEVICE_INFO structure
type BLUETOOTH_DEVICE_INFO struct {
	Size          uint32
	Address       uint64
	ClassofDevice uint32
	Connected     int32
	Remembered    int32
	Authenticated int32
	LastSeen      windows.Filetime
	LastUsed      windows.Filetime
	Name          [BLUETOOTH_MAX_NAME_SIZE]uint16
}

// BLUETOOTH_DEVICE_SEARCH_PARAMS structure
type BLUETOOTH_DEVICE_SEARCH_PARAMS struct {
	Size                uint32
	ReturnAuthenticated int32
	ReturnRemembered    int32
	ReturnUnknown       int32
	ReturnConnected     int32
	IssueInquiry        int32
	TimeoutMultiplier   uint8
	RadioHandle         windows.Handle
}

func main() {
	fmt.Println("Scanning for Bluetooth devices...")

	// Bluetooth class GUID: {e0cbf06c-cd8b-4647-bb8a-263b43f0f974}
	guidBluetooth := windows.GUID{
		Data1: 0xe0cbf06c,
		Data2: 0xcd8b,
		Data3: 0x4647,
		Data4: [8]byte{0xbb, 0x8a, 0x26, 0x3b, 0x43, 0xf0, 0xf9, 0x74},
	}

	setupapi := syscall.NewLazyDLL("setupapi.dll")
	procSetupDiGetClassDevs := setupapi.NewProc("SetupDiGetClassDevsW")
	procSetupDiEnumDeviceInfo := setupapi.NewProc("SetupDiEnumDeviceInfo")
	procSetupDiGetDeviceRegistryProperty := setupapi.NewProc("SetupDiGetDeviceRegistryPropertyW")
	procSetupDiDestroyDeviceInfoList := setupapi.NewProc("SetupDiDestroyDeviceInfoList")

	DIGCF_PRESENT := 0x2

	// Get device info set for Bluetooth devices
	hDevInfo, _, _ := procSetupDiGetClassDevs.Call(
		uintptr(unsafe.Pointer(&guidBluetooth)),
		0,
		0,
		uintptr(DIGCF_PRESENT),
	)
	if hDevInfo == 0 {
		fmt.Println("Failed to get device info set.")
		return
	}
	defer procSetupDiDestroyDeviceInfoList.Call(hDevInfo)

	type SP_DEVINFO_DATA struct {
		cbSize    uint32
		ClassGuid windows.GUID
		DevInst   uint32
		Reserved  uintptr
	}

	var deviceCount int
	for i := 0; ; i++ {
		devInfoData := SP_DEVINFO_DATA{cbSize: uint32(unsafe.Sizeof(SP_DEVINFO_DATA{}))}
		ret, _, _ := procSetupDiEnumDeviceInfo.Call(
			hDevInfo,
			uintptr(i),
			uintptr(unsafe.Pointer(&devInfoData)),
		)
		if ret == 0 {
			break // No more devices
		}

		// Get device name
		const SPDRP_FRIENDLYNAME = 0x0000000C
		var nameBuf [256]uint16
		var reqLen uint32
		procSetupDiGetDeviceRegistryProperty.Call(
			hDevInfo,
			uintptr(unsafe.Pointer(&devInfoData)),
			SPDRP_FRIENDLYNAME,
			0,
			uintptr(unsafe.Pointer(&nameBuf[0])),
			uintptr(len(nameBuf)*2),
			uintptr(unsafe.Pointer(&reqLen)),
		)
		name := windows.UTF16ToString(nameBuf[:])
		if name == "" {
			name = "Unknown Device"
		}

		fmt.Printf("\n--- Device %d ---\n", deviceCount+1)
		fmt.Printf("Name: %s\n", name)
		// MAC address and other info are not directly available via SetupAPI
		deviceCount++
	}

	if deviceCount == 0 {
		fmt.Println("No Bluetooth devices found.")
	} else {
		fmt.Printf("\nTotal devices found: %d\n", deviceCount)
	}

	var mxMasterFound bool
	var mxMasterIndex int
	var mxMasterDevInfoData SP_DEVINFO_DATA

	for i := 0; ; i++ {
		devInfoData := SP_DEVINFO_DATA{cbSize: uint32(unsafe.Sizeof(SP_DEVINFO_DATA{}))}
		ret, _, _ := procSetupDiEnumDeviceInfo.Call(
			hDevInfo,
			uintptr(i),
			uintptr(unsafe.Pointer(&devInfoData)),
		)
		if ret == 0 {
			break // No more devices
		}

		// Get device name
		const SPDRP_FRIENDLYNAME = 0x0000000C
		var nameBuf [256]uint16
		var reqLen uint32
		procSetupDiGetDeviceRegistryProperty.Call(
			hDevInfo,
			uintptr(unsafe.Pointer(&devInfoData)),
			SPDRP_FRIENDLYNAME,
			0,
			uintptr(unsafe.Pointer(&nameBuf[0])),
			uintptr(len(nameBuf)*2),
			uintptr(unsafe.Pointer(&reqLen)),
		)
		name := windows.UTF16ToString(nameBuf[:])
		if name == "" {
			name = "Unknown Device"
		}

		if !mxMasterFound && strings.Contains(strings.ToLower(name), "mx master") {
			mxMasterFound = true
			mxMasterIndex = i + 1
			mxMasterDevInfoData = devInfoData
		}
	}

	if !mxMasterFound {
		fmt.Println("MX Master mouse not found. Exiting.")
		return
	}

	fmt.Printf("Found MX Master mouse (device %d). Attempting to unpair...\n", mxMasterIndex)

	// Remove the device using SetupDiCallClassInstaller with DIF_REMOVE
	const DIF_REMOVE = 0x00000005
	setupDiCallClassInstaller := setupapi.NewProc("SetupDiCallClassInstaller")
	ret, _, err := setupDiCallClassInstaller.Call(
		DIF_REMOVE,
		hDevInfo,
		uintptr(unsafe.Pointer(&mxMasterDevInfoData)),
	)
	if ret == 0 {
		fmt.Printf("Failed to remove MX Master mouse: %v\n", err)
		return
	}
	fmt.Println("MX Master mouse has been unpaired and removed.")
}

func printDeviceInfo(deviceInfo *BLUETOOTH_DEVICE_INFO, index int) {
	// Convert device name from UTF-16 to string
	name := windows.UTF16ToString(deviceInfo.Name[:])
	if name == "" {
		name = "Unknown Device"
	}

	// Format MAC address
	address := deviceInfo.Address
	macAddress := fmt.Sprintf("%02X:%02X:%02X:%02X:%02X:%02X",
		byte(address&0xFF),
		byte((address>>8)&0xFF),
		byte((address>>16)&0xFF),
		byte((address>>24)&0xFF),
		byte((address>>32)&0xFF),
		byte((address>>40)&0xFF))

	fmt.Printf("\n--- Device %d ---\n", index)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("MAC Address: %s\n", macAddress)
	fmt.Printf("Class of Device: 0x%08X\n", deviceInfo.ClassofDevice)
	fmt.Printf("Connected: %t\n", deviceInfo.Connected != 0)
	fmt.Printf("Remembered: %t\n", deviceInfo.Remembered != 0)
	fmt.Printf("Authenticated: %t\n", deviceInfo.Authenticated != 0)
}
