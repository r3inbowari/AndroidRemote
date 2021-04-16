package com.mujin.androidremoteservant.core.session;

import android.util.Log;

import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.TouchGrpc;
import com.r3inb.pb.TouchReply;
import com.r3inb.pb.TouchRequest;

import io.grpc.stub.StreamObserver;

/**
 * 单向 REPLY 流
 * 即当 Touch被Attach后 REPLY流不可以完成,知道通信结束
 * 当前一个touch未被服务器detach时, 重新进行attach操作, 那么新的信道将会代替旧的
 */
public class Touch {

    // touch 的存根
    private TouchGrpc.TouchStub touchStub = null;

    private static Touch touchInstance = null;

    private String deviceID = null;

    public static Touch getInstance() {
        if (touchInstance == null)
            touchInstance = new Touch();
        return touchInstance;
    }

    public void sendAndProcess() {
        // touch事件流
        touchStub = TouchGrpc.newStub(gRPCChannelPool.get().getChannel("touch"));
        // 准备就绪 发送 attach touch (REG) 请求
        TouchRequest request1 = TouchRequest.newBuilder().setId(deviceID).setType(ChatTypeEnum.REG).build();

        // reply 流的处理
        touchStub.touchReq(request1, new StreamObserver<TouchReply>() {
            @Override
            public void onNext(TouchReply value) {
                // touch event 事件处理
                try {
                    if (value.getType().equals(TouchReply.TouchType.RELEASE)) {
                        MiniTouch.getInstance().getEventManager().release(value.getContact());
                    } else if (value.getType().equals(TouchReply.TouchType.TAP)) {
                        MiniTouch.getInstance().getEventManager().tap(value.getContact(), value.getX(), value.getY());
                    } else if (value.getType().equals(TouchReply.TouchType.SWIPE)) {
                        MiniTouch.getInstance().getEventManager().move(value.getContact(), value.getX(), value.getY());
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


    public Touch setDeviceID(String deviceID) {
        this.deviceID = deviceID;
        return this;
    }
}
