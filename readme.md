## 一个随机图片展示服务
一个基础的随机图片展示服务，修改[docker-compose.yml](./docker-compose.yml)文件中的image路径为自己本地的图片根目录，会自动展示目录下的所有图片，每10分钟切换一次图片，随机展示，适合用于不想折腾的轻量图床使用。
## 如何使用
1. 克隆项目到本地
2. 构建镜像
   ```
   docker build -t random_image:local -f dockerfile .
   ```
3. 修改[docker-compose.yml](./docker-compose.yml)文件中的image路径为自己本地的图片根目录
4. 启动服务
   ```
   docker-compose -f docker-compose.yml up -d
   ```
5. 访问ip:8080即可看到文件夹下的随机一张图片