FROM node:lts-alpine
# install simple http server for serving static content
RUN npm install -g http-server
# make the 'app' folder the current working directory
WORKDIR /app
# copy 'package.json' to install dependencies
COPY package*.json ./
# install dependencies
RUN npm install
# copy files and folders to the current working directory (i.e. 'app' folder)
COPY . .
# build app for production with minification
RUN npm run build
# provide server port
EXPOSE 5173
CMD [ "http-server", "dist", "-p", "5173" ]