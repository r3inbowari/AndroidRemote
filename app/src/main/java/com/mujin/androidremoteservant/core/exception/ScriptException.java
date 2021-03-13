package com.mujin.androidremoteservant.core.exception;

/**
 * @author Stardust
 * @create 2017.4.1
 */
public class ScriptException extends RuntimeException {


    public ScriptException(String message, Throwable cause) {
        super(message, cause);
    }

    public ScriptException(String message) {
        super(message);
    }

    public ScriptException() {
    }

    public ScriptException(Throwable cause) {
        super(cause);
    }
}