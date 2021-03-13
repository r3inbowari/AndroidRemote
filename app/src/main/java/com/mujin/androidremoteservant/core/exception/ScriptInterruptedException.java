package com.mujin.androidremoteservant.core.exception;

/**
 * @author Stardust
 * @create 2017.4.1
 */
public class ScriptInterruptedException extends ScriptException {

    public ScriptInterruptedException() {

    }

    public ScriptInterruptedException(Throwable e) {
        super(e);
    }

    public static boolean causedByInterrupted(Throwable e) {
        while (e != null) {
            if (e instanceof ScriptInterruptedException || e instanceof InterruptedException) {
                return true;
            }
            e = e.getCause();
        }
        return false;
    }

}