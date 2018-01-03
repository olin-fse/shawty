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

- Install MYSQL, then start it (if not the app will not work)

- Build React code from `shawty/frontend`:

```
npm run build
```

- Run server from root directory:

```
go run main.go
```