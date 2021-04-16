package com.mujin.androidremoteservant.core.session;

import android.util.Log;

import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.EventGrpc;
import com.r3inb.pb.EventRequest;
import com.r3inb.pb.EventResponse;

import io.grpc.stub.StreamObserver;

/**
 * 单向 REPLY 流
 * 即当 Touch被Attach后 REPLY流不可以完成,知道通信结束
 * 当前一个touch未被服务器detach时, 重新进行attach操作, 那么新的信道将会代替旧的
 */
public class Event {

    // touch 的存根
    private EventGrpc.EventStub touchStub = null;

    private static Event touchInstance = null;

    private String deviceID = null;

    public static Event getInstance() {
        if (touchInstance == null)
            touchInstance = new Event();
        return touchInstance;
    }

    public void sendAndProcess() {
        // touch事件流
        touchStub = EventGrpc.newStub(gRPCChannelPool.get().getChannel("touch"));
        // 准备就绪 发送 attach touch (REG) 请求
        EventRequest request1 = EventRequest.newBuilder().setId(deviceID).setType(ChatTypeEnum.REG).build();

        // reply 流的处理
        touchStub.eventReq(request1, new StreamObserver<EventResponse>() {
            @Override
            public void onNext(EventResponse value) {
                // touch event 事件处理
                try {
                    if (value.getType() == EventResponse.EventType.RELEASE) {
                        MiniTouch.getInstance().getEventManager().release(value.getContact());
                    } else if (value.getType() == EventResponse.EventType.TAP) {
                        MiniTouch.getInstance().getEventManager().tap(value.getContact(), value.getX(), value.getY());
                    } else if (value.getType() == EventResponse.EventType.SWIPE) {
                        MiniTouch.getInstance().getEventManager().move(value.getContact(), value.getX(), value.getY());
                    } else if (value.getType() == EventResponse.EventType.KEYBOARD) {
                        Log.i("test", "test");
                    }
                } catch (InterruptedException e) {
                    Log.i("SessionTouch", "Touch failed..");
                }
            }

            @Override
            public void onError(Throwable t) {
                // 错误处理
            }

            @Override
            public void onCompleted() {
                // 通信结束
                System.out.println("touch end off");
            }
        });
    }

    // 断开连接 不需要处理请求测断开, 因为请求不是流, 而是单次数据

    public Event setDeviceID(String deviceID) {
        this.deviceID = deviceID;
        return this;
    }
}
