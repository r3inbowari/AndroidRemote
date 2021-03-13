package com.mujin.androidremoteservant.core.stf.cap;

import android.text.TextUtils;
import android.util.Log;

import com.mujin.androidremoteservant.core.utils.ScreenMetrics;
import com.mujin.androidremoteservant.core.utils.Utils;

import java.io.DataOutputStream;

public class MiniCap extends AbstractMiniCap {

    private static MiniCap sInstance = null;

    MiniCap() {
        super("minicap");
    }

    @Override
    public ThreadReceiver getLocalReceiver() {
        if (threadReceiver == null)
            this.threadReceiver = new ThreadReceiver(this);
        return threadReceiver;
    }

    public static MiniCap getInstance() {
        if (sInstance == null) {
            synchronized (MiniCap.class) {
                if (sInstance == null) {
                    sInstance = new MiniCap();
                }
            }
        }
        return sInstance;
    }

    @Override
    public void run() {
        MiniCapNDK rt = new MiniCapNDK(
                ScreenMetrics.getDeviceScreenWidth(),
                ScreenMetrics.getDeviceScreenHeight()
        );
        new Thread(
                rt,
                "MiniCap NDK"
        ).start();
    }
}

class MiniCapNDK extends Thread {

    static final String TAG = "Binary Cap";
    static final String ENV_PATH = "LD_LIBRARY_PATH=/data/local/tmp";
    static final String execName = "minicap";

    static int mx;
    static int my;
    static String execString;
    static String logStr = "> " + execName + ".log 2>&1";

    MiniCapNDK(int width, int height) {
        mx = width;
        my = height;
        if (mx > 5000 || my > 5000) {
            Log.e(TAG, "run failed used a error x or y metrics pixel");
        }
        String dp = "-P " + mx + "x" + my + "@" + mx + "x" + my + "/0";

        execString = TextUtils.join(" ",
                new Object[]{
                        ENV_PATH,
                        Utils.DIR_TEMP + "/" + execName,
                        dp,
                        logStr
                });
    }

    // LD_LIBRARY_PATH=/data/local/tmp /data/local/tmp/minicap -P 123x123@123x123/0 > minicap.log 2>&1
    @Override
    public void run() {
        Process process = null;
        DataOutputStream os = null;
        try {
            process = Runtime.getRuntime().exec("su");
            os = new DataOutputStream(process.getOutputStream());
            os.writeBytes("cd /data/local/tmp" + "\n");
            os.flush();
            os.writeBytes(execString + "\n");
            os.writeBytes("exit\n");
            os.flush();
            process.waitFor();
        } catch (Exception e) {
            Log.d("DEBUG", e.getMessage());
            return;
        } finally {
            try {
                if (os != null) {
                    os.close();
                }
                process.destroy();
            } catch (Exception e) {
                Log.d(TAG, "os exception warming -> " + Thread.currentThread().getName());
            }
        }
        Log.d(TAG, "running mini cap on thread -> " + Thread.currentThread().getName());
    }
}

