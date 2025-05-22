FROM node:lts-alpine AS build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:lts-alpine AS production-stage

WORKDIR /app

COPY --from=build-stage /app /app

RUN npm install --production

EXPOSE 3000

CMD [ "npm", "run", "start" ]