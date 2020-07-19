/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2020-03-30 17:42:12
 * LastEditTime: 2020-03-30 17:49:06
 * Description: stl 优先队列测试
 */

#include <iostream>
#include <functional>
#include <vector>
#include <queue> //优先队列
using namespace std;
priority_queue<int, vector<int>,std::greater<int> > que; // 默认优先队列是个大顶堆；
// 返回当前结果
int main(){
    que.push(10);
    que.push(11);
    cout << que.top() << endl;
    return 0;
}