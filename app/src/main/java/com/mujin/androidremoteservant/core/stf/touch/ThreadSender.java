package com.mujin.androidremoteservant.core.stf.touch;

import android.util.Log;

import java.io.IOException;
import java.io.OutputStream;
import java.util.concurrent.BlockingQueue;

public class ThreadSender extends Thread {

    MiniTouch threadMaster = null;

    private final BlockingQueue<String> taskQueue;

    boolean Running = false;
    String Tag = Thread.currentThread().getName();
    Thread thread = null;

    @Override
    public void run() {
        OutputStream mRace = null;
        try {
            mRace = threadMaster.localSocket.getOutputStream();
            while (Running) {
                String event = taskQueue.take();
                executeTask(mRace, event);
            }

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
        }

    }

    public void start() {
        if (thread == null) {
            thread = new Thread(this);
            thread.start();
            Log.i(Tag, "sender start");
        }
    }

    public void close() {
        this.Running = false;
    }

    private void executeTask(OutputStream os, String n) throws IOException {
        os.write(n.getBytes());
        os.flush();
        // System.out.println("Thread:" + Thread.currentThread().getName() + " consume: " + n);
    }

    ThreadSender(MiniTouch miniTouchUDS, BlockingQueue<String> taskQueue) {
        this.threadMaster = miniTouchUDS;
        this.taskQueue = taskQueue;
        this.Running = true;
    }

    public void Send() {

    }
}
