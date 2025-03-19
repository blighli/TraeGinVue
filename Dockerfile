FROM node:18-alpine as build

# 设置 npm 镜像
RUN npm config set registry https://registry.npmmirror.com

WORKDIR /app

# 只复制依赖文件
COPY package*.json ./
RUN npm install

# 分层复制其他文件
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]