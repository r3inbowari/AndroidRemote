package com.mujin.androidremoteservant.core.session;

import android.app.Activity;
import android.util.Log;

import java.security.NoSuchAlgorithmException;

import static com.mujin.androidremoteservant.core.SystemInit.DID;

public class ChatHeartbeat implements Runnable {

    Activity context = null;
    private static final String TAG = "CHB";

    public ChatHeartbeat(Activity context) {
        this.context = context;
    }

    @Override
    public void run() {
        Log.i(TAG, "heartbeat start");
        while (true) {
            try {
                Chat.getInstance().sendMsg(ChatTypeEnum.REG, DeviceProvider.getDeviceInfo(context));
                Thread.sleep(80 * 1000);
            } catch (InterruptedException | NoSuchAlgorithmException e) {
                e.printStackTrace();
            }
        }
    }
}