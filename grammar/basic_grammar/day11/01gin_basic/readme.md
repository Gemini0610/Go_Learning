Gin框架的基础学习
  1、渲染: json  c.JSON(状态码, 数据)
          HTML  c.HTML(状态码, 模板, 数据) 
          XML   c.XML(状态码, 模板, 数据) 
          YAML   c.YAML(状态码, 数据) 
    
  2、常用的状态码:
        200 OK：表示请求成功，并且服务器成功返回了请求的数据。
        201 Created：表示请求成功，并且服务器成功创建了新的资源。
        204 No Content：表示请求成功，但服务器没有返回任何内容。
        400 Bad Request：表示客户端发送的请求有错误，服务器无法理解或处理该请求。
        401 Unauthorized：表示客户端未经身份验证或身份验证失败，无权访问请求的资源。
        403 Forbidden：表示客户端无权限访问请求的资源，服务器拒绝该请求。
        404 Not Found：表示请求的资源不存在，服务器无法找到请求的资源。
        500 Internal Server Error：表示服务器在处理请求时发生了错误，导致无法完成请求。
        503 Service Unavailable：表示服务器当前无法处理请求，通常是因为服务器过载或维护。

  3、HTTP四种请求:
        GET     获取资源
        POST    创建资源
        PUT/PATCH 更新资源
        DELETE  删除资源
 
  4、参数获取:
        querystring: URL地址后面?的键值对 /search?name=123
        path params: URL地址参数    /posts/2019/06/30
        form params: 表单提交的数据