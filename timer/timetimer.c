/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-18 13:03:53
 * LastEditTime: 2019-08-18 13:14:33
 * Description: 使用 C 语言库 time 函数 对程序运行计时 以秒为单位
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
    time_t start, stop; // time_t aka long 
    int def = 1000;
    if (argc == 2)
    {
        def = str2int(argv[argc-1]);
    }
    start = time(NULL);
    for (int i = 0; i < def ; i++)
    {
        float tmp = sqrt(i);
    }
    stop =  time(NULL);
    time_t total =  stop - start; // 使用运行的时间 以秒为单位
    printf("Start: %ld s\n", start);
    printf("Stop: %ld s\n", stop);
    printf("Time: %ld s\n", total); // 总共使用的时钟
    return 0;
}