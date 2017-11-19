## Intro

TBD

## Test

http POST http://127.0.0.1:8443/rest/v1/users name=1 phone=137 password=123456
http GET http://127.0.0.1:8443/rest/v1/users/1
http GET http://127.0.0.1:8443/rest/v1/users/1 'Authorization:Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIxIiwiZXhwIjoxNTEwNDE0MTExfQ.Q19pgjX9u3jl_1fQWGF9AGzx5RUJqRLQTTamw6uJpAgt1Ti27yQtIE9E5_rsK57JzYNL4XuHd42hD566eK5dEsFbXeCIbJT_YCBRT30u8Bkxbpg0FWnbmuEZPpXxnh9sHylEyblph5vN8hbuDa8uXzxFSuU-5OUxV7W1Kf3XB08'

http POST http://127.0.0.1:8443/rest/v1/auth phone=137 password=123456

http POST http://127.0.0.1:8443/rest/v1/users/1/channels name=gggg phone=13712131313 intro=intro
http GET http://127.0.0.1:8443/rest/v1/users/1/channels/1


http POST http://127.0.0.1:8443/rest/v1/users/1/customs weixin=gggg phone=13712131313 address=shenzhen postCode=12345
http GET http://127.0.0.1:8443/rest/v1/users/1/customs/1


http POST http://127.0.0.1:8443/rest/v1/users/1/goods name=gggg catalog=food intro=intro channelID:=1 sellPrice:=60.01 purchasePrice:=50.02
http GET http://127.0.0.1:8443/rest/v1/users/1/goods/1


http POST http://127.0.0.1:8443/rest/v1/users/1/orders customID:=1 goodsID:=1 sellPrice:=60.01
http GET http://127.0.0.1:8443/rest/v1/users/1/orders/1



http GET http://127.0.0.1:8443/rest/v1/users/1/stats/orders