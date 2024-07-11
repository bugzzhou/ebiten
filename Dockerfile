FROM ubuntu:latest

# 安装基本工具和依赖项
RUN apt-get update && apt-get install -y \
    build-essential \
    libxrandr-dev \
    libx11-dev \
    libxcursor-dev \
    libxinerama-dev \
    libxi-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev

# 设置工作目录
WORKDIR /app

# 复制代码到容器中
COPY . .

# 默认命令，可以根据需要修改
CMD ["make"]

