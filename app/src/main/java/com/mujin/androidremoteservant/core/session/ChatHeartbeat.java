package com.mujin.androidremoteservant.core.session;

import android.app.Activity;
import android.util.Log;

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
                Chat.getInstance().sendMsg(ChatTypeEnum.REG, DeviceProvider.getDeviceInfo(context), DID);
                Thread.sleep(80 * 1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}