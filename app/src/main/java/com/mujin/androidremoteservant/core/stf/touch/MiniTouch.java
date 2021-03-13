package com.mujin.androidremoteservant.core.stf.touch;

import android.text.TextUtils;
import android.util.Log;


import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.touch.event.EventUtils;

import java.io.DataOutputStream;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

public class MiniTouch extends AbstractMiniTouch {

    private static MiniTouch sInstance = null;

    private final BlockingQueue<String> taskQueue;

    private MiniTouchInfo miniTouchInfo = null;

    private final String TAG = "MiniTouch";

    public MiniTouch() {
        super("minitouch");
        taskQueue = new ArrayBlockingQueue<>(100);

        new Thread(() -> {
            while (true) {
                try {
                    Thread.sleep(3000);
                    taskQueue.put(EventUtils.tap(0, 200, 300, 1));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 210, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 220, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 230, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 240, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 250, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 260, 310, 0));
                    Thread.sleep(10);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(10);

                    taskQueue.put(EventUtils.move(0, 270, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 280, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 290, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 300, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);

                    taskQueue.put(EventUtils.move(0, 310, 310, 0));
                    Thread.sleep(4);
                    taskQueue.put(EventUtils.commit());
                    Thread.sleep(4);


                    taskQueue.put(EventUtils.release(0));

                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

        }).start();

    }

    public static MiniTouch getInstance() {
        if (sInstance == null) {
            synchronized (MiniTouch.class) {
                if (sInstance == null) {
                    sInstance = new MiniTouch();
                }
            }
        }
        return sInstance;
    }

    public void setMiniTouchInfo(MiniTouchInfo miniTouchInfo) {
        this.miniTouchInfo = miniTouchInfo;
    }

    @Override
    public ThreadReceiver getLocalReceiver() {
        if (threadReceiver == null)
            this.threadReceiver = new ThreadReceiver(this);
        return threadReceiver;
    }

    @Override
    public ThreadSender getLocalSender() {
        if (threadSender == null)
            this.threadSender = new ThreadSender(this, taskQueue);
        return threadSender;
    }

    @Override
    public void touch(int contact, int x, int y, int pressure) {

    }

    @Override
    public void swipe(int contact, int x, int y, int duration) {

    }

    @Override
    public void executeTouch(String touchCommand) {

    }

    public void destroy() {
        if (this.miniTouchInfo != null) {
            Log.i(TAG, "destroy and release ndk resource -> pid ---- " + this.miniTouchInfo.PID);
            ProcessShell.exec(TextUtils.join(" ", new Object[]{"kill", this.miniTouchInfo.PID}), true);
        }
    }

    @Override
    public void run() {
        new Thread(new MiniTouchNDK(), "MiniTouch NDK").start();
    }
}

class MiniTouchNDK extends Thread {

    static final String TAG = "Binary Touch";

    @Override
    public void run() {
        Process process = null;
        DataOutputStream os = null;
        try {
            process = Runtime.getRuntime().exec("su");
            os = new DataOutputStream(process.getOutputStream());
            os.writeBytes("cd /data/local/tmp" + "\n");
            os.flush();
            os.writeBytes("./minitouch > minitouch.log 2>&1" + "\n");
            os.writeBytes("exit\n");
            os.flush();
            process.waitFor();
        } catch (Exception e) {
            Log.d("DEBUG", e.getMessage());
            return;
        } finally {
            try {
                if (os != null) {
                    os.close();
                }
                process.destroy();
            } catch (Exception e) {
                Log.d(TAG, "os exception warming -> " + Thread.currentThread().getName());
            }
        }
        Log.d(TAG, "running mini touch on thread -> " + Thread.currentThread().getName());
    }
}
