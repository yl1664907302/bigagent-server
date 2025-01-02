# 基础镜像
FROM nginx:latest
# 镜像维护
LABEL maintainer=yeling
# 将 dist 文件夹拷贝到 Nginx 的静态资源目录
COPY dist-pro/. /usr/share/nginx/html/web/
# 将前端 Nginx 配置覆盖基础镜像的配置文件（最好在启动容器的时候挂载到宿主机由运维维护）
COPY bigagent.conf /etc/nginx/conf.d/bigagent.conf
