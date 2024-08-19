package bluetooth

//
///*
//#cgo LDFLAGS: -landroid -llog
//#include <jni.h>
//#include <android/log.h>
//
//JNIEnv* getJNIEnv();
//
//void StartBLEScan(JNIEnv* env, jobject obj) {
//    jclass bluetoothManagerClass = (*env)->FindClass(env, "android/bluetooth/BluetoothAdapter");
//    jmethodID getDefaultAdapterMethod = (*env)->GetStaticMethodID(env, bluetoothManagerClass, "getDefaultAdapter", "()Landroid/bluetooth/BluetoothAdapter;");
//    jobject bluetoothAdapter = (*env)->CallStaticObjectMethod(env, bluetoothManagerClass, getDefaultAdapterMethod);
//
//    jclass adapterClass = (*env)->GetObjectClass(env, bluetoothAdapter);
//    jmethodID startDiscoveryMethod = (*env)->GetMethodID(env, adapterClass, "startDiscovery", "()Z");
//    jboolean result = (*env)->CallBooleanMethod(env, bluetoothAdapter, startDiscoveryMethod);
//
//    if (result == JNI_TRUE) {
//        __android_log_print(ANDROID_LOG_INFO, "GoBluetooth", "开始蓝牙设备发现");
//    } else {
//        __android_log_print(ANDROID_LOG_ERROR, "GoBluetooth", "无法开始蓝牙设备发现");
//    }
//}
//*/
//import "C"
//
//func Android() {
//	// 需要确保这是在 Android 环境下运行的，并且有适当的 JNI 调用
//	env := C.getJNIEnv()
//	C.StartBLEScan(env, nil)
//}
