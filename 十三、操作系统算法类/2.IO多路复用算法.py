"""
2.IO多路复用算法伪代码
时间复杂度：O(??)
空间复杂度：O(??)
"""
# select
class Solution1:
    def UserThread(self, socket):
        select(socket)  #将socket添加到select监视中

        while True:
            sockets = select()  #获取激活的socket
            for s in sockets:
                if can_read(s): # 一旦socket可读，则从socket中读取内容到缓冲区，然后在从缓冲区中取数据
                    read(s, buffer) # 设置缓冲区的目的，非阻塞IO，读写的时候数据存入buffer不会被阻塞，
                    process(buffer)

    def UserThread1(self):
        read_events, write_events = select(read_fds, write_fds, timeout)
        for event in read_events:
            handler_read(event)
        for event in write_events:
            handler_write(event)
        handler_others()
        
 
# poll       
class Solution2:
    
    
# epoll
class Solution3: