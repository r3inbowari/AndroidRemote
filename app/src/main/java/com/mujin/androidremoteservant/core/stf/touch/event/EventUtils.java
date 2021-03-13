package com.mujin.androidremoteservant.core.stf.touch.event;

import android.text.TextUtils;

public class EventUtils {

    private static final String EVENT_TAP = "d";
    private static final String EVENT_MOV = "m";
    private static final String EVENT_RELEASE = "u";
    private static final String EVENT_COMMIT = "c";
    private static final String EVENT_KNOWN = "k";
    private static final String END_LINE = "\n";

//    Coords coords = null;
//    int pressure = 0;
//    int event = -1;

    public static String tap(int contact, int x, int y, int pressure) {
        return TextUtils.join(" ", new Object[]{EVENT_TAP, contact, x, y, pressure}) + END_LINE;
    }

    // only m can call
    public static String move(int contact, int x, int y, int pressure) {
        return TextUtils.join(" ", new Object[]{EVENT_MOV, contact, x, y, pressure}) + END_LINE;
    }

    // only c can call
    public static String commit() {
        return EVENT_COMMIT + END_LINE;
    }

    // only u can call
    public static String release(int contact) {
        return TextUtils.join(" ", new Object[]{EVENT_RELEASE, contact}) + END_LINE;
    }
}
