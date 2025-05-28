FROM node:lts-alpine AS build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM node:lts-alpine AS prod

WORKDIR /app

COPY --from=build-stage /app /app

RUN npm install

EXPOSE 5173

CMD [ "npm", "run", "dev" ]