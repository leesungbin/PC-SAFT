# SAFT-GO Documentation

이 문서에서는 SAFT-GO 서비스를 어떻게 사용할 수 있는지에 관해서 다룹니다.



## saftgo.app

> 현재 이 서비스는 Google Cloud Run 을 통해 배포되었습니다. 따라서 계산에 시간이 오래 걸릴 수 있습니다.

* saftgo.app 은 ternary mixture의 평형도를 시각적으로 확인할 수 있는 서비스입니다.

* Program 탭에서 해당 작업을 할 수 있습니다.

* 계산을 하기 위한 조건이 갖추어지지 않으면 페이지 하단에 오류 메세지가 등장합니다.

* response timeout은 300초로, 이 안에 계산결과를 얻지 못하거나, 모든 조성에서 계산 오류가 난 경우에는 **계산이 잘 되지 않는 물질입니다.** 라는 메세지를 볼 수 있습니다.

* 계산 결과를 5분이 지나도 얻지 못한 경우에는 페이지를 **새로고침**하여 다시 시도하거나, saftwithgo@gmail.com 으로 문의 바랍니다. 

* 대부분의 계산 오류는 압축인자(Z)가 0 보다 작아지거나, 반복문을 최대한으로 돌아도 수렴이 되지 않을 때 생깁니다.

* 이 서비스에서 사용된 api endpoint 는 ***/equil*** , ***/flashes***, ***/datas***, ***/search*** 입니다. Command Line을 통해서 직접 POST request를 보내볼 수 있습니다. (REST API 탭을 참고하십시오.)

* 계산 결과를 통해 얻은 파란색 점은 liquid, 분홍색 점은 vapor phase를 나타냅니다. 점 위에 마우스를 올리면 조성과 tieline을 볼 수 있습니다.

  

## REST API
| Endpoint       | Required | Optional | Not Required |
| -------------- | -------- | -------- | ------------ |
| ***/equil***   | id       | T or P   | x, y         |
| ***/flashes*** | id, T, P |          | x, y         |
| ***/datas***   |          |          | everything   |
| ***/search***  | name     |          |              |
| ***/flash***   | id, T, P | x or y   |              |
| ***/bublp***   | id, T, x |          |              |
| ***/bublt***   | id, P, x |          |              |
| ***/dewp***    | id, T, y |          |              |
| ***/dewt***    | id, P, y |          |              |


* https://saftgo.app/api 에는 여러개의 endpoint 가 있습니다. 각 endpoint로 적절한 input을 보내면 json 형태로 response 를 받아볼 수 있습니다. (별다른 Authentication 은 없습니다.)

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

