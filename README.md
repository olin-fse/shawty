# How to set up this project

- Add this line to your `~/.bashrc` (Linux) or `~/.bash_profile` (Mac):

```
export GOPATH=$HOME/go
```

- Clone this repository:

```
mkdir -p ~/go/src/github.com/olin-fse
cd !$
git clone https://github.com/olin-fse/shawty
cd shawty
```

- Install dependencies from root directory:

```
go get ./...

cd frontend/
npm install
```

- Install MYSQL. Then create a user with the username `url` and password `password`. Start the MYSQL service.

- Build React code from `shawty/frontend`:

```
npm run build
```

- Run server from root directory:

```
go run main.go
```