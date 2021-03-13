package com.mujin.androidremoteservant;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;

import java.io.IOException;

import static com.mujin.androidremoteservant.core.SystemInit.prepareRuntimeNDKEnv;
import static com.mujin.androidremoteservant.core.utils.Utils.rootGrant;

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

        rootGrant(getPackageCodePath());

        System.out.println(getPackageCodePath());

        prepareRuntimeNDKEnv(this);

        ProcessShell ps = new ProcessShell(true);
        ps.Swipe(100, 200, 400, 500);


        MiniTouch.getInstance().run();
        MiniCap.getInstance().run();

        try {
            // TMP ANR
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        try {

            MiniTouch.getInstance().connect();
            MiniTouch.getInstance().getLocalReceiver().start();
            MiniTouch.getInstance().getLocalSender().start();

            MiniCap.getInstance().connect();
            MiniCap.getInstance().getLocalReceiver().start();

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}