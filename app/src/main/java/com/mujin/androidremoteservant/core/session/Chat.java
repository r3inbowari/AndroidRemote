package com.mujin.androidremoteservant.core.session;

import android.util.Log;

import androidx.annotation.VisibleForTesting;

import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.core.utils.app.AppUtils;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.ChatGrpc;
import com.r3inb.pb.ChatRequest;
import com.r3inb.pb.ChatRequestOrBuilder;
import com.r3inb.pb.ChatResponse;

import io.grpc.stub.StreamObserver;

import static com.mujin.androidremoteservant.core.session.ChatTypeEnum.ASK_OPEN_APP;
import static com.mujin.androidremoteservant.core.session.ChatTypeEnum.FAILED;
import static com.mujin.androidremoteservant.core.session.ChatTypeEnum.REQ_OPEN_APP;
import static com.mujin.androidremoteservant.core.session.ChatTypeEnum.SUCCEED;

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

    private int heartbeatCnt = 0;

    private String deviceID = "null"; // must do

    private AppUtils appUtils = null;

    // connect
    // 连接和处理
    public Chat connectAndProcess() {
        chatStub = ChatGrpc.newStub(gRPCChannelPool.get().getChannel("chat"));
        StreamObserver<ChatResponse> requestStreamObserver = new StreamObserver<ChatResponse>() {

            @Override
            public void onNext(ChatResponse value) {
                // received data
                // System.out.println(value.getOutput());
                int ass = value.getType();
                switch (value.getType()) {
                    case ChatTypeEnum.REQ_START_SENDER:
                        Log.i(TAG, "sender start method");
                        MiniCap.getInstance().openSend();
                        sendMsg(ChatTypeEnum.ASK_START_SENDER, "null"); // ask
                        break;
                    case ChatTypeEnum.REQ_PAUSE_SENDER:
                        Log.i(TAG, "sender pause method");
                        MiniCap.getInstance().pauseSend();
                        sendMsg(ChatTypeEnum.ASK_PAUSE_SENDER, "null"); // ask
                        break;
                    case ChatTypeEnum.ASK:
                        Log.i(TAG, "heartbeat ask");
                        heartbeatCnt++;
                        break;
                    case ChatTypeEnum.REG:
                        Log.i(TAG, "device register ask");
                        break;
                    case REQ_OPEN_APP:
                        boolean isSucceed = false;
                        Log.i(TAG, "launcher app: " + value.getOutput());
                        if (appUtils != null)
                            isSucceed = appUtils.launchApp(value.getOutput());
                        else
                            Log.i(TAG, "launcher app failed(appUtils is null) app: " + value.getOutput());
                        sendMsg(ASK_OPEN_APP, isSucceed ? SUCCEED : FAILED);
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
        result = chatStub.bidStream(requestStreamObserver);
        return this;
    }

    // 消息发送
    // @warning throw RuntimeException
    public boolean sendMsg(int type, String payload) throws RuntimeException {
        try {
            ChatRequest request = ChatRequest.newBuilder()
                    .setInput(payload)
                    .setType(type)
                    .setInput(payload)
                    .setId(this.deviceID)
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
    @VisibleForTesting
    public void howtouse() {
        Chat chat = new Chat();
        chat.connectAndProcess().setDeviceID("1231231");
        chat.sendMsg(ChatTypeEnum.NOR, "hello");
        chat.disconnect();
        chat.sendMsg(ChatTypeEnum.NOR, "can not");
        chat.connectAndProcess();
        chat.sendMsg(ChatTypeEnum.NOR, null);
    }

    /**
     * 注册本chat
     * 当该chat被注册时，设备消息将会被添加到服务器中
     */
    public void reg() {

    }

    public int getHeartbeatCnt() {
        return this.heartbeatCnt;
    }

    public Chat setDeviceID(String deviceID) {
        this.deviceID = deviceID;
        return this;
    }

    public Chat setAppUtils(AppUtils appUtils) {
        this.appUtils = appUtils;
        return this;
    }
}

