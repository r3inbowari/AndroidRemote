package com.mujin.androidremoteservant.grpc;

import java.io.InputStream;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.okhttp.OkHttpChannelBuilder;

public class gRPCChannelUtils {
    private gRPCChannelUtils() {
        throw new UnsupportedOperationException("u can't instantiate me...");
    }

    /**
     * 构建一条SSLChannel
     *
     * @param host         主机服务地址
     * @param port         端口
     * @param authority    域名
     * @param certificates 证书
     * @return
     */
    public static ManagedChannel newSSLChannel(String host, int port, String authority, InputStream... certificates) {
        HttpsUtils.SSLParams sslParams = HttpsUtils.getSslSocketFactory(certificates);
        return OkHttpChannelBuilder.forAddress(host, port)
                //overrideAuthority非常重要，必须设置调用
                .overrideAuthority(authority)
                .sslSocketFactory(sslParams.sSLSocketFactory)
                .build();
    }


    /**
     * 构建一条普通的Channel
     * TODO 在此处使用拦截器进行调试
     *
     * @param host 主机服务地址
     * @param port 端口
     * @return
     */
    public static ManagedChannel newChannel(String host, int port) {
        return ManagedChannelBuilder.forAddress(host, port)
                .usePlaintext()
                .build();
    }


    /**
     * 关闭Channel
     *
     * @param channel 端口
     * @return
     */
    public static boolean shutdown(ManagedChannel channel) {
        if (channel != null) {
            try {
                return channel.shutdown().awaitTermination(1, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        }
        return false;
    }
}
