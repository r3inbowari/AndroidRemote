package com.mujin.androidremoteservant.core.stf.cap;

import android.text.TextUtils;
import android.util.Log;

import com.google.protobuf.ByteString;
import com.mujin.androidremoteservant.core.session.Jpeg;
import com.mujin.androidremoteservant.core.utils.ScreenMetrics;
import com.mujin.androidremoteservant.core.utils.Utils;
import com.r3inb.pb.JPEGRequest;

import java.io.DataOutputStream;

public class MiniCap extends AbstractMiniCap {

    private static MiniCap sInstance = null;

    // private final BlockingQueue<byte[]> taskQueue;
    // 帧缓冲区阻塞队列
    private final FrameAlloc frameAlloc;

    // private JPEGGrpc.JPEGStub jpegStub;

    private final String TAG = "MiniCap";

    // private StreamObserver<JPEGRequest> requestJPEG;

    // send runFlag (atomic)
    private boolean senderRunFlag = false;

    MiniCap() {
        super("minicap");
        // taskQueue = new ArrayBlockingQueue<>(100);
        this.frameAlloc = new FrameAlloc(100, 1024 * 256);

        // 初始化 RPC 存根
//        this.jpegStub = JPEGGrpc.newStub(gRPCChannelPool.get().getChannel("jpeg"));
//
//        StreamObserver<JPEGReply> streamObserver = new StreamObserver<JPEGReply>() {
//            @Override
//            public void onNext(JPEGReply value) {
//                Log.i(TAG, "jpeg sender closed");
//            }
//
//            @Override
//            public void onError(Throwable t) {
//                t.printStackTrace();
//            }
//
//            @Override
//            public void onCompleted() {
//                System.out.println("aa");
//            }
//        };
//
//        requestJPEG = jpegStub.sendJPEG(streamObserver);
    }

    private JPEGRequest jpegRequest;

    Thread senderThread = null;

    // pause sender thread
    public void pauseSend() {
        this.senderRunFlag = false;
        // don't pause the thread here, the thread will sleep on next loop process
    }

    // notify the sender thread
    public void openSend() {
        synchronized (senderThread) {
            if (!senderRunFlag) {
                this.senderRunFlag = true;
                senderThread.notify();
            }
        }
    }

    public void startSender() {
        System.out.println("------------------- Sender Probe Start -------------------");
        senderThread = new Thread(() -> {
            int count = 0;

            try {
                while (true) {
                    // sender thread wait
                    // sender thread will notified by user through the rpc network
                    if (!senderRunFlag) {
                        synchronized (senderThread) {
                            senderThread.wait();
                        }
                    }

                    Frame frame = frameAlloc.take();
                    // count++;

                    // byte[] a = frame.getFrameBuffer();
                    // this.jpegStub = JPEGGrpc.newStub(gRPCChannelPool.get().getChannel("jpeg"));
                    // 数据转发
//                    jpegRequest = JPEGRequest.newBuilder()
//                            .setData(ByteString.copyFrom(frame.getFrameBuffer(), 0, frame.getLen()))
//                            .build();
//                    requestJPEG.onNext(jpegRequest);
                    // 数据转发
                    Jpeg.getInstance().sendFrame(
                            ByteString.copyFrom(
                                    frame.getFrameBuffer(),
                                    0,
                                    frame.getLen()
                            )
                    );

//                    jpegStub.sendJPEG(jpegRequest, new SimpleStreamObserver<Reply>() {
//                        @Override
//                        protected void onSuccess(Reply value) {
//                            // System.out.println("OK");
//                        }
//                    });

//                    if (count == 30) {
//                        count = 0;
//                        System.out.println("30 FP By Sender Probe");
//                    }
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        });
        senderThread.start();
    }

    @Override
    public ThreadReceiver getLocalReceiver() {
        if (threadReceiver == null)
            this.threadReceiver = new ThreadReceiver(this, frameAlloc);
        // this.threadReceiver = new ThreadReceiver(this, taskQueue);
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

