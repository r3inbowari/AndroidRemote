package com.mujin.androidremoteservant.core.stf.cap;

import android.util.Log;

import java.io.IOException;
import java.io.InputStream;
import java.nio.ByteBuffer;
import java.nio.ByteOrder;

public class CapProbe {

    static final int HEADER_LEN = 24;
    static public int VERSION = 1;

    static final byte[] HEADER_DEFINE = {
            0, 1,  // Version (currently 1)
            1, 1,  // Size of the header (from byte 0)
            2, 4,  // Pid of the process
            6, 4,  // Real display width in pixels
            10, 4, // Real display height in pixels
            14, 4, // Virtual display width in pixels
            18, 4, // Virtual display height in pixels
            22, 1, // Display orientation
            23, 1  // Quirk bitflags (see below)
    };

    // public int size;
    public static int pid;
    public static int maxX;
    public static int maxY;
    public static int maxXV;
    public static int maxYV;
    public static int orientation;
    public static int quirk;

    public static byte[] raw = null;

    private static boolean headerProbe(byte[] data) {
        if (data.length == HEADER_LEN && data[1] == HEADER_LEN && data[0] == VERSION) {
            for (int i = 2; i < HEADER_DEFINE.length; i += 2) {
                ByteBuffer byteBuffer = ByteBuffer.wrap(data, HEADER_DEFINE[i], HEADER_DEFINE[i + 1]).order(ByteOrder.LITTLE_ENDIAN);

                switch (i) {
                    case 4:
                        pid = byteBuffer.getInt();
                        break;
                    case 6:
                        maxX = byteBuffer.getInt();
                        break;
                    case 8:
                        maxY = byteBuffer.getInt();
                        break;
                    case 10:
                        maxXV = byteBuffer.getInt();
                        break;
                    case 12:
                        maxYV = byteBuffer.getInt();
                        break;
                    case 14:
                        orientation = byteBuffer.get();
                        break;
                    case 16:
                        quirk = byteBuffer.get();
                        break;
                }
            }
            return true;
        }
        return false;
    }

    public static void HeadScanner(InputStream inputStream) throws IOException {
        raw = new byte[HEADER_LEN];
        int n = inputStream.read(raw);
        Log.i("Probe", "find header frame bit count -> " + n);
        headerProbe(raw);
    }
}
