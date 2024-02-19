# Compiling the proto files

1) Head to https://github.com/protocolbuffers/protobuf/releases and download the zip for your machine
    - Note: for windows, you may to expand the list for the download to appear
2) Extract the zip to your C: drive under the name of the zip folder
3) Add to your system environment variables
    - Under System variables, find Path, and double click. 
    - Add name_of_folder\bin to the list of paths

    Run the following command to compile the proto file:
    ```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative file.proto
    ```

    If protoc isn't being detected as a command, run it in cmd.


# Getting the server and client to run 

The following has already been executed, but to replicate on your machine, you can do the following.

```go mod init example.com/fileserverproject```

The above command will allow you to have the pb import in the server main.go file with the import path of "example.com/fileserverproject"

To recognize the grpc imports, run the following:

```go get google.golang.org/grpc```

# Other Notes
If there is an issue where vsCode deletes your code (unused imports), add the following to your vscode settings.json 
```
  "[go]": {

        "editor.formatOnSave": false,
        "editor.codeActionsOnSave": {
            "source.organizeImports": false
        },
    },
    "go.formatTool": "gofmt",
```
