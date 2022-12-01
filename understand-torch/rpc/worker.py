import os
import torch
import torch.distributed.rpc as rpc

os.environ['MASTER_ADDR'] = '127.0.0.1'
os.environ['MASTER_PORT'] = '7030'

# FIXME: 函数的定义必须在master和worker中都有，类似存根的概念;
# TODO: 可以不同，但是函数名字要相同；
def hello():
    return

# FIXME: rank数量还没达到 world size，程序会block;
rpc.init_rpc("worker", rank=1, world_size=2)

ret = rpc.rpc_async("master", hello, args=())
print("server:", ret.wait())

rpc.shutdown()