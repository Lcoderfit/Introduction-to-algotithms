package main

/*
TCP粘包：指发送方发送的多个包粘在一起形成一个包（TCP是发送数据是流式发送的，不知道一个包的开头和结尾）
所以出现一个包的头部紧挨着另一个包的尾部发送

2.导致粘包的主要原因：tcp数据传递模式是流模式，保持长连接的时候可以进行多次收和发

2.1 发送端粘包：Nagle算法（改善网络传输效率的方法，在未收到确认数据已发送的ACK时，
发送器会将数据放到缓存里，直到收到ACK或达到小于MSS的最大值时才发送）

为什么Nagle能提高效率：因为如果一个数据包包括首部和实体，
例如如果我发送两个实体为10字节的数据包，由于每个数据包首部为40，
所以两个数据包首部耗费了80字节；
如果我把两个包合成一个，则发送同样的实体数据，

2.2 接收端粘包：接收端会把接收到的数据放入缓存，然后通知应用层取数据，如果
由于某些原因不能及时把数据取出来，则会造成TCP缓冲区中存放了几段数据


mspaint: 打开画图功能
*/

func main() {
}
