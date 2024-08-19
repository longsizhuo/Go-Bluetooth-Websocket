package main

import (
	device "Go-Bluetooth-Websocket/bluetooth"
	"Go-Bluetooth-Websocket/internalbluetooth"
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"strings"
	"tinygo.org/x/bluetooth"
)

func main() {
	switch runtime.GOOS {
	//case "darwin":
	//	// 如果当前设备是 macOS
	//	device.Mac()
	//case "linux":
	//	// 如果当前设备是 Linux
	//	device.Linux()
	case "windows":
		// 如果当前设备是 Windows
		devices, err := device.GetConnectedDevicesWindows()
		if err != nil {
			println("无法获取已连接的设备:", err.Error())
			return
		}
		for _, d := range devices {
			println(d)
		}
		// Choose a device to connect to through the command line
		device_connect := bufio.NewReader(os.Stdin)
		println("请输入要连接的MAC:")
		device_name, _ := device_connect.ReadString('\n')
		device_name = device_name[:len(device_name)-1]
		// 需要转 type MAC [6]byte
		macBytes, err := parseMAC(device_name)
		if err != nil {
			fmt.Println("无效的MAC地址:", err)
			return
		}

		address := bluetooth.MACAddress{
			MAC: macBytes,
		}
		deviceAddress := bluetooth.Address{
			MACAddress: address,
		}
		err = internalbluetooth.ConnectBLEDeviceAndChat(deviceAddress)
		if err != nil {
			println("无法连接到设备:", err.Error())
			return
		}

	//	device.Windows()
	default:
		// 其他未知的操作系统
		println("不支持的操作系统:", runtime.GOOS)
	}
}

func parseMAC(mac string) ([6]byte, error) {
	var macBytes [6]byte

	// Check if the MAC address is in the format 00805F9B34FB
	if len(mac) == 12 {
		for i := 0; i < 6; i++ {
			byteVal, err := hex.DecodeString(mac[i*2 : i*2+2])
			if err != nil {
				return macBytes, err
			}
			macBytes[i] = byteVal[0]
		}
		return macBytes, nil
	}

	// Original code for MAC addresses with colons
	parts := strings.Split(mac, ":")
	if len(parts) != 6 {
		return macBytes, fmt.Errorf("MAC地址格式错误")
	}

	for i, part := range parts {
		byteVal, err := hex.DecodeString(part)
		if err != nil {
			return macBytes, err
		}
		macBytes[i] = byteVal[0]
	}

	return macBytes, nil
}
