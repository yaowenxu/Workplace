/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2020-03-31 22:19:28
 * LastEditTime: 2020-03-31 22:38:16
 * Description: C++ 数据结构链表的使用 std::list 
 * 编译运行选项: c++ -std=c++11 链表.cc ;./a.out 
 */

#include <iostream>
#include <list> // c++ 中的双向链表；
#include <forward_list> // c++ 中的单向链表；当不需要双向链表的时候，使用forward_list来实现，更为高效；
using namespace std;

// 链表可以常数时间删除和增加一个元素；
int main(){
    std::list<int> l = {7, 5, 16, 8};
    l.push_front(1);
    l.push_back(111);
    // 使用迭代器，打印出list中的元素
    for (int n : l){
        cout << n << endl;
    }
    // 使用迭代器，找到16号元素的位置，然后把在16号元素前插入一个元素；
    auto it = std::find(l.begin(), l.end(), 16);
    if (it != l.end())
    {
        l.insert(it, 122);
    }
    cout << endl;
    // 使用迭代器，打印出list中的元素
    for (int n : l){
        cout << n << endl;
    }
    return 0;
}
