package com.mujin.androidremoteservant.core.session;

import android.nfc.Tag;
import android.util.Log;

import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.ChatGrpc;
import com.r3inb.pb.ChatRequest;
import com.r3inb.pb.ChatResponse;

import io.grpc.stub.StreamObserver;

/**
 * 设备会话使用的chat通道类
 * 不继承到StreamObserver
 * 双向流
 */

public class Chat {
    // chat的存根
    private ChatGrpc.ChatStub chatStub = null;
    private StreamObserver<ChatRequest> result = null;

    private static Chat chatInstance = null;

    public static Chat getInstance() {
        if (chatInstance == null)
            chatInstance = new Chat();
        return chatInstance;
    }

    private final String TAG = "Chat";

    // connect
    // 连接和处理
    public void connectAndProcess() {
        chatStub = ChatGrpc.newStub(gRPCChannelPool.get().getChannel("chat"));
        StreamObserver<ChatResponse> requestStreamObserver = new StreamObserver<ChatResponse>() {

            @Override
            public void onNext(ChatResponse value) {
                // received data
                System.out.println(value.getOutput());
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

        // 观察者attach
        result = chatStub.bidStream(requestStreamObserver);
    }

    // 消息发送
    // @warning throw RuntimeException
    public boolean sendMsg(int type, String payload, String id) throws RuntimeException {
        try {
            ChatRequest request = ChatRequest.newBuilder()
                    .setInput(payload)
                    .setType(type)
                    .setId(id)
                    .build();
            result.onNext(request);
            return true;
        } catch (RuntimeException e) {
            // e.printStackTrace();
            Log.i("SessionChat", "Send failed..");
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
        Chat chat = new Chat();
        chat.connectAndProcess();
        chat.sendMsg(ChatTypeEnum.NOR, "hello", "123");
        chat.disconnect();
        chat.sendMsg(ChatTypeEnum.NOR, "can not", "123");
        chat.connectAndProcess();
        chat.sendMsg(ChatTypeEnum.NOR, "hello", "123");
    }

    /**
     * 注册本chat
     * 当该chat被注册时，设备消息将会被添加到服务器中
     */
    public void reg() {

    }
}

