# BMR
A generalized data messaging protocol designed to carry messages across chains.

#### Installation
```bash
go install
```

#### Development Server
```bash
make watch # hot reloads on server
```
or
```bash
air # https://github.com/cosmtrek/air required
```

### Production
```bash
go build
bmr start 
```

### Generate ABI in GoLang
```bash
abigen --abi=Loopso.abi --pkg=contracts --type=Loopso --out=Loopso.go
```
