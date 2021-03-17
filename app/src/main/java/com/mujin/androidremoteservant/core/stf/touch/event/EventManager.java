package com.mujin.androidremoteservant.core.stf.touch.event;

import android.text.TextUtils;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ConcurrentHashMap;

public class EventManager {

    private static final String EVENT_TAP = "d";
    private static final String EVENT_MOV = "m";
    private static final String EVENT_RELEASE = "u";
    private static final String EVENT_COMMIT = "c";
    private static final String EVENT_KNOWN = "k";
    private static final String END_LINE = "\n";

    private static final int PRESSURE = 0;
//    Coords coords = null;
//    int pressure = 0;
//    int event = -1;

    private BlockingQueue<String> tasQueue = null;

    public EventManager(BlockingQueue<String> taskQueue) {
        this.tasQueue = taskQueue;
        for (int i = 0; i < 10; i++) {
            isPressed.put(i, false);
        }
    }

    private Map<Integer, Boolean> isPressed = new ConcurrentHashMap<>();

    public synchronized void tap(int contact, int x, int y) throws InterruptedException {
        if (isPressed.get(contact)) {
            release(contact);
        }
        String order = TextUtils.join(" ", new Object[]{EVENT_TAP, contact, x, y, PRESSURE}) + END_LINE;
        isPressed.put(contact, true);
        this.tasQueue.put(order);
        commit();
    }

    public synchronized void move(int contact, int x, int y) throws InterruptedException {
        if (isPressed.get(contact)) {
            String order = TextUtils.join(" ", new Object[]{EVENT_MOV, contact, x, y, PRESSURE}) + END_LINE;
            this.tasQueue.put(order);
        } else {
            tap(contact, x, y);
            this.isPressed.put(contact, true);
        }
        this.commit();
    }

    public synchronized void commit() throws InterruptedException {
        String order = EVENT_COMMIT + END_LINE;
        this.tasQueue.put(order);
    }

    public synchronized void release(int contact) throws InterruptedException {
        isPressed.put(contact, false);
        String order = TextUtils.join(" ", new Object[]{EVENT_RELEASE, contact}) + END_LINE;
        this.tasQueue.put(order);
        commit();
    }
}
