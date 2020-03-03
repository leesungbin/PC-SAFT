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

- The line that connects the blue point obtained by the calculation represents the boundary of the liquid phase, and the line connects the pink point represents the boundary of the vapor phase. Move the mouse over the point to see the composition and the tieline.

- If the phase is complicated, the binodal line can be drawn strangely on the current sorting algorithm.

- Turn on the Exp data switch, to plot points on ternary diagram.

  ```text
  0.1 0.2 0.7
  0.2 0.21 0.59
  ```

  **Please be careful of spacing.**

<br/>

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

<br/>

## REST API

| Endpoint              | Required | Optional | Not Required | Description                                                  |
| --------------------- | -------- | -------- | ------------ | ------------------------------------------------------------ |
| ***/api/equil***      | id       | T or P   | x, y         | Calculate the compositions of liquid, vapor phases. If T is fixed(or given), Bubble P calculation occurs.(P given -> Bubble T. Prefer bubble calculation than dew) |
| ***/api/flashes***    | id, T, P |          | x, y         | Calculate the compositions of liquid, vapor phases when T, P are fixed. |
| ***/api/datas***      |          |          |              | Get all components' thermodynamic properties in the database. |
| ~~***/api/search***~~ | ~~name~~ |          |              | ~~name이 이름에 포함된 component들을 가져옵니다.~~           |
| ***/api/flash***      | id, T, P | x or y   |              | Do flash calculation when T, P are fixed.                    |
| ***/api/bublp***      | id, T, x |          |              | Do bubble p calculation when T, x are fixed.                 |
| ***/api/bublt***      | id, P, x |          |              | Do bubble p calculation when P, x are fixed.                 |
| ***/api/dewp***       | id, T, y |          |              | Do dew p calculation when T, y are fixed.                    |
| ***/api/dewt***       | id, P, y |          |              | Do dew p calculation when P, y are fixed.                    |

- https://saftgo.app/api has several endpoints. Each endpoints wants adequate input in json, and you'll get json response. (There are no special authentications.)

- All API endpoints only listen to POST requests.

- The shape of the input needed for the calculation is like below:

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
  > id: array of components' id (refer ***search*** tab)
  >
  > x: array of components' liquid mole fraction
  >
  > y: array of components' vapor mole fraction

  You should contain adequate values in the body and send the request,

- The shape of the response is like below:

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