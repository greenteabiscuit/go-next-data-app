# go-next-data-app

## To debug:

1. clone repository
2. run `docker-compose up` (この時点ではapi serverは立ち上がっておらず、dlvのみ立ち上がっている)
3. Open goland
4. Open `Run` -> `Debug` -> `edit configurations`
5. Choose `go remote`

<img width="1072" alt="Screen Shot 2021-03-09 at 17 43 13" src="https://user-images.githubusercontent.com/44530189/110443279-173fb900-80ff-11eb-8288-2af2117b5654.png">

6. Click `Create Configuration`, and click `Apply` without changing anything.

<img width="699" alt="Screen Shot 2021-03-09 at 17 50 12" src="https://user-images.githubusercontent.com/44530189/110444067-f330a780-80ff-11eb-8425-038b3fe6d84b.png">


7. Click `Debug` to start. 

If successful:

<img width="680" alt="Screen Shot 2021-03-09 at 17 44 00" src="https://user-images.githubusercontent.com/44530189/110443463-47875780-80ff-11eb-8187-cd00c1ffb508.png">

ターミナルでも以下のように表示されるはず

```
go_1  | [GIN-debug] Listening and serving HTTP on :8080
```

8. Set breakpoint and access `localhost:8080`. 以下のように変数が表示されるはずです。

<img width="1053" alt="Screen Shot 2021-03-09 at 17 48 38" src="https://user-images.githubusercontent.com/44530189/110443870-bd8bbe80-80ff-11eb-91e1-dbe979c3a451.png">
