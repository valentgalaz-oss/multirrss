# cli

Ultra-simple Go CLI module that uses [`github.com/felipeinf/instago`](https://github.com/felipeinf/instago).

## Install

```bash
go install github.com/valentgalaz-oss/multirrss/cli@latest
```

## Usage (PowerShell)

Create a session file:

```powershell
$env:INSTAGO_USERNAME="your_username"
$env:INSTAGO_PASSWORD="your_password"
cli login -session session.json
```

Use an existing session:

```powershell
cli me -session session.json
```

