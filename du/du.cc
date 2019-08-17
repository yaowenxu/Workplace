#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <string>
#include <iostream>

/*
 * 函数名：du
 * 功能：根据文件夹路径，获取该路径对应的文件夹的总容量，并以KB为单位输出；
 * 输入：文件夹对应的全局路径；
 * 输出：改文件夹的大小，以KB为单位；
 */
int64_t du(std::string path){
    int cmdsize = 150; // 默认命令的总长度不超过150个字符
    int bufsize = 200;
    char cmd[cmdsize];
    char buf[bufsize];
    memset(cmd, '\0', cmdsize);
    memset(buf, '\0', bufsize);
    if(sprintf(cmd, "du -h -d 0 %s", path.c_str()) <0 ){
        perror("sprintf()");
        return -1;
    }
    FILE *fp = popen(cmd, "r");
    if (fp == NULL) {
        perror("popen()");
        return -1;
    }
    fgets(buf, bufsize-1, fp);
    std::string cmdresult(buf);
    std::cout << "du 运行结果：" << cmdresult;
    // 下面进行根据获取的返回值解析容量大小
    int formal_size = 0;
    size_t start = 0;
    start = cmdresult.find_first_of(" ", start);
    size_t end = cmdresult.find_first_of("\t", start+1);
    // 获取容量字符串
    std::string size_str = cmdresult.substr(start+1, end-start-1);
    // 获取容量数字大小
    int unit = size_str.size()-1;
    // std::cout  << size_str.substr(0, unit).c_str() << std::endl;
    sscanf(size_str.substr(0, unit).c_str(), "%d", &formal_size);
    std::cout << "解析后容量大小(数字):"<< formal_size << size_str[unit]<< std::endl;
    // 根据容量后单位G,M,K统一转化为KB为单位的容量大小
    int64_t size = 0;
    if(size_str[unit] == 'G'){
        size = formal_size * 1000 * 1000;
    }
    if(size_str[unit] == 'M'){
        size = formal_size*1000;
    }
    if(size_str[unit] == 'K'){
        size = formal_size;
    }
    // 输出字符串形式的容量大小
    std::cout << "解析后容量大小(字符串):" << size_str << std::endl;
    return size;
}

int main(int argc, char* argv[]){
    std::string path(argv[argc-1]); // 构造command命令
    std::string command(argv[0]);
    if(command == path){
        path = "."; // 如果当前命令没有输出参数，默认路径为当前路径 「.」
    }
    long long size = du(path); // 在此处调用函数，进行执行命令；
    std::cout << path << " 的大小：" << size << " KB" << std::endl;
    return 0;
}