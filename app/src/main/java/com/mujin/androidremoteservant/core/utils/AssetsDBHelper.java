package com.mujin.androidremoteservant.core.utils;

import android.content.Context;
import android.database.sqlite.SQLiteDatabase;
import android.database.sqlite.SQLiteOpenHelper;

public class AssetsDBHelper extends SQLiteOpenHelper {
    private static final int DB_VERSION = 3;
    private static final String DB_NAME = "local.db";
    public static final String TABLE_PIDs = "PIDs";

    private static AssetsDBHelper assetsDBHelper = null;

    public AssetsDBHelper(Context context) {
        super(context, DB_NAME, null, DB_VERSION);
    }

    @Override
    public void onCreate(SQLiteDatabase db) {
        String sql = "create table if not exists " + TABLE_PIDs + " (Id text primary key, PID integer)";
        db.execSQL(sql);
    }

    @Override
    public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion) {
        String sql = "DROP TABLE IF EXISTS " + TABLE_PIDs;
        db.execSQL(sql);
        onCreate(db);
    }

    public static AssetsDBHelper getDB(Context context) {
        if (assetsDBHelper == null) {
            assetsDBHelper = new AssetsDBHelper(context);
        }
        return assetsDBHelper;
    }

}
