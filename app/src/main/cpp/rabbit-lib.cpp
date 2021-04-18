//
// Created by r3inbowari on 2021/4/18.
//

#include "rabbit-lib.h"
#include <jni.h>
#include <string>
#include <unistd.h>
#include <fcntl.h>


extern "C"
JNIEXPORT jstring JNICALL
Java_com_mujin_androidremoteservant_MainActivity_stringFromJNI(
        JNIEnv *env,
        jobject /* this */) {
    std::string hello = "Hello from C++";
    int fd = open("/dev/input/event5", O_RDWR);
//    env->GetFieldID()
    printf("%d", fd);

    return env->NewStringUTF(hello.c_str());
}