FROM node:12.18 as builder

# set working directory
WORKDIR /src

# add `/app/node_modules/.bin` to $PATH
ENV PATH /src/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY yarn.lock ./
RUN yarn install

# add app
COPY . /src

RUN yarn build

FROM caddy:2

COPY deploy/Caddyfile /etc/caddy/Caddyfile
COPY --from=builder /src/build /app
