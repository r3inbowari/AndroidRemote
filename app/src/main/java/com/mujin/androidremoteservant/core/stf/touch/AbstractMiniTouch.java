package com.mujin.androidremoteservant.core.stf.touch;


import com.mujin.androidremoteservant.core.stf.AbstractUDS;

public abstract class AbstractMiniTouch extends AbstractUDS {

    protected ThreadReceiver threadReceiver = null;
    protected ThreadSender threadSender = null;

    public AbstractMiniTouch(String local) {
        super(local);
    }

    public abstract ThreadReceiver getLocalReceiver();

    public abstract ThreadSender getLocalSender();

    public abstract void touch(int contact, int x, int y, int pressure);

    public abstract void swipe(int contact, int x, int y, int duration);

    /**
     * ndk 资源回收
     */
    public abstract void destroy();

    /**
     * minitouch run on thread ?
     */
    public abstract void run();
}
