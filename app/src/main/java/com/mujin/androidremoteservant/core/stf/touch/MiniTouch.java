package com.mujin.androidremoteservant.core.stf.touch;

import android.text.TextUtils;
import android.util.Log;


import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.touch.event.EventManager;

import java.io.DataOutputStream;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

public class MiniTouch extends AbstractMiniTouch {

    private static MiniTouch sInstance = null;

    private final BlockingQueue<String> taskQueue;

    private MiniTouchInfo miniTouchInfo = null;

    private EventManager eventManager = null;

    private final String TAG = "MiniTouch";

    public MiniTouchInfo getMiniTouchInfo() {
        return miniTouchInfo;
    }

    public MiniTouch() {
        super("minitouch");
        this.taskQueue = new ArrayBlockingQueue<>(100);
        this.eventManager = new EventManager(taskQueue);

        new Thread(() -> {
            while (true) {
                try {
                    Thread.sleep(3000);
                    this.eventManager.tap(0, 400, 400);
                    Thread.sleep(16);
                    this.eventManager.release(0);

                    this.eventManager.move(3, 400, 200);
                    Thread.sleep(1621);
                    this.eventManager.move(3, 410, 210);
                    Thread.sleep(1621);
                    this.eventManager.move(3, 420, 230);
                    Thread.sleep(1621);
                    this.eventManager.move(3, 450, 220);
                    Thread.sleep(1621);
                    this.eventManager.move(3, 460, 320);
                    Thread.sleep(1621);
                    this.eventManager.release(3);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }

        });

    }

    public EventManager getEventManager() {
        return this.eventManager;
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
        // taskQueue.put(EventUtils.move(0, 270, 310, 0));
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
