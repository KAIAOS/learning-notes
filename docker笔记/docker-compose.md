```yml
#docker-compose.yml 固定的名称
version: '2'

services:
  deamon:   #容器的名称
    image: tomcat  #镜像名称 :1.12 指定端口号
    ports:
      - 8008:8080  #-p端口映射
    links:
      - app
    volumes:  #-v数据卷
      - /home/yuyi/app:/usr/local/tomcat/webapps
    restart: alway
  app:
    image: 
```

docker-compose up -d

