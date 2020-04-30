FROM golang:1.13 as server_builder
WORKDIR /app
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download
COPY ./server ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o linkbucket .

FROM node:alpine as client_builder
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY ./client/package.json ./
COPY ./client/package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent
COPY ./client ./
ARG REACT_APP_LINKBUCKET_URL=
ARG REACT_APP_PAGINATE_LIMIT=15
ENV REACT_APP_LINKBUCKET_URL=${REACT_APP_LINKBUCKET_URL}
ENV REACT_APP_PAGINATE_LIMIT=${REACT_APP_PAGINATE_LIMIT}
RUN npm run build

FROM alpine
WORKDIR /app
RUN mkdir server && echo "" >> ./server/.env
COPY --from=server_builder /app/linkbucket ./server
COPY --from=client_builder /app/build/ ./client/build
RUN pwd && ls
WORKDIR /app/server
CMD ["./linkbucket"]
