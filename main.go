package main

import (
	device "Go-Bluetooth-Websocket/bluetooth"
	"runtime"
)

func main() {
	switch runtime.GOOS {
	case "darwin":
		// 如果当前设备是 macOS
		device.Mac()
	case "android":
		// 如果当前设备是 Android
		device.Android()
	//case "linux":
	//	// 如果当前设备是 Linux
	//	device.Linux()
	//case "windows":
	//	// 如果当前设备是 Windows
	//	device.Windows()
	default:
		// 其他未知的操作系统
		println("不支持的操作系统:", runtime.GOOS)
	}
}
