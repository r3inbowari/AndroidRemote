package com.mujin.androidremoteservant.core;

import android.app.Activity;
import android.os.Build;

import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.cap.CapProbe;
import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.core.utils.AssetsOperation;
import com.mujin.androidremoteservant.core.utils.PID;
import com.mujin.androidremoteservant.core.utils.ScreenMetrics;
import com.mujin.androidremoteservant.core.utils.Utils;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;

import java.io.IOException;

import static com.mujin.androidremoteservant.core.utils.Utils.*;

public class SystemInit {
    public static void prepareRuntimeNDKEnv(Activity context) {
        // root 权限获取
        rootGrant(context.getPackageCodePath());

        // 检查并试图清除旧的相关后台进程
        PID.checkAndKillProcess(context, "minicap");
        PID.checkAndKillProcess(context, "minitouch");

        // 架构获取
        String arch = Utils.getHardwareArch();
        // 屏幕适配
        ScreenMetrics.initIfNeeded(context);

        // System.out.println(ScreenMetrics.getDeviceScreenHeight());
        // System.out.println(ScreenMetrics.getDeviceScreenWidth());

        // sdk version
        int sdk = Build.VERSION.SDK_INT;

        // minicap 的动态库加载地址
        String dllPath = "/lib/android-" + String.valueOf(sdk) + "/minicap.so";

        // 目录清除
        AssetsOperation.clean("/data/local/tmp/nohup.out");
        AssetsOperation.clean("/data/local/tmp/minitouch.log");
        AssetsOperation.clean("/data/local/tmp/minicap.log");
        AssetsOperation.clean("/data/local/tmp/minicap");
        AssetsOperation.clean("/data/local/tmp/minicap.so");
        AssetsOperation.clean("/data/local/tmp/minitouch");
        AssetsOperation.clean("/data/local/tmp/run.sh");

        String a = context.getCacheDir().getPath();

        // 文件权限处理
        AssetsOperation.autoZipBySU(context, arch + "/bin/minitouch", DIR_TEMP, EXEC_MINI_TOUCH);
        AssetsOperation.autoZipBySU(context, "run.sh", DIR_TEMP, "run.sh");
        AssetsOperation.autoZipBySU(context, arch + "/bin/minicap", DIR_TEMP, EXEC_MINI_CAP);
        AssetsOperation.autoZipBySU(context, arch + dllPath, DIR_TEMP, DLL_MINI_CAP);

        // 这里进行一次 autojs 中的滑动测试
        ProcessShell ps = new ProcessShell(true);
        ps.Swipe(100, 200, 400, 500);
    }

    // 程序入口
    public static void applicationEntry(Activity context) {
        // gRPC pool 初始化
        gRPCChannelPool.init("192.168.5.67", 5005);

        // 后台相关进程的启动
        MiniTouch.getInstance().run();
        MiniCap.getInstance().run();

        // 启动 minicap 的发送任务
        MiniCap.getInstance().startSender();

        try {
            // TMP ANR
            Thread.sleep(1000);
            MiniTouch.getInstance().connect();
            MiniTouch.getInstance().getLocalReceiver().start();
            MiniTouch.getInstance().getLocalSender().start();

            MiniCap.getInstance().connect();
            MiniCap.getInstance().getLocalReceiver().start();
            Thread.sleep(1000);
        } catch (InterruptedException | IOException e) {
            // sleep interruptedException
            // thread io exception
            e.printStackTrace();
        }

        // 记录本次 PID
        if (CapProbe.pid != 0 && MiniTouch.getInstance().getMiniTouchInfo() != null) {
            PID.Attach(context, "minicap", CapProbe.pid);
            PID.Attach(context, "minitouch", Integer.valueOf(MiniTouch.getInstance().getMiniTouchInfo().PID));
        }
    }
}
