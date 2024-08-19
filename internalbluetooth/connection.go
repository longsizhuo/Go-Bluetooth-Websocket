package internalbluetooth

import (
	"log"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

// ConnectToDevice connects to a Bluetooth device by its address.
func ConnectToDevice(address string) (*bluetooth.Device, error) {
	// 启用蓝牙适配器
	err := adapter.Enable()
	if err != nil {
		return nil, err
	}

	// 将字符串形式的地址解析为 BLE 地址对象
	addr, err := bluetooth.ParseMAC(address)
	if err != nil {
		return nil, err
	}

	// 尝试连接到设备
	log.Printf("尝试连接到设备: %s", address)
	device, err := adapter.Connect(addr, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, err
	}

	log.Printf("成功连接到设备: %s", address)
	return device, nil
}
