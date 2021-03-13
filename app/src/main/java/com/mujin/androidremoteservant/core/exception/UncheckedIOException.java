package com.mujin.androidremoteservant.core.exception;

import java.io.IOException;

/**
 * @author Stardust
 * @create 2017.4.1
 */
public class UncheckedIOException extends RuntimeException {

    public UncheckedIOException(IOException cause) {
        super(cause);
    }

    @Override
    public synchronized IOException getCause() {
        return (IOException) super.getCause();
    }
}
