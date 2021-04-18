package com.mujin.androidremoteservant;

import androidx.appcompat.app.AppCompatActivity;

import android.app.ActivityManager;
import android.content.Context;
import android.os.Bundle;

import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.core.utils.app.AppUtils;

import static com.mujin.androidremoteservant.core.SystemInit.applicationEntry;
import static com.mujin.androidremoteservant.core.SystemInit.prepareRuntimeNDKEnv;

// demo 仅供测试，不保活
public class MainActivity extends AppCompatActivity {
    static {
        System.loadLibrary("rabbit-lib");
    }

    public native String stringFromJNI();

    @Override
    protected void onDestroy() {
        MiniTouch.getInstance().destroy();
        super.onDestroy();
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        // may be start on a new thread (avoid ANR)...


        prepareRuntimeNDKEnv(this);
        new Thread(() -> {
            applicationEntry(this);
        }).start();

//         AppUtils a = new AppUtils(this);
//         a.launchApp("辐射避难所");


         // System.out.println(stringFromJNI());
        ActivityManager aam = (ActivityManager) getSystemService(Context.ACTIVITY_SERVICE);
        ProcessShell ps = new ProcessShell(true);

//        new Thread(() -> {
//            try {
//                Thread.sleep(20000);
//                ps.Home();
//                Thread.sleep(7000);
//            } catch (InterruptedException e) {
//                e.printStackTrace();
//            }
//            aam.killBackgroundProcesses("com.shandagames.falloutshelterUc.uc");
//
//        }).start();
    }
}
