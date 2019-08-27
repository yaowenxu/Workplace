/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-27 16:42:24
 * LastEditTime: 2019-08-27 17:03:54
 * Description: 打印小于 1000 的所有素数
 */

#include <iostream>
using namespace std;


// 使用开方来进行提速；
// 使用整除来进行提速；
// 使用性能工具来进行提速；
bool prime(int n){
    int i;
    for (i=2;i<n;i++){
        if(n%i==0){
            return false;
        }
    }
    return true;
}

int main(){
    int i, n;
    n = 1000;
    for(i = 2; i <= n; i++){
        if(prime(i)){
            printf("%d\n", i);
        }
    }
    return 0;
}