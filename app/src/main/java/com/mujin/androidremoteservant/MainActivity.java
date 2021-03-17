package com.mujin.androidremoteservant;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.cap.MiniCap;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.grpc.SimpleStreamObserver;
import com.mujin.androidremoteservant.grpc.gRPCChannelPool;
import com.r3inb.pb.ChatGrpc;
import com.r3inb.pb.ChatRequest;
import com.r3inb.pb.ChatResponse;
import com.r3inb.pb.Control;
import com.r3inb.pb.HelloGrpc;
import com.r3inb.pb.HelloReply;
import com.r3inb.pb.HelloRequest;
import com.r3inb.pb.Reply;
import com.r3inb.pb.Request;

import java.io.IOException;

import io.grpc.stub.StreamObserver;

import static com.mujin.androidremoteservant.core.SystemInit.prepareRuntimeNDKEnv;
import static com.mujin.androidremoteservant.core.utils.Utils.rootGrant;

public class MainActivity extends AppCompatActivity {

    private ChatGrpc.ChatStub chatStub = null;

    @Override
    protected void onDestroy() {
        MiniTouch.getInstance().destroy();
        super.onDestroy();
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        gRPCChannelPool.init("192.168.5.67", 5005);

        rootGrant(getPackageCodePath());

        System.out.println(getPackageCodePath());

        prepareRuntimeNDKEnv(this);

        ProcessShell ps = new ProcessShell(true);
        ps.Swipe(100, 200, 400, 500);


//        MiniTouch.getInstance().run();
//        MiniCap.getInstance().run();
//
//        MiniCap.getInstance().startSender();
//
//        try {
//            // TMP ANR
//            Thread.sleep(4000);
//        } catch (InterruptedException e) {
//            e.printStackTrace();
//        }
//
//        try {
//
//            MiniTouch.getInstance().connect();
//            MiniTouch.getInstance().getLocalReceiver().start();
//            MiniTouch.getInstance().getLocalSender().start();
//
//            MiniCap.getInstance().connect();
//            MiniCap.getInstance().getLocalReceiver().start();
//
//        } catch (IOException e) {
//            e.printStackTrace();
//        }


//        helloGrpc = HelloGrpc.newStub(gRPCChannelPool.get().getChannel("hello"));
//        HelloRequest request = HelloRequest.newBuilder().setName("hello").build();
//        helloGrpc.sayHello(request, new SimpleStreamObserver<HelloReply>() {
//            @Override
//            protected void onSuccess(HelloReply value) {
//                System.out.println("OK");
//            }
//        });
        chatStub = ChatGrpc.newStub(gRPCChannelPool.get().getChannel("chat"));

        StreamObserver<ChatResponse> requestStreamObserver = new StreamObserver<ChatResponse>() {

            @Override
            public void onNext(ChatResponse value) {
                System.out.println("s");
            }

            @Override
            public void onError(Throwable t) {
                t.printStackTrace();
            }

            @Override
            public void onCompleted() {

            }
        };

        ChatRequest request = ChatRequest.newBuilder().setInput("结束对222话\n").build();
        StreamObserver<ChatRequest> result = chatStub.bidStream(requestStreamObserver);

        result.onNext(request);
        result.onCompleted();

        try {
            Thread.sleep(5000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

    }
}