package com.mujin.androidremoteservant;

import androidx.annotation.RequiresApi;
import androidx.appcompat.app.AppCompatActivity;

import android.app.ActivityManager;
import android.content.Context;
import android.media.AudioFormat;
import android.media.AudioRecord;
import android.media.MediaRecorder;
import android.os.Build;
import android.os.Bundle;
import android.util.Log;
import android.view.Window;
import android.widget.Button;

import com.mujin.androidremoteservant.core.shell.ProcessShell;
import com.mujin.androidremoteservant.core.stf.touch.MiniTouch;
import com.mujin.androidremoteservant.core.utils.app.AppUtils;

import java.io.FileOutputStream;
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.SocketException;

import static com.mujin.androidremoteservant.core.SystemInit.applicationEntry;
import static com.mujin.androidremoteservant.core.SystemInit.prepareRuntimeNDKEnv;

// demo 仅供测试，不保活
public class MainActivity extends AppCompatActivity {
    static {
        System.loadLibrary("rabbit-lib");
    }

    public native String stringFromJNI();

    @Override
    protected void onDestroy() {
        MiniTouch.getInstance().destroy();
        super.onDestroy();
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        // may be start on a new thread (avoid ANR)...
        prepareRuntimeNDKEnv(this);
        new Thread(() -> {
            applicationEntry(this);
        }).start();

        try {
            initMR();
            recordAndPlay();
        } catch (SocketException e) {
            e.printStackTrace();
        }
    }

    AudioRecord mRecord = null;
    boolean mReqStop = false;

    private final int kSampleRate = 44100;
    private final int kChannelMode = AudioFormat.CHANNEL_IN_STEREO;
    private final int kEncodeFormat = AudioFormat.ENCODING_PCM_16BIT;
    private final int kFrameSize = 2048;
    private static String TAG = "JZJ12";

    DatagramSocket socket = null;

    private void initMR() throws SocketException {
        socket = new DatagramSocket(5004);

        int minBufferSize = AudioRecord.getMinBufferSize(kSampleRate, kChannelMode,
                kEncodeFormat);
        mRecord = new AudioRecord(MediaRecorder.AudioSource.MIC,
                kSampleRate, kChannelMode, kEncodeFormat, minBufferSize * 2);

    }
    int cnt = 0;
    private void recordAndPlay() {
        // FileOutputStream os = null;
        mRecord.startRecording();
        try {
            // os = new FileOutputStream(filePath);
            byte[] buffer = new byte[kFrameSize];
            int num = 0;
            while (!mReqStop) {
                num = mRecord.read(buffer, 0, kFrameSize);
                Log.d(TAG, "buffer = " + buffer.toString() + ", num = " + num);
                // os.write(buffer, 0, num);
                // write here
                // byte[] buf = new byte[1024];
                // 解析数据包
                DatagramPacket packet = new DatagramPacket(buffer, buffer.length);
                cnt ++;
                socket.send(packet);
                System.out.println(cnt);
            }

            Log.d(TAG, "exit loop");
            // os.close();
        } catch (IOException e) {
            e.printStackTrace();
            Log.e(TAG, "Dump PCM to file failed");
        }
        mRecord.stop();
        mRecord.release();
        mRecord = null;
        Log.d(TAG, "clean up");
    }

//    public void stop(View view) {
//        mReqStop = true;
//        Button stopBtn = (Button) findViewById(R.id.stopBtn);
//        stopBtn.setText("Stopped");
//        stopBtn.setEnabled(false);
//    }

}
