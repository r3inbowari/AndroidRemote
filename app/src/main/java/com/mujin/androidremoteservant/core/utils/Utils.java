package com.mujin.androidremoteservant.core.utils;

import android.content.Context;
import android.os.Environment;
import android.text.TextUtils;
import android.util.Log;

import com.mujin.androidremoteservant.core.shell.AbstractShell;
import com.mujin.androidremoteservant.core.shell.ProcessShell;

public class Utils {

    public static final String ARCH_X86 = "x86";
    public static final String ARCH_X86_64 = "x86_64";
    public static final String ARCH_MIPS = "mips";
    public static final String ARCH_MIPS64 = "mips64";
    public static final String ARCH_ARMEABI_V7A = "armeabi-v7a";
    public static final String ARCH_ARMEABI = "armeabi";
    public static final String ARCH_ARM64_V8A = "arm64_v8a";
    public static final String ARCH_UNKNOWN = "arm64_v8a";

    public static final String ORDER_CHMOD = "chmod";
    public static final String ORDER_PUSH = "push";
    public static final String ORDER_LS = "ls";
    public static final String ORDER_CD = "cd";
    public final static String ORDER_CHOWN = "chown";
    public final static String ORDER_CHGRP = "chgrp";

    public static final String EXEC_MINI_TOUCH = "minitouch";
    public static final String EXEC_MINI_CAP = "minicap";
    public static final String DLL_MINI_CAP = "minicap.so";

    public final static String LEVEL_ALL = "777";

    public static final String DIR_BINARY = "libs/";
    public static final String DIR_TEMP = "/data/local/tmp";

    public final static String OWNER_ROOT = "root";
    public final static String GROUP_ROOT = "root";
    public final static String ORDER_DEL = "rm -rf";

    public static final String TAG = "SysUtils";

    /**
     * 设备架构获取
     *
     * @return ARCH
     */
    public static String getHardwareArch() {
        AbstractShell.Result a = ProcessShell.exec("getprop ro.product.cpu.abi", false);
        if (a.code == 0)
            return a.result.substring(0, a.result.length() - 1);
        else
            return ARCH_UNKNOWN;
    }


    /**
     * root 权限提升
     *
     * @param code  提升地址
     * @param level 提升程度
     */
    public static void rootGrant(String code, String level) {
        ProcessShell.exec(TextUtils.join(" ", new Object[]{"chmod", level, code}), true);
    }

    public static void rootGrant(String code) {
        rootGrant(code, LEVEL_ALL);
    }

    /**
     * 文件推送 假的。。。
     *
     * @deprecated 弃用的方法
     */
    public static boolean pushFile() {
        String abi = getHardwareArch();
        Log.i(TAG, "ABI -> " + abi);

        if (abi.equals(ARCH_ARMEABI))
            return false;

        // String pushString = DIR_BINARY + abi + "/bin/" + EXEC_MINI_TOUCH;
        Object[] pushObject = new Object[]{
                ORDER_PUSH,
                DIR_BINARY + abi + "/bin/" + EXEC_MINI_TOUCH,
                DIR_TEMP
        };
        String pushString = TextUtils.join(" ", pushObject);

        AbstractShell.Result result = ProcessShell.exec(pushString, false);
        Log.i(TAG, "ABI2 -> " + result.result);
        // ProcessShell.exec("cd /data/local/tmp ", false);
        result = ProcessShell.exec("cd data/local/tmp & ls -l", false);
        Log.i(TAG, "ABI2 -> " + result.result);
        return true;
    }


    private static final char[] hexCode = "0123456789ABCDEF".toCharArray();

    /**
     * hex dump
     *
     * @param bytes
     * @return
     */
    public static String printHexBinary(byte[] bytes) {
        String strHex = "";
        StringBuilder sb = new StringBuilder("");
        for (int n = 0; n < bytes.length; n++) {
            strHex = Integer.toHexString(bytes[n] & 0xFF);
            sb.append((strHex.length() == 1) ? "0" + strHex : strHex);
            sb.append(" ");
        }
        return sb.toString().trim().toUpperCase();
    }

    /**
     * bytes 截取
     *
     * @param src
     * @param begin
     * @param count
     * @return
     */
    public static byte[] subBytes(byte[] src, int begin, int count) {
        byte[] bs = new byte[count];
        System.arraycopy(src, begin, bs, 0, count);
        return bs;
    }

    /**
     * 缓存地址获取
     *
     * @param context
     * @return
     */
    public String getDiskCacheDir(Context context) {
        String cachePath = null;
        if (Environment.MEDIA_MOUNTED.equals(Environment.getExternalStorageState())
                || !Environment.isExternalStorageRemovable()) {
            cachePath = context.getExternalCacheDir().getPath();
        } else {
            cachePath = context.getCacheDir().getPath();
        }
        return cachePath;
    }

    /**
     * 大小端转换 byte[4] -> int
     *
     * @param src
     * @param offset
     * @return
     */
    public static int bytesToIntByLE(byte[] src, int offset) {
        return ((src[offset + 3] & 0xFF) << 24)
                | ((src[offset + 2] & 0xFF) << 16)
                | ((src[offset + 1] & 0xFF) << 8)
                | (src[offset] & 0xFF);
    }

    public static int bytesToIntByBE(byte[] src, int offset) {
        return ((src[offset] & 0xFF) << 24)
                | ((src[offset + 1] & 0xFF) << 16)
                | ((src[offset + 2] & 0xFF) << 8)
                | (src[offset + 3] & 0xFF);
    }

    /**
     * 快速帧长度扫描 max 16777215
     *
     * @param src 长度一定是 4,并且不可以做其他用途
     * @return
     */
    public static int parseFrameLen(byte[] src) {
        return ((src[2] & 0xFF) << 16)
                | ((src[1] & 0xFF) << 8)
                | (src[0] & 0xFF);
    }

    // -9
    public static void kill(int pid) {
        ProcessShell.exec("kill -9 " + String.valueOf(pid), true);
    }
}
