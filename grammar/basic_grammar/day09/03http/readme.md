HTTP:应用层的协议
    1、底层依赖传输层：tcp(短连接)、网络层(IP)
    2、无状态的，每一次请求都是独立的，下次请求需要重新建立连接

HTTPS：
    1、http是标准协议，https不是标准协议。
    2、https = http + ssl(非对称加密, 数字证书)
    3、现在所有网站都会尽量要求使用https开发

一个htto请求体可以分为四部分:
    1、请求行: 请求方法、请求URL、HTTP协议及版本
            方法+URL+协议版本号
            方法：GET获取数据、POST上传数据层(表单\JSON格式)、PUT修改数据、Delete用于删除数据

    2、请求头: 报文头、报文体 K-V格式
            格式：Key：Value
            可以有很多个键值对(包含协议自带或者自定义)
            重要字段：(1)Accpet:接收数据的格式
                     (2)User-Agent:描述用户浏览器的信息
                     (3)Connection:Keep-Alive(长连接)，Close(短连接)
                     (4)Accept-Encoding:gzip, xxx描述可以接收的编码
                     (5)Cookie:由服务器设置的key=value数据，客户端下次请求的时候可以携带过来，快速建立连接。
                     (6)Content-Type: application/-form表示上传的数据是表单格式  application/-json表示上传的数据是json格式
                     (7)用户自定义的

    3、空行
        告诉服务器请求头结束，用于分割

    4、请求包体(可选的)
        一般POST方法，会配套提供BODY
        在GET的时候也可能提供BODY
        上传两种数据格式:
            (1)表单form (2)json

前后端传输数据方法：
    1、放在请求头中
    2、放在请求包中
    3、放在url中:GET/user>id=1011&score=90&school=sut


HTTP响应消息格式
    1、状态行
        协议格式：协议版本号+状态码+状态描述
                    HTTP/1.1 + 200 + OK
        常用状态码:
            1xx===>客户端可以继续发送请求
            2xx===>正常访问，200
            3xx===>重定向
            4xx===>401未授权、404notfound
            5xx===>501服务器内部错误

    2、响应头
        实例:Connect-Type:application/json
             Server:Apache
             Data：Mon,12 Sep

    3、空行

    4、响应包体
        通常是返回json数据