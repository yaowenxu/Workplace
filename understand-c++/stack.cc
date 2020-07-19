/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-27 11:12:07
 * LastEditTime: 2019-08-27 11:25:03
 * Description: c++ stack 的使用
 */

#include <stack>
#include <vector>
#include <iostream>
using namespace std;

int main(){
    std::stack<int, std::vector<int> > singlestack;
    singlestack.push(1);
    singlestack.push(2);
    singlestack.push(3);
    singlestack.push(4);
    cout << singlestack.top() << endl;
    singlestack.pop();
    cout << singlestack.top() << endl;
    singlestack.pop();
    cout << singlestack.top() << endl;
    singlestack.pop();
    cout << singlestack.top() << endl;
    singlestack.pop();
    return 0;
}