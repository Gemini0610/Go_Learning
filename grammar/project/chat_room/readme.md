实现了一个简单的网络聊天室:
    刚学习完网络编程,熟悉一下基本语法、基本知识
    功能如下：
        1、用户的上线下线
        2、聊天、其他人，自己都可以看见聊天信息
        3、查询当前聊天室用户名字
        4、可以修改自己的名字
        5、用户连接超时，系统自动踢出聊天室

    技术点:
        1、socket tcp编程
        2、map结构：存储所有的用户、map遍历、map删除
        3、goroutine、channel
        4、select 超时退出、主动退出
        5、timer定时器