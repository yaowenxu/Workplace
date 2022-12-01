import os
import torch
import torch.distributed.rpc as rpc

os.environ['MASTER_ADDR'] = '127.0.0.1'
os.environ['MASTER_PORT'] = '7030'

# FIXME: 函数的定义必须在master和worker中都有，类似存根的概念;
def hello():
    print("hello!")
    return "hi, shuang"

# FIXME: rank数量还没达到 world size，程序会block;
rpc.init_rpc("master", rank=0, world_size=2)

rpc.shutdown()
