```shell script
path = rocksdb
url = https://gitee.com/mirrors/rocksdb.git
pre = apt-get install make g++ libz-dev libsnappy-dev
cmd = git submodule update --init
init = cd rocksdb && make static_lib
post = go run xxx/main.go -h # 进行运行测试
```