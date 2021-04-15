package com.mujin.androidremoteservant.core.stf;

import android.net.LocalSocket;
import android.net.LocalSocketAddress;

import java.io.IOException;

public abstract class AbstractUDS {

    public String local = null;
    public LocalSocket localSocket = null;

    public boolean status = false;

//    public AbstractUDS() {
//        this("minitouch");
//    }

    public AbstractUDS(String local) {
        this.local = local;
        this.localSocket = new LocalSocket();
    }

//    public void isOnlyOne() {
//        if (OnlyOnce) {
//            return ;
//        }
//    }

    public void connect() throws IOException {
        LocalSocketAddress localSocketAddress = new LocalSocketAddress(local);
        this.localSocket.connect(localSocketAddress);
        this.status = true;
    }

    /**
     * @throws IOException
     * @deprecated can't not close local socket safely due to the native impl
     */
    public void disconnect() throws IOException {
        if (localSocket.isConnected()) {
            localSocket.shutdownInput();
            localSocket.shutdownOutput();
            localSocket.close();
            boolean a = localSocket.isConnected();
            System.out.println("adsssssssssssssssssssssssssssssssssssssssssssssssssssss" + a);
        }
    }

    /**
     * @throws IOException
     * @deprecated can't not close local socket safely due to the native impl
     */
    public void reconnect() throws IOException, InterruptedException {

        // destroy and release
        // threadReceiver.close();
        // threadReceiver = null;
        disconnect();
        this.localSocket = new LocalSocket();
        connect();
    }

    /**
     * @throws IOException
     * @deprecated can't not close local socket safely due to the native impl
     */
    public void reconnect(String local) throws IOException, InterruptedException {

        // destroy and release
        // threadReceiver.close();
        // threadReceiver = null;
        disconnect();
        this.local = local;
        this.localSocket = new LocalSocket();
        connect();

    }

//    public static TouchManager getInstance() {
//        if (sInstance == null) {
//            synchronized (TouchManager.class) {
//                if (sInstance == null) {
//                    sInstance = new TouchManager();
//                }
//            }
//        }
//        return sInstance;
//    }

    public abstract void executeTouch(String touchCommand);

    public static AbstractUDS getInstance() {
        return null;
    }


}
