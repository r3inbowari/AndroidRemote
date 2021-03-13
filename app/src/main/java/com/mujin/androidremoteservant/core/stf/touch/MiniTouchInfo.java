package com.mujin.androidremoteservant.core.stf.touch;

public class MiniTouchInfo {

    public String PID = null;
    public String MAX_X = null;
    public String MAX_Y = null;
    public String MAX_CONTACT = null;
    public String VERSION = null;
    public String MAX_PRESSURE = null;

    MiniTouchInfo(String PID, String VERSION, String MAX_CONTACT, String MAX_X, String MAX_Y, String MAX_PRESSURE) {
        this.PID = PID;
        this.VERSION = VERSION;
        this.MAX_CONTACT = MAX_CONTACT;
        this.MAX_Y = MAX_Y;
        this.MAX_X = MAX_X;
        this.MAX_PRESSURE = MAX_PRESSURE;
    }

    public static MiniTouchInfo parse(String src) {
        String[] arrayInfo = src.split("\n");
        if (arrayInfo.length == 4) {
            String version = arrayInfo[0].split(" ")[1];
            String max_contact = arrayInfo[1].split(" ")[1];
            String max_x = arrayInfo[1].split(" ")[2];
            String max_y = arrayInfo[1].split(" ")[3];
            String max_pressure = arrayInfo[1].split(" ")[4];
            String pid = arrayInfo[2].split(" ")[1];
            return new MiniTouchInfo(pid, version, max_contact, max_x, max_y, max_pressure);
        }
        return null;
    }
}
