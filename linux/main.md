1. lsof -i:8080      
根据端口查看对应的进程    

2. ps -ef|grep java ： 找出所有java进程     
    -a：显示所有用户的所有进程      
    -e：显示所有进程，环境变量f用树形格式来显示进程

3. find -name 文件名 ：找匹配的文件名

4. netstat  -anp  | grep   端口号
    查看某个端口是否被占用

5. tail -f 文件名
    查看文件倒数几行

6. tcpdump抓包
   
    tcpdump是一个用于截取网络分组，并输出分组内容的工具。
 
   1. tcpdump核心参数解析
      1. 