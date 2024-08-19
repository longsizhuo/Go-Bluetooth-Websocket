package internalbluetooth

import (
	"fmt"
	"tinygo.org/x/bluetooth"
)

func ConnectBLEDeviceAndChat(address bluetooth.Address) error {
	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	if err != nil {
		return err
	}

	// 连接到BLE设备
	device, err := adapter.Connect(address, bluetooth.ConnectionParams{})
	if err != nil {
		return err
	}
	defer func(device bluetooth.Device) {
		err := device.Disconnect()
		if err != nil {
			return
		}
	}(device)

	// 发现服务
	services, err := device.DiscoverServices(nil)
	if err != nil {
		return err
	}

	for _, service := range services {
		// 发现特征
		characteristics, err := service.DiscoverCharacteristics(nil)
		if err != nil {
			return err
		}

		for _, char := range characteristics {
			fmt.Printf("发现特征: %s\n", char.UUID().String())

			// 写入消息到特征
			message := []byte("Hello from Go BLE!")
			_, err = char.WriteWithoutResponse(message)
			if err != nil {
				return err
			}
			fmt.Println("消息已发送!")

			// 读取特征的响应
			buf := make([]byte, 512) // 512字节的缓冲区用于读取消息
			n, err := char.Read(buf)
			if err != nil {
				return err
			}
			fmt.Printf("收到消息: %s\n", string(buf[:n]))
		}
	}
	return nil
}
