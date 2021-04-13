package com.mujin.androidremoteservant.core.session;

import android.util.Log;

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

    // connect
    // 连接和处理
    public void connectAndProcess() {
        chatStub = ChatGrpc.newStub(gRPCChannelPool.get().getChannel("chat"));
        StreamObserver<ChatResponse> requestStreamObserver = new StreamObserver<ChatResponse>() {

            @Override
            public void onNext(ChatResponse value) {
                // 接收到数据时
                System.out.println(value.getOutput());
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
    public boolean sendMsg(String payload) throws RuntimeException {
        try {
            ChatRequest request = ChatRequest.newBuilder().setInput(payload).build();
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
        chat.sendMsg("hello");
        chat.disconnect();
        chat.sendMsg("cant");
        chat.connectAndProcess();
        chat.sendMsg("hello2");
    }

    /**
     * 注册本chat
     * 当该chat被注册时，设备消息将会被添加到服务器中
     */
    public void reg() {

    }
}
