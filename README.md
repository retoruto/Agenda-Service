# Agenda-Service

## **the Api UI :**</br>
![Image text](https://github.com/retoruto/Agenda-Service/blob/master/photos/1.png)
![Image text](https://github.com/retoruto/Agenda-Service/blob/master/photos/2.png)
![Image text](https://github.com/retoruto/Agenda-Service/blob/master/photos/3.png)
![Image text](https://github.com/retoruto/Agenda-Service/blob/master/photos/4.png)
![Image text](https://github.com/retoruto/Agenda-Service/blob/master/photos/5.png)

## **监听端口状态：**</br>
```
$ go run main.go
[negroni] listening on :8080

```
## **UserRegister:**</br>
**客户端**
```
$ ./agenda register -u Alice -p 123 -e Alice@163.com -t 123
Register successfully.The account detail is :
{
  "Name": "Alice",
  "Password": "123",
  "Email": "Alice@163.com",
  "Phone": "123"
}
```
**服务器**
```
[negroni] 2017-12-13T04:10:34-08:00 | 201 | 	 288.047584ms | localhost:8080 | POST /v1/user 

```



