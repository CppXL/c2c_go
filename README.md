## Server基本功能
1. 维护Client连接列表
2. 下发命令(单个下发/批量下发)

    * 执行系统命令
    * 更新Client
    * 获取Client系统信息
    
        * 进程列表
        * IP地址
        * 系统版本
        
    * 扫描同网段其他机器
    * DDoS
    * 下载文件


3. 保活(呼吸功能)
4. 根据配置文件载入配置


#### 指令列表

* cmd command args
* download url save_path
* update
* getinfo 
* dos atk_method ipaddr/mask
* 


#### 指令格式



参考

| https://blog.netlab.360.com/pinkbot/
```
    Token字段，长度4字节，该字段值由服务器端指定，指定后将一直使用这个值。设置方式为：Bot启动后首先会向CC发送新生成的ECDH的公钥，此刻Token为0，当服务端接受后，会分配一个Token值给Bot，这就算指定成功了。
    指令字段，长度1字节。CC发出指令后，Bot也要用相同的指令码把执行结果返回。
    内容长度字段，长度2字节。当指令不包含具体内容时，设置为零，否则这里填充内容的字节长度数，并追加密文内容。
    指令内容。当指令包含内容时，此处填写密文的指令内容。解密方法请继续向下阅读。
```