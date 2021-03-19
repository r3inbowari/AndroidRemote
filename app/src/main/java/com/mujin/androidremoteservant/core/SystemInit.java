package com.mujin.androidremoteservant.core;

import android.app.Activity;
import android.os.Build;

import com.mujin.androidremoteservant.core.utils.AssetsOperation;
import com.mujin.androidremoteservant.core.utils.PID;
import com.mujin.androidremoteservant.core.utils.ScreenMetrics;
import com.mujin.androidremoteservant.core.utils.Utils;

import static com.mujin.androidremoteservant.core.utils.Utils.*;

public class SystemInit {
    public static void prepareRuntimeNDKEnv(Activity context) {

        PID.checkAndKillProcess(context, "minicap");
        PID.checkAndKillProcess(context, "minitouch");

        String arch = Utils.getHardwareArch();

        ScreenMetrics.initIfNeeded(context);

        // System.out.println(ScreenMetrics.getDeviceScreenHeight());
        // System.out.println(ScreenMetrics.getDeviceScreenWidth());

        int sdk = Build.VERSION.SDK_INT;
        String dllPath = "/lib/android-" + String.valueOf(sdk) + "/minicap.so";

        AssetsOperation.clean("/data/local/tmp/nohup.out");
        AssetsOperation.clean("/data/local/tmp/minitouch.log");
        AssetsOperation.clean("/data/local/tmp/minicap.log");
        AssetsOperation.clean("/data/local/tmp/minicap");
        AssetsOperation.clean("/data/local/tmp/minicap.so");
        AssetsOperation.clean("/data/local/tmp/minitouch");
        AssetsOperation.clean("/data/local/tmp/run.sh");

        String a = context.getCacheDir().getPath();
        AssetsOperation.autoZipBySU(context, arch + "/bin/minitouch", DIR_TEMP, EXEC_MINI_TOUCH);
        AssetsOperation.autoZipBySU(context, "run.sh", DIR_TEMP, "run.sh");

        AssetsOperation.autoZipBySU(context, arch + "/bin/minicap", DIR_TEMP, EXEC_MINI_CAP);
        AssetsOperation.autoZipBySU(context, arch + dllPath, DIR_TEMP, DLL_MINI_CAP);

    }
}
