#### 1、对称加密

f1(k,data) = X  加密

f2(k, X) = data  解密   

服务端客户端双向都是安全的，但是一旦客户端可以获得k的话 加密是无效的

#### 2、非对称加密

**用私钥加密的数据可以用公钥解密，同样用公钥加密的数据可以用私钥解密**

- f(pk,data) = Y
- f(sk,Y) = data
- f(sk,data) = Y'
- f(pk,Y') = data

客户端向服务端发送请求是安全的，服务端向客户端发送请求是不安全的

客户端请求一个pk，之后请求用pk加密，服务端用sk解密；服务端返回数据用sk加密，客户端用pk解密

#### 3、对称+非对称

先用非对称加密 协商一个对称加密的key，之后用对称加密就安全了

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201224151151818.png" alt="image-20201224151151818" style="zoom:50%;" />

**但是会有中间人问题，黑客模拟服务端发送对称加密的key**

<img src="C:\Users\hanka\AppData\Roaming\Typora\typora-user-images\image-20201224151341036.png" alt="image-20201224151341036" style="zoom:33%;" />

**解决办法CA（Certificate Authority）  CPK、CSK    在CA端 f(CSK,pk) = lisence**   

客户端不再请求一个pk，而是请求license,之后再用CPK解密得到pk

 **HTTPS 是 对称加密+非对称加密+Hash+CA**