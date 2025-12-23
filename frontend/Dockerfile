FROM node:24-alpine3.21

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 1234

RUN npm run dev

ENTRYPOINT [ "npm", "run", "start" ]