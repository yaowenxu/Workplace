/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-18 11:59:54
 * LastEditTime: 2019-08-18 12:45:45
 * Description: 使用 C 语言标准库函数clock来进行测试程序运行时间
 */
#include <time.h>
#include <stdio.h>
#include <math.h>

int str2int(char* str){
    char *p = str;
    int sum = 0;
    while (*p != '\0')
    {
        sum = sum*10 + (*p-'0');
        p++;
    }
    return sum;
}

int main(int argc, char* argv[]){
    clock_t start, stop;
    int def = 1000;
    if (argc == 2)
    {
        def = str2int(argv[argc-1]);
    }
    start = clock();
    for (int i = 0; i < def ; i++)
    {
        float tmp = sqrt(i);
    }
    stop =  clock();
    double total =  stop - start; // 使用运行的时间
    printf("Clocks: %.1f\n", total); // 总共使用的时钟
    printf("Time: %.1f us\n", total*1e6/(CLOCKS_PER_SEC)); // 转换运行时间为微秒
    return 0;
}

/*
 * Approximates the processor time used by the program, 
 * since the beginning of an implementation-defined time 
 * period that is related to the program invocation. 
 * To measure the time spent in a program, call the clock() 
 * function at the start of the program, and subtract its 
 * returned value from the value returned by subsequent calls 
 * to clock(). Then, to obtain the time in seconds, divide 
 * the value returned by clock() by CLOCKS_PER_SEC.
 * 
 * If you use the system() function in your program, do not 
 * rely on clock() for program timing, because calls to system() 
 * may reset the clock.
 * 
 * In a multithread POSIX C application, if you are creating threads 
 * with a function that is based on a POSIX.4a draft standard, 
 * the clock() function is thread-scoped.
 * 
 * Refer：https://www.ibm.com/support/knowledgecenter/en/SSLTBW_2.3.0/com.ibm.zos.v2r3.bpxbd00/clock.htm
*/