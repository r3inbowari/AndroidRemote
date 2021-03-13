package com.mujin.androidremoteservant.core.stf.cap;

import android.util.Log;

import com.mujin.androidremoteservant.core.utils.Utils;

import java.io.IOException;
import java.io.InputStream;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;

public class ThreadReceiver extends Thread {

    String Tag = Thread.currentThread().getName();
    MiniCap threadMaster = null;
    boolean Running = false;
    Thread thread = null;

    ThreadReceiver(MiniCap miniCap) {
        this.Running = true;
        this.threadMaster = miniCap;
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
            CapProbe.HeadScanner(mRace);

            byte[] data;
            int receiveLen = 0;
            // data = new byte[2048]; // 2 MB


            // 1. 获取第一帧的帧长度 分配4个字节的空间并进行读取操作
            data = new byte[4]; // 2 MB
            mRace.read(data);
            ByteBuffer byteBuffer = ByteBuffer.wrap(data).order(ByteOrder.LITTLE_ENDIAN);
            int frameLen = byteBuffer.getInt();
            Log.e("frameframe", Integer.toString(frameLen));

            int sum = 0;
            data = new byte[frameLen];
            int loopCount = 0;
            while (Running) {
                // receiveLen = localSocket.getReceiveBufferSize();

                // int a = mRace.read(data);
                // ByteBuffer byteBuffer = ByteBuffer.wrap(Utils.subBytes(data, 2, 4)).order(ByteOrder.LITTLE_ENDIAN);
                // Log.e("header", CapProbe.maxX + "12");

                // 2. 分配 frameLen 个长度进行读取则成功获得第一帧的部分内容, 完成一次帧获取
                loopCount++;

                sum += mRace.read(data);
                if (frameLen == sum) {
                    Log.e("frameframe", "ok!");
                    String kg = Utils.printHexBinary(data);
                    System.out.println(kg);

                    data = new byte[4]; // 2 MB
                    mRace.read(data);
                    ByteBuffer bb = ByteBuffer.wrap(data).order(ByteOrder.LITTLE_ENDIAN);
                    int sd = bb.getInt();
                    Log.e("frameframe", Integer.toString(frameLen));

                    break;
                }


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
