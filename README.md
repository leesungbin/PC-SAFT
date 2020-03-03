# PC-SAFT EOS Calculating Service

> This project is the result of the "Project Design" in University of Seoul department of chemical engineering. You can meet the service on https://saftgo.app .

<br/>

## Project structure

There are 3 folders in the root.

1. **client** : React app(typescript)

   Used [react-konva](https://github.com/konvajs/react-konva) to draw ternary diagram.

2. **schema** : DB schema ( not using now )

   I used PostgreSQL to save chemical properties and deployed to google cloud sql. But now, because of the financial problem I detached cloud sql service and the chemical properties are saved in json.

3. **server** : Go API server

   Server is developed with go, the server can calculate several components' PC-SAFT equation of state.

   There are some codes left which have been used to connect with postgresql, and those codes are not approached for now.

   * **/server/api** : Bubble P , Bubble T, Dew P, Dew T, Flash calculation is available.

   * **/server/jdb** : To make json db work. (data.go file has properties of chemical components)

   * **/server/parser** : not using

   * **/server/ternary** : To calculate ternary coordinates.

   * **/server/ttp** : To communicate with react app.

   * **/server/env** : APP_NAME, PORT, DEBUG value should be given in environment variables. In this folder you should create env file for yourself to run the server.

   * **/server/saft** : shell script to run api server in local. To run the server refer the command below.

     ```bash
     # read environment variables and run the server.
     saft (envfile name in /server/env) run
     ```

<br/>

## SAFT-GO API Documentation

There are korean documentation on [saftgo.app](https://saftgo.app).

In this documents, you'll know how to use the SAFT-GO service.

<br/>

### saftgo.app

> For now, this service is deployed on Google Cloud Run. The calculation power of the server is low, so the calculation may take some time.
>
> Sometime you'll meet weird binodal curve, and in some inadequate conditions you'll be not able to get the desired response.

- saftgo.app is the service to check ternary mixture's ternary diagram.

- You can do your work in 'Program' tab easily.

- If there are not enough conditions to calculate, you'll meet the error message on the bottom of the page.

- If the server can't calculate within 300 seconds, response timeout occurs. If mixtures' all composition  occurs error than the message '계산이 잘 되지 않는 물질입니다.' will be shown.

- If you can't get the calculation results in 5 minutes, you may reload the page or email to saftwithgo@gmail.com .

- Most calculation errors occur when the compressibility factor (Z) is less than zero, or when the iteration is not converging.

- In react app, it uses 4 api endpoint : `/api/equil`, `/api/flashes`, `/api/datas/` `/api/search/`. You may POST request to saftgo with command line. (see 'REST API' document below)

- 계산 결과를 통해 얻은 파란색 점을 연결한 선은 liquid phase의 경계, 분홍색 점을 연결한 선은 vapor phase의 경계를 나타냅니다. 점 위에 마우스를 올리면 조성과 tieline을 볼 수 있습니다.

- phase가 복잡하게 갈릴 경우, 현재의 정렬 알고리즘 상에서는 binodal line이 이상하게 그려질 수 있습니다.

- Exp data 스위치를 켜면, ternary diagram 하단에 값을 입력하여, canvas 내에 plot을 찍어볼 수 있습니다.

  **이 때 띄어쓰기에 유의하십시오.**



### Database

Each column in the database has the following meanings:

| Property Name | Unit      | Description                                |
| ------------- | --------- | ------------------------------------------ |
| Mw            | g/mol     | Molecular weight                           |
| Pc            | atm       | Critical pressure                          |
| Tc            | K         | Critical temperature                       |
| Tb            | K         | Normal boiling point                       |
| ω             |           | Acentric factor                            |
| ε             | ε/k ( K ) | Depth of the potential                     |
| m             |           | Chain length                               |
| σ             | Å         | Segment diameter                           |
| k             |           | Associative volume characteristic constant |
| e             |           | Association energy characteristic constant |
| d             | D         | Dipole moment                              |
| x             |           | non-associated fraction                    |





## REST API

| Endpoint              | Required | Optional | Not Required | Description                                                  |
| --------------------- | -------- | -------- | ------------ | ------------------------------------------------------------ |
| ***/api/equil***      | id       | T or P   | x, y         | Ternary Diagram을 그리기 위한 값을 계산합니다. T가 주어지면 BubbleP, P가 주어지면 BubbleT 계산을 진행합니다. |
| ***/api/flashes***    | id, T, P |          | x, y         | T, P가 고정된 상태에서 Ternary Diagram을 그리기 위한 값을 계산합니다. |
| ***/api/datas***      |          |          |              | Database에서 모든 component들의 열역학적 물성 값을 가져옵니다. |
| ~~***/api/search***~~ | ~~name~~ |          |              | ~~name이 이름에 포함된 component들을 가져옵니다.~~           |
| ***/api/flash***      | id, T, P | x or y   |              | T, P가 고정된 상태에서 조성이 주어졌을 때, Flash 계산을 진행합니다. |
| ***/api/bublp***      | id, T, x |          |              | T, x 가 주어진 상태에서 조성이 주어졌을 때, BubbleP 계산을 진행합니다. |
| ***/api/bublt***      | id, P, x |          |              | P, x 가 주어진 상태에서 조성이 주어졌을 때, BubbleT 계산을 진행합니다. |
| ***/api/dewp***       | id, T, y |          |              | T, y 가 주어진 상태에서 조성이 주어졌을 때, DewP 계산을 진행합니다. |
| ***/api/dewt***       | id, P, y |          |              | P, y 가 주어진 상태에서 조성이 주어졌을 때, DewT 계산을 진행합니다. |

- https://saftgo.app/api 에는 여러개의 endpoint 가 있습니다. 각 endpoint로 적절한 input을 보내면 json 형태로 response 를 받아볼 수 있습니다. (별다른 Authentication 은 없습니다.)

- 모든 endpoint에는 `POST` Method로 요청을 보내야합니다.

- 계산 작업에 필요한 input은 다음과 같은 json 형태입니다.

  ```
  {
    "T": 300,
    "P": 1,
    "id": [10,20,30],
    "x": [0.4,0.5,0.1],
    "y": [0.1,0.2,0.7]
  }
  ```

  > T: Temperature (K)
  >
  > P: Pressure (atm)
  >
  > id: array of components' id (***/search*** 참고)
  >
  > x: array of components' liquid mole fraction
  >
  > y: array of components' vapor mole fraction

  이 중 해당되는 내용을 body에 담아서 request를 보내야합니다.

- 계산이 완료되면 받는 reponse의 형태는 다음과 같습니다.

  ```
  {
    "result": {
      "data": [
        {
          ...
        }
      ],
      "names": [
        "water (polar)", "methanol (polar)", "acetic acid"
      ]
    }
  }
  ```

  > Vvap, Vliq : Molar volume of vapor phase or liquid phase ( m^3/mol )





### CURL example

#### Request

```
curl -X POST https://saftgo.app/api/bublp -d '{ "T": 300, "id": [51, 66, 109], "x": [0.1, 0.2, 0.7] }'
```



#### Response

```
{
    "data": {
        "P": 0.103896178849966,
        "T": 300,
        "x": [
            0.1,
            0.2,
            0.7
        ],
        "y": [
            0.20269533934906378,
            0.49823257919074704,
            0.2991343607762043
        ],
        "Vvap": 0.2326430921880393,
        "Vliq": 0.000026383936903638775
    }
}
```