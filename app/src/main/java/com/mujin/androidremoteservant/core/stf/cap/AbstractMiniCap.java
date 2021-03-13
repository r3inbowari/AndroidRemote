package com.mujin.androidremoteservant.core.stf.cap;

import com.mujin.androidremoteservant.core.stf.AbstractUDS;

public abstract class AbstractMiniCap extends AbstractUDS {

    protected ThreadReceiver threadReceiver = null;

    public AbstractMiniCap(String local) {
        super(local);
    }

    @Override
    public void executeTouch(String touchCommand) {

    }

    public abstract ThreadReceiver getLocalReceiver();

    public abstract void run();
}
