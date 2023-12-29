# go-tdlib

Go wrapper for [TDLib (Telegram Database Library)](https://github.com/tdlib/td) with full support of the TDLib.
Current supported version of TDLib corresponds to the commit hash [1a50ec4](https://github.com/tdlib/td/commit/1a50ec474ce2c2c09017aa3ab9cc9e0c68f483fc), updated on 2023-12-17

## TDLib installation

Use [TDLib build instructions](https://tdlib.github.io/td/build.html) with checkmarked `Install built TDLib to /usr/local instead of placing the files to td/tdlib`.

### Windows

Build with environment variables:

```
CGO_CFLAGS=-IC:/path/to/tdlib/build/tdlib/include
CGO_LDFLAGS=-LC:/path/to/tdlib/build/tdlib/bin -ltdjson
```

Example for PowerShell:

```powershell
$env:CGO_CFLAGS="-IC:/td/tdlib/include"; $env:CGO_LDFLAGS="-LC:/td/tdlib/bin -ltdjson"; go build -trimpath -ldflags="-s -w" -o demo.exe .\cmd\demo.go
```

## Usage

### Client

[Register an application](https://my.telegram.org/apps) to obtain an api_id and api_hash 

```go
package main

import (
    "log"
    "path/filepath"

    "github.com/wildfluss/go-tdlib/client"
)

func main() {
    // client authorizer
    authorizer := client.ClientAuthorizer()
    go client.CliInteractor(authorizer)

    // or bot authorizer
    // botToken := "000000000:gsVCGG5YbikxYHC7bP5vRvmBqJ7Xz6vG6td"
    // authorizer := client.BotAuthorizer(botToken)

    const (
        apiId   = 00000
        apiHash = "8pu9yg32qkuukj83ozaqo5zzjwhkxhnk"
    )

    authorizer.TdlibParameters <- &client.SetTdlibParametersRequest{
        UseTestDc:              false,
        DatabaseDirectory:      filepath.Join(".tdlib", "database"),
        FilesDirectory:         filepath.Join(".tdlib", "files"),
        UseFileDatabase:        true,
        UseChatInfoDatabase:    true,
        UseMessageDatabase:     true,
        UseSecretChats:         false,
        ApiId:                  apiId,
        ApiHash:                apiHash,
        SystemLanguageCode:     "en",
        DeviceModel:            "Server",
        SystemVersion:          "1.0.0",
        ApplicationVersion:     "1.0.0",
        EnableStorageOptimizer: true,
        IgnoreFileNames:        false,
    }

	_, err := client.SetLogVerbosityLevel(&client.SetLogVerbosityLevelRequest{
		NewVerbosityLevel: 1,
	})
	if err != nil {
		log.Fatalf("SetLogVerbosityLevel error: %s", err)
	}
	
    tdlibClient, err := client.NewClient(authorizer)
    if err != nil {
        log.Fatalf("NewClient error: %s", err)
    }

    optionValue, err := client.GetOption(&client.GetOptionRequest{
        Name: "version",
    })
    if err != nil {
        log.Fatalf("GetOption error: %s", err)
    }

    log.Printf("TDLib version: %s", optionValue.(*client.OptionValueString).Value)

    me, err := tdlibClient.GetMe()
    if err != nil {
        log.Fatalf("GetMe error: %s", err)
    }

    log.Printf("Me: %s %s [%s]", me.FirstName, me.LastName, me.Username)
}

```

### Receive updates

```go
tdlibClient, err := client.NewClient(authorizer)
if err != nil {
    log.Fatalf("NewClient error: %s", err)
}

listener := tdlibClient.GetListener()
defer listener.Close()
 
for update := range listener.Updates {
    if update.GetClass() == client.ClassUpdate {
        log.Printf("%#v", update)
    }
}
```

### Proxy support

```go
proxy := client.WithProxy(&client.AddProxyRequest{
    Server: "1.1.1.1",
    Port:   1080,
    Enable: true,
    Type: &client.ProxyTypeSocks5{
        Username: "username",
        Password: "password",
    },
})

tdlibClient, err := client.NewClient(authorizer, proxy)

```

## Example

[Example application](https://github.com/wildfluss/go-tdlib/tree/master/example)

```
cd example
docker build --network host --build-arg TD_COMMIT=daf4801 --tag tdlib-test .
docker run --rm -it -e "API_ID=00000" -e "API_HASH=abcdef0123456789" tdlib-test ash
./app
```

## Notes

* WIP. Library API can be changed in the future
* The package includes a .tl-parser and generated [json-schema](https://github.com/wildfluss/go-tdlib/tree/master/data) for creating libraries in other languages

## Author

[Aleksandr Zelenin](https://github.com/wildfluss/), e-mail: [aleksandr@zelenin.me](mailto:aleksandr@zelenin.me)
