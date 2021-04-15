package com.mujin.androidremoteservant.core.stf.cap;

import android.util.Log;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

// unsafe
public class FrameAlloc {

    private final BlockingQueue<Frame> taskQueue;
    private final List<Frame> framePool;
    private final String TAG = "ALLOCATOR";

    public FrameAlloc(int capacity, int bufferCap) {
        Log.i(TAG, "--------------------- frame allocator ---------------------");
        this.framePool = new ArrayList<Frame>();
        taskQueue = new ArrayBlockingQueue<>(capacity);
        for (int i = 0; i < capacity + 1; i++) {
            this.framePool.add(new Frame(bufferCap));
        }
    }

    public Frame take() throws InterruptedException {
        return taskQueue.take();
    }

    public void put(Frame frame) throws InterruptedException {
        taskQueue.put(frame);
    }

    public Frame alloc() {
        int a = taskQueue.size();
        return framePool.get(taskQueue.size());
    }
}

class Frame {
    private byte[] frame;
    private int len;
    private long ts;
    // ...

    public Frame(int maxCap) {
        this.frame = new byte[maxCap];
    }

    public byte[] getFrameBuffer() {
        return this.frame;
    }

    public Frame setLen(int len) {
        this.len = len;
        return this;
    }

    public int getLen() {
        return len;
    }
}