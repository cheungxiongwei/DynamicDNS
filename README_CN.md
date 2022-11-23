### 如何使用 ？
运行 `DynamicDNS.exe` 程序，同时提供一个 `config.json` 配置文件.

程序行为默认读取程序根目录 config.json 配置文件.
例如如下配置会自动 15 分钟更新一次远程 dns 地址.

host：二级域名 或 @  
domain：一级域名  
password：动态DNS服务密码  
time_out：自动更新时间，单位秒
```json
{
  "host": "ddns",
  "domain": "google.com",
  "password": "0656f00f07714951a3974ada7a485559",
  "time_out": 900
}
```


### 如何获取本机 IP 地址 ？

* https://ip4.me/
* https://test-ipv4.com/
* https://ipv6-test.com/

通过一个 api 列表查询本机 ip 地址，把查询的结果通过域名提供商提供的 api 更新动态DNS，这样通过域名就可以获取电脑 ip 地址了。

域名提供商：https://www.namecheap.com/

namecheap 动态域名服务远程修改 api：
https://www.namecheap.com/support/knowledgebase/article.aspx/29/11/how-to-dynamically-update-the-hosts-ip-with-an-http-request/

```
例子：
https://dynamicdns.park-your-domain.com/update?host=[host]&domain=[domain_name]&password=[ddns_password]&ip=[your_ip]

# Fields that are to be changed in the URL are: [host], [domain_name], [ddns_password] and [your_ip].
```

响应值：
```xml
<?xml version="1.0" encoding="utf-16"?>
<interface-response>
    <Command>SETDNSHOST</Command>
    <Language>eng</Language>
    <IP>xxx.x.x.xxx</IP>
    <ErrCount>0</ErrCount>
    <errors/>
    <ResponseCount>0</ResponseCount>
    <responses/>
    <Done>true</Done>
    <debug><![CDATA[
]]></debug>
</interface-response>
```