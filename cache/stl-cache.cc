/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-17 18:42:10
 * LastEditTime: 2019-08-17 19:05:07
 * Description: 文件描述信息
 */

#include <iostream>
#include <unordered_map>
#include <list>
#include <cstdlib>
#include <utility>
using namespace std;

class LRUCache{
    size_t m_capacity;
    unordered_map<int,  list<pair<int, int> >::iterator> m_map; //m_map_iter->first: key, m_map_iter->second: list iterator;
    list< pair<int, int> > m_list;                               //m_list_iter->first: key, m_list_iter->second: value;
public:
    LRUCache(size_t capacity):m_capacity(capacity) {
    }
    int get(int key) {
        auto found_iter = m_map.find(key);
        if (found_iter == m_map.end()) //key doesn't exist
            return -1;
        m_list.splice(m_list.begin(), m_list, found_iter->second); //move the node corresponding to key to front
        return found_iter->second->second;                         //return value of the node
    }
    void set(int key, int value) {
        auto found_iter = m_map.find(key);
        if (found_iter != m_map.end()) //key exists
        {
            m_list.splice(m_list.begin(), m_list, found_iter->second); //move the node corresponding to key to front
            found_iter->second->second = value;                        //update value of the node
            return;
        }
        if (m_map.size() == m_capacity) //reached capacity
        {
           int key_to_del = m_list.back().first; 
           m_list.pop_back();            //remove node in list;
           m_map.erase(key_to_del);      //remove key in map
        }
        //m_list.emplace_front(key, value);  //create new node in list
        pair<int, int> a(key, value);
        m_list.push_front(a);
        m_map[key] = m_list.begin();       //create correspondence between key and node
    }
}; // use stl to imple the lru cache 

int main(int argc, char* argv[]){
    LRUCache cache(2);
    cache.set(1, 1000);
    cache.set(2, 2000);
    cache.set(3, 3000);
    cache.set(4, 4000);
    cout << cache.get(3) << endl;
    cache.set(5, 5000);
    cout << cache.get(4) << endl;
    cout << cache.get(5) << endl;
    return 0;
}