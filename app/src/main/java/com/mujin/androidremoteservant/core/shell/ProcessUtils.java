package com.mujin.androidremoteservant.core.shell;

import java.lang.reflect.Field;

/**
 * @author Stardust
 * @create 2017.4.1
 */
public class ProcessUtils {

    private static final String LOG_TAG = "ProcessUtils";

    public static int getProcessPid(Process process) {
        try {
            Field pid = process.getClass().getDeclaredField("pid");
            pid.setAccessible(true);
            return (int) pid.get(process);
        } catch (Exception e) {
            e.printStackTrace();
            return -1;
        }
    }
}