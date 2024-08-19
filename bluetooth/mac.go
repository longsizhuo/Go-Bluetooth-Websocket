package bluetooth

///*
//#cgo CFLAGS: -x objective-c
//#cgo LDFLAGS: -framework Foundation -framework CoreBluetooth
//#import <Foundation/Foundation.h>
//#import <CoreBluetooth/CoreBluetooth.h>
//
//@interface BLEScanner : NSObject <CBCentralManagerDelegate>
//@property (nonatomic, strong) CBCentralManager *centralManager;
//@end
//
//@implementation BLEScanner
//
//- (instancetype)init {
//    self = [super init];
//    if (self) {
//        self.centralManager = [[CBCentralManager alloc] initWithDelegate:self queue:nil];
//    }
//    return self;
//}
//
//- (void)centralManagerDidUpdateState:(CBCentralManager *)central {
//    if (central.state == CBManagerStatePoweredOn) {
//        [central scanForPeripheralsWithServices:nil options:nil];
//        NSLog(@"开始扫描设备...");
//    } else {
//        NSLog(@"蓝牙不可用");
//    }
//}
//
//- (void)centralManager:(CBCentralManager *)central didDiscoverPeripheral:(CBPeripheral *)peripheral advertisementData:(NSDictionary<NSString *,id> *)advertisementData RSSI:(NSNumber *)RSSI {
//    NSLog(@"发现设备: %@, RSSI: %@", peripheral.name, RSSI);
//}
//
//@end
//
//BLEScanner* scanner;
//
//void StartBLEScan() {
//    scanner = [[BLEScanner alloc] init];
//}
//*/
//import "C"
//
//func Mac() {
//	C.StartBLEScan()
//	select {} // 保持程序运行
//}
