/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-26 19:24:48
 * LastEditTime: 2019-08-26 19:39:43
 * Description: C++ array demo
 */
#include <iostream>
#include <array>
using namespace std;

int main(int argc, char* argv[]){
    std::array<int, 5> arr = {1,2,3,4,5};
    array<bool,5> arrbool;
    arrbool.fill(true);
    for (auto b : arrbool){
        //cout << b << endl;
    }
    cout << sizeof(bool) << endl; // 1byte
    cout << sizeof(int) << endl; // 4byte
    cout << sizeof(int64_t) << endl; // 8byte
    cout << sizeof(char) << endl; // 1byte
    cout << arrbool.size() << endl;
    cout << arrbool.data() << endl;
    cout << *(arrbool.data()) << endl;
    return 0;
}