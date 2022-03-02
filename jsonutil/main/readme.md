# json 스키마 maker (jsonutils)
## 출처
https://pkg.go.dev/github.com/sdvdxl/go-tools/json
## 사용법
```sh
./jsonschema.exe -f config.json
```

## json 예제
config.json
```json
{
    "servers":[
        {
            "Server": "210.1.1.138",
            "UserName": "tstwas",
            "Password":        "tstwas!!",
            "RemoteDirRoot":   "/APP/WAS/UBMWQ/LOG/log4j",
            "RemoteDirs": [
                "C10",
                "M17",
                "M20",
                "M26",
                "...",
                "PDA"

            ],
            "LocalDir":        "C:/ago/rcv",
            "FilePattern":     "",
            "BackupDirectory": "C:/analogo/backup"
        } ,
        {
            "Server": "210.1.1.136",
            "UserName": "tstwas",
            "Password":        "tstwas!!",
            "RemoteDirRoot":   "/APP/WAS/UBMWQ/LOG/log4j",
            "RemoteDirs": [
                "C10"

            ],
            "LocalDir":        "C:/analogo/rcv",
            "FilePattern":     "",
            "BackupDirectory": "C:/analogo/backup"
        }
    ]
}
```

### golang struct 생성 결과
```go
type Config struct {
        Servers []struct {
                BackupDirectory string   `json:"BackupDirectory"`
                FilePattern     string   `json:"FilePattern"`    
                LocalDir        string   `json:"LocalDir"`       
                Password        string   `json:"Password"`       
                RemoteDirRoot   string   `json:"RemoteDirRoot"`  
                RemoteDirs      []string `json:"RemoteDirs"`     
                RemoteDirs111   []string `json:"RemoteDirs111"`  
                Server          net.IP   `json:"Server"`
                UserName        string   `json:"UserName"`       
        } `json:"servers"`
}
```


### Unmarshal 방법

```go
func Config(propertiesFilePath string) ServerConfig {
	dat, err := ioutil.ReadFile(propertiesFilePath)

	if err != nil {
		panic(err)
	}
	var c ServerConfig
	err = json.Unmarshal(dat, &c)
	if err != nil {
		panic(err)
	}
	return c

}
```