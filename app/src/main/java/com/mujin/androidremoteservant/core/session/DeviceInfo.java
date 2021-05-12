package com.mujin.androidremoteservant.core.session;

import com.mujin.androidremoteservant.core.utils.Device;
import com.mujin.androidremoteservant.core.utils.Hash;
import com.mujin.androidremoteservant.core.utils.Utils;

import java.security.NoSuchAlgorithmException;
import java.util.Date;

public class DeviceInfo {
    private String deviceFingerPrint = "";
    private int deviceWidth = 0;
    private int deviceHeight = 0;
    private String deviceBrand = "";
    private int deviceSDK = 0;
    private String deviceRelease = "";
    private String deviceIMEI = "";
    private String deviceID = "";
    private String deviceMAC = "";
    private long deviceTotalMem = 0;
    private long deviceAvailMem = 0;
    private String deviceWifiAddress = "";
    private String deviceContainerID = "";

    public DeviceInfo() throws NoSuchAlgorithmException {
    }

    public String getDeviceWifiAddress() {
        return deviceWifiAddress;
    }

    public DeviceInfo setDeviceWifiAddress(String deviceWifiAddress) {
        this.deviceWifiAddress = deviceWifiAddress;
        return this;
    }

    private static DeviceInfo mInstance = null;

    public static DeviceInfo getInstance() throws NoSuchAlgorithmException {
        if (mInstance == null)
            mInstance = new DeviceInfo();
        return mInstance;
    }

    public String getDeviceRelease() {
        return deviceRelease;
    }

    public DeviceInfo setDeviceRelease(String deviceRelease) {
        this.deviceRelease = deviceRelease;
        return this;
    }

    public String getDeviceIMEI() {
        return deviceIMEI;
    }

    public DeviceInfo setDeviceIMEI(String deviceIMEI) {
        this.deviceIMEI = deviceIMEI;
        return this;
    }

    public String getDeviceMAC() {
        return deviceMAC;
    }

    public DeviceInfo setDeviceMAC(String deviceMAC) {
        this.deviceMAC = deviceMAC;
        return this;
    }

    public int getDeviceHeight() {
        return deviceHeight;
    }

    public DeviceInfo setDeviceHeight(int deviceHeight) {
        this.deviceHeight = deviceHeight;
        return this;
    }

    public String getDeviceBrand() {
        return deviceBrand;
    }

    public DeviceInfo setDeviceBrand(String deviceBrand) {
        this.deviceBrand = deviceBrand;
        return this;
    }

    public int getDeviceSDK() {
        return deviceSDK;
    }

    public DeviceInfo setDeviceSDK(int deviceSDK) {
        this.deviceSDK = deviceSDK;
        return this;
    }

    public String getDeviceID() {
        return deviceID;
    }

    public DeviceInfo setDeviceID(String deviceID) {
        this.deviceID = deviceID;
        return this;
    }

    public String getDeviceFingerPrint() {
        return deviceFingerPrint;
    }

    public DeviceInfo setDeviceFingerPrint(String deviceFingerPrint) {
        this.deviceFingerPrint = deviceFingerPrint;
        return this;
    }

    public int getDeviceWidth() {
        return deviceWidth;
    }

    public DeviceInfo setDeviceWidth(int deviceWidth) {
        this.deviceWidth = deviceWidth;
        return this;
    }

    public long getDeviceTotalMem() {
        return deviceTotalMem;
    }

    public DeviceInfo setDeviceAvailMem(long deviceAvailMem) {
        this.deviceAvailMem = deviceAvailMem;
        return this;
    }

    public long getDeviceAvailMem() {
        return deviceAvailMem;
    }

    public DeviceInfo setDeviceTotalMem(long deviceTotalMem) {
        this.deviceTotalMem = deviceTotalMem;
        return this;
    }

    public DeviceInfo setDeviceContainerID(String deviceContainerID) {
        this.deviceContainerID = deviceContainerID;
        return this;
    }

    public String getDeviceContainerID() {
        return deviceContainerID;
    }
}
