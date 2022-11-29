
#### 使用说明

> 阿里云短信
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "sms",
    "time": 1000,
    "message_params": {
        "receiver": "18667143169",
        "variables": {
            "signName":"锣号App",
            "template_id":"SMS_1379552570",
            "map": {
                "code":"112233"
            }
        }
    },
    "message_template_id": 6
}'
```

> 腾讯云短信
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "sms",
    "time": 1000,
    "message_params": {
        "receiver": "18667143169",
        "variables": {
            "app_id":"14007536930",
            "sign_name":"涵睿科技",
            "template_id":"15807890",
            "array": ["1233","5"]
        }
    },
    "message_template_id": 7
}'
```

> 邮件消息
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "email",
    "message_params": {
        "receiver": "test@qq.com",
        "variables": {
            "title": "测试操作",
            "content": "Hello <b>Bob</b> and <i>Cora</i>!"
        }
    },
    "message_template_id": 2
}'
```

> 微信公众号消息
```
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "offiaccount",
    "time":1000,
    "message_params": {
        "template_id":"j-OfIahoJGKC1hHCUEU-XapusCHzL6KTN9D3ntHgOD0",
        "receiver": "oyggS5xOHgKvYo_f2GlZQexBOick",
        "variables": {
            "map": {
                "first":"张三12333|#FF0000",
                "keyword1":"12212321|#FF0000",
                "remark":"12312fsfsdfsdf"
            },
            "url":"http://www.baidu.com"
        }
    },
    "message_template_id":8
}'

//参数带颜色的
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "offiaccount",
    "message_params": {
        "template_id":"j-OfIahoJGKC1hHCUEU-XapusCHzL6KTN9D3ntHgOD0",
        "receiver": "openId",
        "variables": {
            "map": {
                "name":"张三12333|#0000FF"
            },
            "url": "https://www.baidu.com/"
        }
    },
    "message_template_id": 4
}'
```

> 钉钉自定义机器人
```
//艾特某些手机号
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "message_params": {
        "receiver": "13588888888,13588888887",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "message_template_id": 5
}'

//艾特全部人
curl --location --request POST 'http://localhost:8888/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "message_params": {
        "receiver": "@all",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "message_template_id": 5
}'
```