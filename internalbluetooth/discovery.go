package internalbluetooth

import (
	"log"
	"time"
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

	// 存储已经发现的设备
	discoverDevices := map[string]bool{}

	for {
		// 开始扫描
		log.Println("开始扫描附近的设备...")

		// 启动扫描
		go func() {
			err = defaultAdapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
				// 检查设备是否已经发现
				if _, ok := discoverDevices[device.Address.String()]; ok {
					return
				}

				// 记录新的设备并打印信息
				discoverDevices[device.Address.String()] = true
				log.Printf("发现设备: %s, RSSI: %d, 名称: %s", device.Address.String(), device.RSSI, device.LocalName())
			})

			if err != nil {
				log.Fatalf("扫描失败: %v", err)
			}
		}()

		// 设置扫描持续时间为 10 秒
		time.Sleep(10 * time.Second)

		// 停止扫描
		log.Println("扫描时间结束，停止扫描...")
		err = defaultAdapter.StopScan()
		if err != nil {
			return err
		}

		// 休眠一段时间后重新开始扫描（例如 5 秒）
		log.Println("休息一会儿，然后重新开始扫描...")
		time.Sleep(5 * time.Second)
	}

	return nil
}
