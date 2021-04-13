package com.mujin.androidremoteservant.core.utils;

import android.content.Context;
import android.database.Cursor;
import android.database.sqlite.SQLiteDatabase;

public class PID {

    // 检查并尝试结束 process
    public static void checkAndKillProcess(Context context, String killName) {
        int pid = PID.getPID(context, killName);
        if (pid != -1) {
            Utils.kill(pid);
        }
    }

    // 将程序与pid记录到db中
    public static void Attach(Context context, String name, Integer pid) {
        SQLiteDatabase db = AssetsDBHelper.getDB(context).getWritableDatabase();
        String sql = "replace into " + AssetsDBHelper.TABLE_PIDs + " (Id, PID) values ('" + name + "', " + pid + ")";
        db.execSQL(sql);
        db.close();
    }

    // 通过name获取pid
    public static int getPID(Context context, String name) {
        SQLiteDatabase db = AssetsDBHelper.getDB(context).getWritableDatabase();
        String sql = "select * from " + AssetsDBHelper.TABLE_PIDs;
        Cursor cursor = db.rawQuery(sql, null);
        if (cursor != null && cursor.getColumnCount() > 0) {
            while (cursor.moveToNext()) {
                if (cursor.getString(0).equals(name))
                    return cursor.getInt(1);
            }
            cursor.close();
        }
        db.close();
        return -1;
    }
}
