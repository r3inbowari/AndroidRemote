package com.mujin.androidremoteservant.core.utils;

import android.content.Context;
import android.text.TextUtils;
import android.util.Log;

import com.mujin.androidremoteservant.core.shell.ProcessShell;

import static com.mujin.androidremoteservant.core.utils.Utils.*;


public class AssetsOperation {


    public final static String TAG = "File Operation";

    public static void unzip(Context context, String assetsPath, String targetPath, String rename) {
        Log.i(TAG, "unzip -> " + assetsPath + " to " + targetPath);
        AssetsCopy.copy(context, assetsPath, targetPath, rename);
    }

    public static void changeMod(String level, String path) {
        Log.i(TAG, "chmod -> " + path);
        ProcessShell.exec(TextUtils.join(" ", new Object[]{ORDER_CHMOD, level, path}), true);
    }

    public static void changeOwn(String owner, String path) {
        Log.i(TAG, "chown -> " + path);
        ProcessShell.exec(TextUtils.join(" ", new Object[]{ORDER_CHOWN, owner, path}), true);
    }

    public static void changeGroup(String group, String path) {
        Log.i(TAG, "chgrp -> " + path);
        ProcessShell.exec(TextUtils.join(" ", new Object[]{ORDER_CHGRP, group, path}), true);
    }

    public static boolean clean(String path) {
//        if (!path.equals("/")) {
//            File file = new File(path);
//            if (file.exists()) {
//                Log.i(TAG, "clean -> " + path);
//                return file.delete();
//            }
//        }
//        return false;
        if (path.equals("/")) {
            return false;
        }
        ProcessShell.exec(TextUtils.join(" ", new Object[]{ORDER_DEL, path}), true);
        return true;
    }

    public static void autoZipBySU(Context context, String srcAssetsPath, String targetPath, String rename) {
        String realTargetPath = targetPath + "/" + rename;
        AssetsOperation.clean(realTargetPath);

        // 比较猥琐的两步操作
        AssetsOperation.unzip(context, srcAssetsPath, context.getCacheDir().getPath(), rename);
        ProcessShell.exec(TextUtils.join(" ", new Object[]{"mv", context.getCacheDir().getPath() + "/" + rename, targetPath}), true);

        // AssetsOperation.unzip(context, srcAssetsPath, targetPath, rename);
        AssetsOperation.changeMod(LEVEL_ALL, realTargetPath);
        AssetsOperation.changeOwn(OWNER_ROOT, realTargetPath);
        AssetsOperation.changeGroup(GROUP_ROOT, realTargetPath);
    }
}
