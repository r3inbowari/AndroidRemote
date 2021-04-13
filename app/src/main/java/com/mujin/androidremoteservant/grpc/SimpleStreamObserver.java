package com.mujin.androidremoteservant.grpc;

import android.os.Handler;
import android.os.Looper;
import android.util.Log;

import io.grpc.Status;
import io.grpc.StatusException;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;

import static com.mujin.androidremoteservant.grpc.gRPCUtils.LTAG;

//  模板
public abstract class SimpleStreamObserver<T> implements StreamObserver<T> {
    private static Handler handler = new Handler(Looper.getMainLooper());

    @Override
    public void onCompleted() {

    }

    @Override
    public void onNext(final T value) {
        // 返回主线程
        handler.post(new Runnable() {
            @Override
            public void run() {
                onSuccess(value);
            }
        });
    }

    @Override
    public void onError(Throwable t) {
        t.printStackTrace();
        if (t instanceof StatusRuntimeException) {
            Status status = ((StatusRuntimeException) t).getStatus();
            Log.e(LTAG, status.getDescription());
        } else if (t instanceof StatusException) {
            Status status = ((StatusException) t).getStatus();
            Log.e(LTAG, status.getDescription());
        } else {
            Log.e(LTAG, t.getMessage());
        }
    }

    /**
     * 成功回调
     *
     * @param value
     */
    protected abstract void onSuccess(T value);
}