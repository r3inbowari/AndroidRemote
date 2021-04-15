package com.mujin.androidremoteservant.core.session;

import android.util.Log;

import com.google.protobuf.ByteString;
import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.JPEGGrpc;
import com.r3inb.pb.JPEGReply;
import com.r3inb.pb.JPEGRequest;

import io.grpc.stub.StreamObserver;

/**
 * 设备会话使用的 jpeg 通道类
 * 不继承到 StreamObserver
 * 双向流
 */

public class Jpeg {
    // jpeg的存根
    private JPEGGrpc.JPEGStub jpegStub = null;
    private StreamObserver<JPEGRequest> result = null;

    private static Jpeg jpegInstance = null;

    public static Jpeg getInstance() {
        if (jpegInstance == null)
            jpegInstance = new Jpeg();
        return jpegInstance;
    }

    private final String TAG = "Jpeg";

    // connect
    // 连接和处理
    public Jpeg connectAndProcess() {
        jpegStub = JPEGGrpc.newStub(gRPCChannelPool.get().getChannel("jpeg"));
        StreamObserver<JPEGReply> requestStreamObserver = new StreamObserver<JPEGReply>() {

            @Override
            public void onNext(JPEGReply value) {
                // received data
//                System.out.println(value.getOutput());
                switch (value.getType()) {
                    case ChatTypeEnum.REQ_START_SENDER:
                        Log.i(TAG, "sender start method");
                        MiniCap.getInstance().openSend();
                        break;
                    case ChatTypeEnum.REQ_PAUSE_SENDER:
                        Log.i(TAG, "sender pause method");
                        MiniCap.getInstance().pauseSend();
                        break;
                    default:
                        Log.i(TAG, "unsupported method");
                }
            }

            @Override
            public void onError(Throwable t) {
                // 处理错误
                t.printStackTrace();
            }

            @Override
            public void onCompleted() {
                // 过程结束
            }
        };

        // 观察者 attach
        result = jpegStub.sendJPEG(requestStreamObserver);
        return this;
    }

    private static String did = null;

    // 消息发送
    // @warning throw RuntimeException
    public boolean sendFrame(ByteString payload) throws RuntimeException {
        try {
            JPEGRequest request = JPEGRequest.newBuilder()
                    .setData(payload)
                    .setType(ChatTypeEnum.REQ_SEND_JPEG)
                    .setId(did)
                    .build();
            result.onNext(request);
            return true;
        } catch (RuntimeException e) {
            // e.printStackTrace();
            Log.i("SessionJpeg", "Send failed..");
            return false;
        }
    }

    public boolean regJpeg(String id) throws RuntimeException {
        did = id;
        try {
            JPEGRequest request = JPEGRequest.newBuilder()
                    .setType(ChatTypeEnum.REG)
                    .setId(id)
                    .build();
            result.onNext(request);
            return true;
        } catch (RuntimeException e) {
            e.printStackTrace();
            Log.i("SessionJpeg", "Send failed..");
            return false;
        }
    }

    // 断开连接
    public void disconnect() {
        result.onCompleted();
    }

    /**
     * @deprecated unit test
     */
    public void howtouse() {
        Jpeg jpeg = new Jpeg();
        jpeg.connectAndProcess();
//        jpeg.sendMsg(ChatTypeEnum.NOR, "hello", "123");
//        jpeg.disconnect();
//        jpeg.sendMsg(ChatTypeEnum.NOR, "can not", "123");
//        jpeg.connectAndProcess();
//        jpeg.sendMsg(ChatTypeEnum.NOR, "hello", "123");
    }

    /**
     * 注册本 jpeg
     * 当该 jpeg 被注册时，设备消息将会被添加到服务器中
     */
    public void reg() {

    }
}
