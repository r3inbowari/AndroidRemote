package com.mujin.androidremoteservant.core.session;

public class ChatTypeEnum {
    public static final int DEF = 0;
    public static final int REG = 1;
    public static final int HRB = 2;
    public static final int NOR = 3;
    public static final int END = 4;
    public static final int ASK = 5;
    public static final int SNY = 6;
    public static final int REQ = 7;
    public static final int OPR = 8;

    public static final String SUCCEED = "succeed";
    public static final String FAILED = "failed";

    public static final int REQ_START_SENDER = 1000;
    public static final int REQ_PAUSE_SENDER = 1001;
    public static final int REQ_SEND_JPEG = 1002;

    public static final int REQ_OPEN_APP = 1003;

    public static final int REQ_CLOSE_APP = 1004;

    public static final int ASK_START_SENDER = 2000;
    public static final int ASK_PAUSE_SENDER = 2001;
    public static final int ASK_SEND_JPEG = 2002;

    public static final int ASK_OPEN_APP = 2003;

    public static final int ASK_CLOSE_APP = 2004;
}