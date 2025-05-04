FROM node:23-alpine

WORKDIR /app

COPY . .
RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["npm", "run", "preview", "--", "--port", "3000", "--host", "0.0.0.0"]
