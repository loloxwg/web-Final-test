//package main //素数问题
//
//import (
//	"fmt"
//)
//
////seq 用以判断的数据源通道，通道首值会作为判断基数，其后数值作为判断值，向下一级别seq通道写值
////wait chan struct{} 作为通知载体，针对空struct内存申请返回的地址都是固定的，避免内存滥用
////level 当前执行 Processor 的级别，以判断Processor执行次数
//
//func Processor(seq <-chan int,wait chan struct{} , level int)  {
//	go func() {
//		//获取seq通道的初始值和状态，当最后级调用 Processor函数时（ seq的数7 ），此时的状态即为false
//		prime,ok := <-seq
//		if !ok {
//			close(wait)
//			//关闭wait通道，主函数中调用wait通道,起到信号传递功能，否则执行过快不会全部输出
//			return
//			//返回，否则会继续向下执行 等同于 if (){ do a ; return } do b ==  if (){ do a }else{ do b}
//		}
//		fmt.Printf("[%d]:%d\n",level,prime)     //打印当前级别和判断基数
//		out := make(chan int)                //作为下一级别调用Processor函数时使用的seq通道
//		Processor(out,wait,level+1)            //增加level，并循环调用Processor函数
//
//		//从seq通道中取值
//		//level ==1 当prime ==2 num 会取到 3 4 5 6 7 8 9 10
//		//level ==2 当prime ==3 num 会取到 5 7 9 (所有被2整除的数都被过滤)
//		//level ==3 当prime ==5 num 会取到 7
//		//level ==4 当prime ==7 num 取不到值 不会进入下面for循环
//		//           导致 out 通道没有数，再次调用Processor时，传入的seq通道状态为false
//
//		for num := range seq{  //从seq通道中读取值，并获取不能整除的数，写入out通道
//			if num%prime !=0 {
//				out <-num
//			}
//		}
//		close(out)    //关闭out通道，防止死锁
//	}()
//}
//
//func main() {
//	origin,wait := make(chan int),make(chan  struct{}) //初始化第一个seq通道和等待信号通道
//	Processor(origin,wait,1)
//
//	for i := 2 ; i < 123456 ; i++{     //向初始通道写入 2345678910 的值
//		origin <- i
//	}
//	close(origin)      //关闭初始通道
//	<-wait               //停止等待信号
//}