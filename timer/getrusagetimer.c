/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-18 13:31:34
 * LastEditTime: 2019-08-18 13:51:26
 * Description: 使用系统 函数 getrusage 获取程序运行相关信息 
 *              此程序主要是关注与总时间和内核空间运行时间与用户
 *              空间运行时间，使用此函数可大致对程序运行时间计算；
 * 查看: 具体使用信息可以在控制台以 man getrusage 命令查看
 */
#include <stdio.h>
#include <math.h>
#include <sys/time.h>
#include <sys/resource.h>

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
    int def = 1000;
    if (argc == 2)
    {
        def = str2int(argv[argc-1]);
    }
    for (int i = 0; i < def ; i++)
    {
        float tmp = sqrt(i);
    }
    struct rusage usage;
    getrusage(RUSAGE_SELF, &usage);
    //getrusage(RUSAGE_CHILDREN, &usage);
    long user =  usage.ru_utime.tv_sec * 1000000 + usage.ru_utime.tv_usec; // user time used
    long sys = usage.ru_stime.tv_sec * 1e6 + usage.ru_stime.tv_usec; // sys time used
    
    printf("User: %ld us\n", user); //  用户空间使用的时间
    printf("Sys: %ld us\n", sys); //  内核空间使用的时间
    printf("Total: %ld us\n", user+sys); // 总共使用的时钟
    return 0;
}