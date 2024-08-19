package internalbluetooth

import (
	"log"
	"tinygo.org/x/bluetooth"
)

var defaultAdapter = bluetooth.DefaultAdapter

// DiscoverDevices scans for Bluetooth devices and logs their addresses and names.
func DiscoverDevices() error {
	// 启用蓝牙适配器
	err := defaultAdapter.Enable()
	if err != nil {
		return err
	}

	// 开始扫描
	log.Println("开始扫描附近的设备...")
	err = defaultAdapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		log.Printf("发现设备: %s, RSSI: %d, 名称: %s", device.Address.String(), device.RSSI, device.LocalName())
	})

	if err != nil {
		return err
	}

	return nil
}
