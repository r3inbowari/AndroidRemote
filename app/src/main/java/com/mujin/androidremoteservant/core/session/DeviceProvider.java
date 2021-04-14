package com.mujin.androidremoteservant.core.session;

import android.app.Activity;

import com.alibaba.fastjson.JSON;
import com.mujin.androidremoteservant.core.utils.Device;
import com.mujin.androidremoteservant.core.utils.WifiTool;

public class DeviceProvider {

    public static String getDeviceInfo(Activity context) {
        Device device = new Device(context);
        try {
            DeviceInfo.getInstance()
                    .setDeviceID(device.getAndroidId())
                    .setDeviceAvailMem(device.getAvailMem())
                    .setDeviceBrand(Device.brand)
                    .setDeviceFingerPrint(Device.fingerprint)
                    .setDeviceHeight(Device.height)
                    .setDeviceWidth(Device.width)
                    .setDeviceMAC(device.getMacAddress())
                    .setDeviceIMEI(device.getIMEI())
                    .setDeviceRelease(Device.release)
                    .setDeviceSDK(Device.sdkInt)
                    .setDeviceTotalMem(device.getTotalMem())
                    .setDeviceWifiAddress(WifiTool.getWifiAddress(context));
        } catch (Exception e) {
            System.out.println("found errors");
        }
        // marshal json
        return JSON.toJSONString(DeviceInfo.getInstance());
    }
}
