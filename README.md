# FullStackAssignment

## System Dependencies

  * Golang installed
  * NVM installed

## Clone the repository

```
git clone https://github.com/charangk01/FullStackAssignment.git
```

## Setting Up the Project
## Backend (Golang)

1. Navigate to the backend directory
2. Setup Golang Environment
3. Run the Backend Server
```
Local ENV : go run main.go
```
```
Using Docker:
docker build -t golang-lru-cache .
docker run -p 8080:8080 golang-lru-cache
```

## Frontend (ReactJs)

1. Navigate to the Frontend directory
2. Setup the Environment
3. Install the dependenies & Start the server
```
Local ENV :
npm install
npm start
```
```
Using Docker :
docker build -t react-lru-cache .
docker run -p 3000:80 react-lru-cache
```

## Using Docker Compose

```
docker-compose up --build
```

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
