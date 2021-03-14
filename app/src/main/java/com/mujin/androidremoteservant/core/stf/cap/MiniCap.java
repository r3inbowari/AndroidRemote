package com.mujin.androidremoteservant.core.stf.cap;

import android.text.TextUtils;
import android.util.Log;

import com.google.protobuf.ByteString;
import com.mujin.androidremoteservant.core.utils.ScreenMetrics;
import com.mujin.androidremoteservant.core.utils.Utils;
import com.mujin.androidremoteservant.grpc.SimpleStreamObserver;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.HelloGrpc;
import com.r3inb.pb.HelloReply;
import com.r3inb.pb.JPEGGrpc;
import com.r3inb.pb.JPEGRequest;
import com.r3inb.pb.Pics;
import com.r3inb.pb.Reply;

import java.io.DataOutputStream;
import java.sql.Time;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.TimeUnit;

public class MiniCap extends AbstractMiniCap {

    private static MiniCap sInstance = null;

    private final BlockingQueue<byte[]> taskQueue;
    JPEGGrpc.JPEGStub jpegStub;


    MiniCap() {
        super("minicap");
        taskQueue = new ArrayBlockingQueue<>(100);
        jpegStub = JPEGGrpc.newStub(gRPCChannelPool.get().getChannel("jpeg"));
    }

    public void startSender() {
        System.out.println("-------------------------------------------- Sender Probe Start --------------------------------------------");
        new Thread(() -> {
            int count = 0;

            try {
                while (true) {
                    byte[] bytes = taskQueue.take();
                    count++;

                    JPEGRequest jpegRequest = JPEGRequest.newBuilder().setData(ByteString.copyFrom(bytes)).build();
                    jpegStub.sendJPEG(jpegRequest, new SimpleStreamObserver<Reply>() {
                        @Override
                        protected void onSuccess(Reply value) {
                            // System.out.println("OK");
                        }
                    });

                    if (count == 30) {
                        count = 0;
                        System.out.println("30 FP By Sender Probe");
                    }
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        }).start();
    }

    @Override
    public ThreadReceiver getLocalReceiver() {
        if (threadReceiver == null)
            this.threadReceiver = new ThreadReceiver(this, taskQueue);
        return threadReceiver;
    }

    public static MiniCap getInstance() {
        if (sInstance == null) {
            synchronized (MiniCap.class) {
                if (sInstance == null) {
                    sInstance = new MiniCap();
                }
            }
        }
        return sInstance;
    }

    @Override
    public void run() {
        MiniCapNDK rt = new MiniCapNDK(
                ScreenMetrics.getDeviceScreenWidth(),
                ScreenMetrics.getDeviceScreenHeight()
        );
        new Thread(
                rt,
                "MiniCap NDK"
        ).start();
    }
}

class MiniCapNDK extends Thread {

    static final String TAG = "Binary Cap";
    static final String ENV_PATH = "LD_LIBRARY_PATH=/data/local/tmp";
    static final String execName = "minicap";

    static int mx;
    static int my;
    static String execString;
    static String logStr = "> " + execName + ".log 2>&1";

    MiniCapNDK(int width, int height) {
        mx = width;
        my = height;
        if (mx > 5000 || my > 5000) {
            Log.e(TAG, "run failed used a error x or y metrics pixel");
        }
        String dp = "-P " + mx + "x" + my + "@" + mx + "x" + my + "/0 -Q 50";

        execString = TextUtils.join(" ",
                new Object[]{
                        ENV_PATH,
                        Utils.DIR_TEMP + "/" + execName,
                        dp,
                        logStr
                });
    }

    // LD_LIBRARY_PATH=/data/local/tmp /data/local/tmp/minicap -P 123x123@123x123/0 > minicap.log 2>&1
    @Override
    public void run() {
        Process process = null;
        DataOutputStream os = null;
        try {
            process = Runtime.getRuntime().exec("su");
            os = new DataOutputStream(process.getOutputStream());
            os.writeBytes("cd /data/local/tmp" + "\n");
            os.flush();
            os.writeBytes(execString + "\n");
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
        Log.d(TAG, "running mini cap on thread -> " + Thread.currentThread().getName());
    }
}

