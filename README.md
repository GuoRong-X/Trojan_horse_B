# Blake_generation
### 项目功能
执行木马后弹出word文档或其他文件，支持弹出网页url  
持续更新中。。。

### 项目思路
![](/images/sl.jpg)
### 使用方法
填写对应的参数  
![](/images/main.jpg)
    
    执行 go build -ldflags "-s -w -H=windowsgui"进行编译即可
    
### 免杀代码
免杀代码使用的k3rwin大神的https://github.com/k3rwin/shellcode-bypass-go
可以修改成自己的免杀
![](/images/免杀代码.jpg)
### golong exe加图标
在线转换链接:  
首推：https://convertio.co/zh/jpg-ico/  
次推：https://onlineconvertfree.com/zh/icon-generator/  
转换之后放到项目根目录，例如名为：logo.ico  重新编译即可
