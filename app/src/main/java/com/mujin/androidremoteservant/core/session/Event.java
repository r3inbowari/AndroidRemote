package com.mujin.androidremoteservant.core.session;

import android.app.Instrumentation;
import android.text.TextUtils;
import android.util.Log;
import android.view.KeyEvent;

import com.google.common.base.Strings;
import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.EventGrpc;
import com.r3inb.pb.EventRequest;
import com.r3inb.pb.EventResponse;

import java.io.IOException;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;

import io.grpc.stub.StreamObserver;

/**
 * 单向 REPLY 流
 * 即当 event 被 Attach 后 REPLY 流不可以完成, 知道通信结束
 * 当前一个 event 未被服务器 detach 时, 重新进行 attach 操作, 那么新的信道将会代替旧的
 */
public class Event {

    // event 的存根
    private EventGrpc.EventStub eventStub = null;

    private static Event eventInstance = null;

    private String deviceID = null;

    public static Event getInstance() {
        if (eventInstance == null)
            eventInstance = new Event();
        return eventInstance;
    }

    //  private Instrumentation inst = new Instrumentation();

    private final BlockingQueue<Integer> keyQueue;
    private final ProcessShell processShell;

    private Event() {
        this.keyQueue = new ArrayBlockingQueue<>(100);
        this.processShell = new ProcessShell(true);

        processShell.Home();

        new Thread(() -> {
            try {
                while (true) {
                    Integer i = keyQueue.take();
                    Log.i("EvtKBT", String.valueOf(i));
                    new Thread(() -> {
                        processShell.KeyCode(i);
                    }).start();
                    // ProcessShell.exec("input keyevent " + String.valueOf(i), true);
                }
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }).start();
    }

    public void sendAndProcess() {
        // event 事件流
        eventStub = EventGrpc.newStub(gRPCChannelPool.get().getChannel("event"));
        // 准备就绪 发送 attach event (REG) 请求
        EventRequest request1 = EventRequest.newBuilder().setId(deviceID).setType(ChatTypeEnum.REG).build();

        Runtime runtime = Runtime.getRuntime();

        // reply 流事件处理
        eventStub.eventReq(request1, new StreamObserver<EventResponse>() {
            @Override
            public void onNext(EventResponse value) {

                try {
                    if (value.getType() == EventResponse.EventType.RELEASE) {
                        MiniTouch.getInstance().getEventManager().release(value.getContact());
                    } else if (value.getType() == EventResponse.EventType.TAP) {
                        MiniTouch.getInstance().getEventManager().tap(value.getContact(), value.getX(), value.getY());
                    } else if (value.getType() == EventResponse.EventType.SWIPE) {
                        MiniTouch.getInstance().getEventManager().move(value.getContact(), value.getX(), value.getY());
                    } else if (value.getType() == EventResponse.EventType.KEYBOARD) {
                        Log.i("EvtKB", String.valueOf(value.getX()));
//                        String keyCommand = "sudo input keyevent " + String.valueOf(value.getX());
                        // Process proc = runtime.exec(keyCommand);
                        keyQueue.put(value.getX());
                    }
                } catch (InterruptedException e) {
                    Log.i("Evt", "Event Resp failed..");
                }
            }

            @Override
            public void onError(Throwable t) {
                // 错误处理
            }

            @Override
            public void onCompleted() {
                // 通信结束
                System.out.println("event end off");
            }
        });
    }

    // 断开连接 不需要处理请求测断开, 因为请求不是流, 而是单次数据

    public Event setDeviceID(String deviceID) {
        this.deviceID = deviceID;
        return this;
    }
}
