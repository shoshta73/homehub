FROM node:22-alpine3.20 AS base

FROM base AS deps

WORKDIR /app

COPY web/package.json web/yarn.lock* web/package-lock.json* web/pnpm-lock.yaml* web/.npmrc* ./
RUN \
  if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
  elif [ -f package-lock.json ]; then npm ci; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm i --frozen-lockfile; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM base AS build

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules

COPY web/src ./src
COPY web/tsconfig.json ./
COPY web/tsconfig.app.json ./
COPY web/tsconfig.node.json ./
COPY web/vite.config.ts ./
COPY web/index.html ./

COPY web/package.json web/yarn.lock* web/package-lock.json* web/pnpm-lock.yaml* web/.npmrc* ./
RUN \
  if [ -f yarn.lock ]; then yarn build; \
  elif [ -f package-lock.json ]; then npm run build; \
  elif [ -f pnpm-lock.yaml ]; then corepack enable pnpm && pnpm run build; \
  else echo "Lockfile not found." && exit 1; \
  fi

FROM nginx:alpine3.20

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
