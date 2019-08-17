/**
 * Author: Yaowen Xu
 * Github: https://github.com/yaowenxu
 * Organization: 北航系统结构研究所
 * Date: 2019-08-17 16:18:34
 * LastEditTime: 2019-08-17 18:41:52
 * Description: 使用链表实现一个简单的cache,使用lru淘汰策略。没有使用容错，不支持并发；
 */
#include <unordered_map>
#include <cstdlib>
#include <cstring>
#include <iostream>
using namespace std;

struct  Node
{
    int _key;
    int _value;
    Node* _next;
    Node(int key, int value, Node* next):_key(key),_value(value),_next(next){}
};

class LRUCache
{
private:
    int _size;
    int _capacity;
    int _last;
    Node* _head;
    unordered_map <int, Node *> umap_prenodes;
    char* _begin; // 缓存的起始位置
    char* _cur_begin; // 用于分配节点内存的起始位置
public:
    LRUCache(int capacity);
    ~LRUCache();
    int get(int key) {
        int value             = -1;//初始时假设key对应的结点不存在
        Node* pre_node_of_key = umap_prenodes[key];//key对应的结点的前驱结点

        if(pre_node_of_key !=NULL){//key结点存在

            Node* node             = pre_node_of_key->_next;//key对应的结点
            pre_node_of_key->_next = node->_next;
            if(pre_node_of_key->_next!=NULL){
                umap_prenodes[pre_node_of_key->_next->_key] = pre_node_of_key;
            }

            node->_next            = _head->_next;
            if(node->_next!=NULL){//node有后继，更新后继的前驱结点
                umap_prenodes[node->_next->_key] = node;
            }

            _head->_next           = node;
            umap_prenodes[key]     = _head;              

            /*更新_last*/
            if(_last == key ){
                _last = ( pre_node_of_key == _head ? key : pre_node_of_key->_key ); 
            }

            value = node->_value;
        }
        return value;
    }

    void set(int key, int value) {
        Node* node            = NULL;
        Node* pre_node_of_key = umap_prenodes[key];//key对应的结点的前驱结点

        if(pre_node_of_key != NULL){//key对应的结点存在，孤立key对应的结点，也就是从链表中把结点取出来，重新链接链表

            node                   = pre_node_of_key->_next;//key对应的结点
            pre_node_of_key->_next = node->_next;

            if(pre_node_of_key->_next!=NULL){
                umap_prenodes[pre_node_of_key->_next->_key] = pre_node_of_key;//更新前驱
            }

            node->_value           = value; //重置结点值

            /*更新_last*/
            if(_last == key ){
                _last = ( pre_node_of_key == _head ? key : pre_node_of_key->_key ); 
            }
        }else{//结点不存在

            if(_capacity == 0){//缓冲区为空
                return ;
            }

            if(_size == _capacity){//缓存满，重用最后一个结点

                Node* pre_node_of_last    = umap_prenodes[_last];//最后一个结点的前驱结点

                umap_prenodes[pre_node_of_last->_next->_key] = NULL;
                
                node                      = new (pre_node_of_last->_next) Node(key,value,NULL);//重用最后一个结点

                pre_node_of_last->_next   = NULL;//移出最后一个结点

                _last = ( pre_node_of_last == _head ? key : pre_node_of_last->_key ); //更新指向最后一个结点的key

            }else{//缓冲未满，使用新结点

                node    = new (_cur_begin) Node(key,value,NULL);
                _cur_begin += sizeof(Node);
                _size++;
                if(_size==1){
                    _last = key;
                }
            }
        }
        /*把node插入到第一个结点的位置*/
        node->_next            = _head->_next;
        if(node->_next!=NULL){//node有后继，更新后继的前驱结点
            umap_prenodes[node->_next->_key] = node;
        }
        _head->_next           = node;
        umap_prenodes[key]     = _head;  

    }
};

LRUCache::LRUCache(int capacity)
{
    _capacity = capacity;
    _size = 0;
    _last = 0; // last 是链表中最后一个结点的key
    _cur_begin = _begin = (char*) malloc(sizeof(Node)*(capacity+1));
    _head = new (_cur_begin) Node(0, 0, NULL); // 在指定内存上构造对象
    _cur_begin += sizeof(Node);
}

LRUCache::~LRUCache()
{
    if(_begin != NULL){
        while (_cur_begin > _begin)
        {
            _cur_begin -= sizeof(Node);
            ((Node*) _cur_begin) -> ~Node(); // 进行释放内存上的对象
        }
        free(_begin);// 进行释放内存
    }
}

int main(int argc, char* argv[]){
    LRUCache cache(2);
    cache.set(1, 1000);
    cache.set(2, 2000);
    cache.set(3, 3000);
    cache.set(4, 4000);
    cout << cache.get(1) << endl;
    cache.set(5, 5000);
    cout << cache.get(4) << endl;
    cout << cache.get(5) << endl;
    return 0;
}