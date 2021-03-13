package com.mujin.androidremoteservant.core.stf.touch;

import android.util.Log;

import java.io.IOException;
import java.io.InputStream;

public class ThreadReceiver extends Thread {

    // LocalSocket localSocket = null;
    MiniTouch threadMaster = null;
    boolean Running = false;
    String Tag = Thread.currentThread().getName();
    Thread thread = null;

    ThreadReceiver(MiniTouch miniTouch) {
        this.Running = true;
        this.threadMaster = miniTouch;
    }

    public void start() {
        if (thread == null) {
            thread = new Thread(this);
            thread.start();
        }
    }

    public void close() {
        this.Running = false;
    }

    @Override
    public void run() {
        if (threadMaster.localSocket == null) {
            Log.e(Tag, "thread no init");
            return;
        }

        InputStream mRace = null;
        try {
            mRace = threadMaster.localSocket.getInputStream();
            byte[] data;
            int receiveLen = 0;

            while (Running) {
                // receiveLen = localSocket.getReceiveBufferSize();
                data = new byte[128];
                int a = mRace.read(data);

                String ret = new String(data);
                threadMaster.setMiniTouchInfo(MiniTouchInfo.parse(ret));
                // Log.i(Tag, "receiver.getReceiveBufferSize()" + receiveLen + " --- " + ret + " ---");

            }

            mRace.close();
            // localSocket.close();

        } catch (IOException e) {
            e.printStackTrace();
        } finally {
            this.thread = null;
        }
    }
}
