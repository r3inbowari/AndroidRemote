package com.mujin.androidremoteservant;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import com.mujin.androidremoteservant.core.session.DeviceProvider;
import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.core.utils.Device;

import static com.mujin.androidremoteservant.core.SystemInit.applicationEntry;
import static com.mujin.androidremoteservant.core.SystemInit.prepareRuntimeNDKEnv;

public class MainActivity extends AppCompatActivity {


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

        // System.out.println(DeviceProvider.getDeviceInfo(this));
    }
}
