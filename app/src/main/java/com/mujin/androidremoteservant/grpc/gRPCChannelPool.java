package com.mujin.androidremoteservant.grpc;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;

public class gRPCChannelPool {

    // 池单例
    private static volatile gRPCChannelPool sInstance = null;

    /**
     * Channel 池
     *
     * @param key 连接名
     * @param value 连接通道
     */
    private ConcurrentHashMap<Object, ManagedChannel> maps;

    /**
     * 池构造
     */
    private gRPCChannelPool() {
        maps = new ConcurrentHashMap<>();
    }

    // 地址
    private String mHost;
    // 端口
    private int mPort;

    /**
     * 单例获取
     *
     * @return gRPCChannelPool 返回池
     */
    public static gRPCChannelPool get() {
        if (sInstance == null) {
            synchronized (gRPCChannelPool.class) {
                if (sInstance == null) {
                    sInstance = new gRPCChannelPool();
                }
            }
        }
        return sInstance;
    }

    /**
     * 初始化
     *
     * @param host 服务器地址
     * @param port 服务器端口
     */
    public static void init(String host, int port) {
        get().mHost = host;
        get().mPort = port;
    }

    /**
     * 获取 channel 附带生成效果
     *
     * @param name channel 名称
     * @param host 服务地址
     * @param port 服务端口
     * @return ManagedChannel channel
     */
    public ManagedChannel getChannel(Object name, String host, int port) {
        ManagedChannel channel = maps.get(name);
        if (channel == null || channel.isShutdown()) {
            channel = gRPCChannelUtils.newChannel(host, port);
            maps.put(name, channel);
        }
        return channel;
    }

    /**
     * 获取 channel
     *
     * @param name
     * @return
     */
    public ManagedChannel get(Object name) {
        return maps.get(name);
    }

    /**
     * 获取 channel
     *
     * @param name channel 名称
     * @return
     */
    public ManagedChannel getChannel(Object name) {
        return getChannel(name, mHost, mPort);
    }

    /**
     * 添加 channel
     *
     * @param name    channel 名称
     * @param channel channel
     * @return
     */
    public boolean addChannel(Object name, ManagedChannel channel) {
        ManagedChannel tmp = maps.get(name);
        if (tmp == null || tmp.isShutdown()) {
            tmp = channel;
            maps.put(name, tmp);
            return true;
        }
        return false;
    }

    /**
     * 添加 channel
     *
     * @param name channel 名称
     * @param host ...
     * @param port ...
     * @return
     */
    public boolean addChannel(Object name, String host, int port) {
        return addChannel(name, gRPCChannelUtils.newChannel(host, port));
    }

    /**
     * 添加 channel
     *
     * @param name channel 名称
     * @return
     */
    public boolean addChannel(Object name) {
        return addChannel(name, mHost, mPort);
    }

    /**
     * channel 移除
     *
     * @param name
     * @return
     */
    public ManagedChannel remove(Object name) {
        return maps.remove(name);
    }

    /**
     * channel 添加 不附带生成效果 针对下面三条 func
     *
     * @param name
     * @param channel
     * @return
     */
    public ManagedChannel add(Object name, ManagedChannel channel) {
        return maps.put(name, channel);
    }

    public ManagedChannel add(Object name, String host, int port) {
        return maps.put(name, gRPCChannelUtils.newChannel(host, port));
    }

    public ManagedChannel add(Object name) {
        return maps.put(name, gRPCChannelUtils.newChannel(mHost, mPort));
    }

    /**
     * 通道回收
     *
     * @param name
     * @return
     */
    public boolean shutdown(Object name) {
        ManagedChannel channel = maps.get(name);
        if (channel != null) {
            try {
                return channel.shutdown()
                        .awaitTermination(1, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            } finally {
                maps.remove(name);
            }
        }
        return false;
    }
}
