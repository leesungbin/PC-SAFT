# SAFT-GO Documentation

이 문서에서는 SAFT-GO 서비스를 어떻게 사용할 수 있는지에 관해서 다룹니다.



## saftgo.app

> 현재 이 서비스는 Google Cloud Run 을 통해 배포되었습니다. 따라서 계산에 시간이 오래 걸릴 수 있습니다.
>
> 또한 Binodal curve가 이상하게 나올 수 있으며, 적절하지 않은 조건에서는 결과를 얻지 못할 수 있습니다.

* saftgo.app 은 ternary mixture의 평형도를 시각적으로 확인할 수 있는 서비스입니다.

* Program 탭에서 해당 작업을 할 수 있습니다.

* 계산을 하기 위한 조건이 갖추어지지 않으면 페이지 하단에 오류 메세지가 등장합니다.

* response timeout은 300초로, 이 안에 계산결과를 얻지 못하거나, 모든 조성에서 계산 오류가 난 경우에는 **계산이 잘 되지 않는 물질입니다.** 라는 메세지를 볼 수 있습니다.

* 계산 결과를 5분이 지나도 얻지 못한 경우에는 페이지를 **새로고침**하여 다시 시도하거나, saftwithgo@gmail.com 으로 문의 바랍니다. 

* 대부분의 계산 오류는 압축인자(Z)가 0 보다 작아지거나, 반복문을 최대한으로 돌아도 수렴이 되지 않을 때 생깁니다.

* 이 서비스에서 사용된 api endpoint 는 ***/api/equil*** , ***/api/flashes***, ***/api/datas***, ***/api/search*** 입니다. Command Line을 통해서 직접 POST request를 보내볼 수 있습니다. (REST API 탭을 참고하십시오.)

* 계산 결과를 통해 얻은 파란색 점은 liquid, 분홍색 점은 vapor phase를 나타냅니다. 점 위에 마우스를 올리면 조성과 tieline을 볼 수 있습니다.

* Exp data 스위치를 켜면, ternary diagram 하단에 값을 입력하여, canvas 내에 plot을 찍어볼 수 있습니다.

  **이 때 띄어쓰기에 유의하십시오.**

  

<br/>

### Database

Database의 각 column의 의미는 다음과 같습니다.

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

<br/>

<br/>

## REST API
| Endpoint           | Required | Optional | Not Required | Description                                                  |
| ------------------ | -------- | -------- | ------------ | ------------------------------------------------------------ |
| ***/api/equil***   | id       | T or P   | x, y         | Ternary Diagram을 그리기 위한 값을 계산합니다. T가 주어지면 BubbleP, P가 주어지면 BubbleT 계산을 진행합니다. |
| ***/api/flashes*** | id, T, P |          | x, y         | T, P가 고정된 상태에서 Ternary Diagram을 그리기 위한 값을 계산합니다. |
| ***/api/datas***   |          |          |              | Database에서 모든 component들의 열역학적 물성 값을 가져옵니다. |
| ***/api/search***  | name     |          |              | name이 이름에 포함된 component들을 가져옵니다.               |
| ***/api/flash***   | id, T, P | x or y   |              | T, P가 고정된 상태에서 조성이 주어졌을 때, Flash 계산을 진행합니다. |
| ***/api/bublp***   | id, T, x |          |              | T, x 가 주어진 상태에서 조성이 주어졌을 때, BubbleP 계산을 진행합니다. |
| ***/api/bublt***   | id, P, x |          |              | P, x 가 주어진 상태에서 조성이 주어졌을 때, BubbleT 계산을 진행합니다. |
| ***/api/dewp***    | id, T, y |          |              | T, y 가 주어진 상태에서 조성이 주어졌을 때, DewP 계산을 진행합니다. |
| ***/api/dewt***    | id, P, y |          |              | P, y 가 주어진 상태에서 조성이 주어졌을 때, DewT 계산을 진행합니다. |


* https://saftgo.app/api 에는 여러개의 endpoint 가 있습니다. 각 endpoint로 적절한 input을 보내면 json 형태로 response 를 받아볼 수 있습니다. (별다른 Authentication 은 없습니다.)

* 모든 endpoint에는 `POST` Method로 요청을 보내야합니다.

* 계산 작업에 필요한 input은 다음과 같은 json 형태입니다.

  ```json
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

  <br/>

* 계산이 완료되면 받는 reponse의 형태는 다음과 같습니다.

  ```json
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

<br/>

<br/>

### CURL example

#### Request

```shell
curl -X POST https://saftgo.app/api/bublp -d '{ "T": 300, "id": [51, 66, 109], "x": [0.1, 0.2, 0.7] }'
```

<br/>

#### Response

```json
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