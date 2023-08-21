配置文件

```json
{
  "stdout": false,
  "fileName": "%Y%m%d.log",
  "timeFormat": "2006-01-02 15:04:05",
  "path": "./logs/",
  "level": "info",
  "maxFile": 7,
  "triggerNum": 50,
  "rotationInterval": 30,
  "notifyInterval": 30,
  "email": {
    "host": "",
    "port": 0,
    "username": "",
    "password": "",
    "recipient": ["example.gmail.com"]
  },
  "wxPusher": {
    "apiToken": "",
    "topicIds": [1],
    "recipient": ["example"]
  }
}
```

```yaml
stdout: false
fileName: "%Y%m%d.log"
timeFormat: "2006-01-02 15:04:05"
path: "./logs/"
level: "info"
maxFile: 7
triggerNum: 50
rotationInterval: 30
notifyInterval: 30
email:
  host: ""
  port: 0
  username: ""
  password: ""
  recipient:
    - sss
wxPusher:
  apiToken: ""
  topicIds:
    - 1
  recipient:
    - example
```