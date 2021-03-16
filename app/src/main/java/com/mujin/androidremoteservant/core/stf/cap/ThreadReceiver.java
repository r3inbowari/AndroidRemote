package com.mujin.androidremoteservant.core.stf.cap;

import android.util.Log;

import com.mujin.androidremoteservant.core.utils.Utils;

import java.io.IOException;
import java.io.InputStream;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;
import java.util.concurrent.BlockingQueue;

public class ThreadReceiver extends Thread {

    String Tag = Thread.currentThread().getName();
    MiniCap threadMaster = null;
    boolean Running = false;
    Thread thread = null;

    // private final BlockingQueue<byte[]> taskQueue;
    private final FrameAlloc frameAlloc;

    // ThreadReceiver(MiniCap miniCap, BlockingQueue<byte[]> taskQueue) {
    ThreadReceiver(MiniCap miniCap, FrameAlloc taskQueue) {
        this.Running = true;
        this.threadMaster = miniCap;
        // this.taskQueue = taskQueue;
        this.frameAlloc = taskQueue;
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

        // 帧长预
        int frameLen = 0;

        InputStream mRace = null;
        try {
            mRace = threadMaster.localSocket.getInputStream();
            CapProbe.HeadScanner(mRace);

            // 预分配缓冲区 64 KB
            // byte[] data = new byte[256 * 1024];
            // 预分配长度空间 4 Bytes
            byte[] len = new byte[4];

            // testCount
            int testCount = 0;

            long nanoStart = System.nanoTime();

            while (Running) {
                mRace.read(len);
                frameLen = Utils.parseFrameLen(len);
                // byte[] data = new byte[frameLen];

                Frame frame = frameAlloc.alloc();

                while (frameLen != 0) {
                    int reduce = mRace.read(frame.getFrameBuffer(), 0, frameLen);
                    // long reduce = mRace.skip(frameLen);
                    frameLen -= reduce;
                }

                frameAlloc.put(frame);
                testCount++;
                if (testCount == 30) {
                    long nanoEnd = System.nanoTime();
                    System.out.println("30 FP -> Time Consuming: " + Long.toString((nanoEnd - nanoStart) / 1000000));
                    nanoStart = nanoEnd;
                    testCount = 0;
                }
            }

            mRace.close();
            // localSocket.close();

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
            System.out.println(frameLen);
        } finally {
            this.thread = null;
        }
    }
}
