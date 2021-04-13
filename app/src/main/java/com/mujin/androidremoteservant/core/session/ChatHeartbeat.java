package com.mujin.androidremoteservant.core.session;

import android.app.Activity;
import android.util.Log;

import static com.mujin.androidremoteservant.core.SystemInit.DID;
import static com.mujin.androidremoteservant.core.session.DeviceProvider.getDeviceInfo;

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
                String bb = DeviceProvider.getDeviceInfo(context);
                Chat.getInstance().sendMsg(ChatTypeEnum.REG, bb, DID);
                Thread.sleep(120 * 1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}