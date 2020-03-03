package jdb

const JSON = `{
  "data": [
    {
      "id": 0,
      "data": {
        "name": "1,2-propylene oxide (polar)",
        "Mw": 58.080002,
        "Tc": 482.20001,
        "Pc": 49.200001,
        "w": 0.26899999,
        "Tb": 308,
        "m": 2.0105,
        "sigma": 3.6094999,
        "epsilon": 258.82001,
        "k": 0,
        "e": 0,
        "d": 2,
        "x": 0.39789999
      }
    },
    {
      "id": 1,
      "data": {
        "name": "1-butanol (polar)",
        "Mw": 74.123001,
        "Tc": 563.09998,
        "Pc": 44.200001,
        "w": 0.59299999,
        "Tb": 390.89999,
        "m": 2.9286001,
        "sigma": 3.5235,
        "epsilon": 239.23,
        "k": 0.016349999,
        "e": 2371.8,
        "d": 1.7,
        "x": 0.23902
      }
    },
    {
      "id": 2,
      "data": {
        "name": "1-butanol",
        "Mw": 74.123001,
        "Tc": 563.09998,
        "Pc": 44.200001,
        "w": 0.59299999,
        "Tb": 390.89999,
        "m": 2.7514999,
        "sigma": 3.6138999,
        "epsilon": 259.59,
        "k": 0.0066920002,
        "e": 2544.6001,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 3,
      "data": {
        "name": "1-butene",
        "Mw": 56.108002,
        "Tc": 419.60001,
        "Pc": 40.200001,
        "w": 0.191,
        "Tb": 266.89999,
        "m": 2.2864001,
        "sigma": 3.6431,
        "epsilon": 222,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 4,
      "data": {
        "name": "1-chlorobutane",
        "Mw": 92.569,
        "Tc": 542,
        "Pc": 36.799999,
        "w": 0.21799999,
        "Tb": 351.60001,
        "m": 2.8585,
        "sigma": 3.6424,
        "epsilon": 258.66,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 5,
      "data": {
        "name": "1-heptanol (polar)",
        "Mw": 116.204,
        "Tc": 633,
        "Pc": 30.4,
        "w": 0.56,
        "Tb": 449.79999,
        "m": 4.1961999,
        "sigma": 3.6119001,
        "epsilon": 254.99001,
        "k": 0.0026,
        "e": 2662.7,
        "d": 1.7,
        "x": 0.16682
      }
    },
    {
      "id": 6,
      "data": {
        "name": "1-heptanol",
        "Mw": 116.204,
        "Tc": 633,
        "Pc": 30.4,
        "w": 0.56,
        "Tb": 449.79999,
        "m": 4.3985,
        "sigma": 3.5450001,
        "epsilon": 253.46001,
        "k": 0.0011549999,
        "e": 2878.5,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 7,
      "data": {
        "name": "1-hexanol (polar)",
        "Mw": 102.177,
        "Tc": 611,
        "Pc": 40.5,
        "w": 0.56,
        "Tb": 430.20001,
        "m": 2.8971,
        "sigma": 3.9467001,
        "epsilon": 267.32999,
        "k": 0.0081500001,
        "e": 2769.3999,
        "d": 1.7,
        "x": 0.24162
      }
    },
    {
      "id": 8,
      "data": {
        "name": "1-hexanol",
        "Mw": 102.177,
        "Tc": 611,
        "Pc": 40.5,
        "w": 0.56,
        "Tb": 430.20001,
        "m": 3.5146,
        "sigma": 3.6735001,
        "epsilon": 262.32001,
        "k": 0.0057470002,
        "e": 2538.8999,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 9,
      "data": {
        "name": "1-hexene",
        "Mw": 84.163002,
        "Tc": 504,
        "Pc": 31.700001,
        "w": 0.285,
        "Tb": 336.60001,
        "m": 2.9853001,
        "sigma": 3.7753,
        "epsilon": 236.81,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 10,
      "data": {
        "name": "1-octanol (polar)",
        "Mw": 130.231,
        "Tc": 652.5,
        "Pc": 28.6,
        "w": 0.58700001,
        "Tb": 468.29999,
        "m": 4.3683,
        "sigma": 3.7077999,
        "epsilon": 260.48999,
        "k": 0.0028200001,
        "e": 2660.3999,
        "d": 1.7,
        "x": 0.16024999
      }
    },
    {
      "id": 11,
      "data": {
        "name": "1-octanol",
        "Mw": 130.231,
        "Tc": 652.5,
        "Pc": 28.6,
        "w": 0.58700001,
        "Tb": 468.29999,
        "m": 4.3555002,
        "sigma": 3.7145,
        "epsilon": 262.73999,
        "k": 0.002197,
        "e": 2754.8,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 12,
      "data": {
        "name": "1-octene",
        "Mw": 112.216,
        "Tc": 566.70001,
        "Pc": 26.200001,
        "w": 0.38600001,
        "Tb": 394.39999,
        "m": 3.7423999,
        "sigma": 3.8132999,
        "epsilon": 243.02,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 13,
      "data": {
        "name": "1-pentanol (polar)",
        "Mw": 88.150002,
        "Tc": 588.20001,
        "Pc": 39.099998,
        "w": 0.579,
        "Tb": 411.10001,
        "m": 3.8132999,
        "sigma": 3.3910999,
        "epsilon": 239.75999,
        "k": 0.01303,
        "e": 2079.3999,
        "d": 1.7,
        "x": 0.18357
      }
    },
    {
      "id": 14,
      "data": {
        "name": "1-pentanol",
        "Mw": 88.150002,
        "Tc": 588.20001,
        "Pc": 39.099998,
        "w": 0.579,
        "Tb": 411.10001,
        "m": 3.6259999,
        "sigma": 3.4507999,
        "epsilon": 247.28,
        "k": 0.010319,
        "e": 2252.1001,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 15,
      "data": {
        "name": "1-pentene",
        "Mw": 70.135002,
        "Tc": 464.79999,
        "Pc": 35.299999,
        "w": 0.233,
        "Tb": 303.10001,
        "m": 2.6006,
        "sigma": 3.7399001,
        "epsilon": 231.99001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 16,
      "data": {
        "name": "1-propanol (polar)",
        "Mw": 60.096001,
        "Tc": 536.79999,
        "Pc": 51.700001,
        "w": 0.62300003,
        "Tb": 370.29999,
        "m": 2.6268001,
        "sigma": 3.3917999,
        "epsilon": 219.13,
        "k": 0.020959999,
        "e": 2479.3999,
        "d": 1.7,
        "x": 0.26625001
      }
    },
    {
      "id": 17,
      "data": {
        "name": "1-propanol",
        "Mw": 60.096001,
        "Tc": 536.79999,
        "Pc": 51.700001,
        "w": 0.62300003,
        "Tb": 370.29999,
        "m": 2.9997001,
        "sigma": 3.2521999,
        "epsilon": 233.39999,
        "k": 0.015268,
        "e": 2276.8,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 18,
      "data": {
        "name": "1-propylamine (N-PROPYL AMINE)",
        "Mw": 59.112,
        "Tc": 497,
        "Pc": 48.099998,
        "w": 0.303,
        "Tb": 321.70001,
        "m": 2.4539001,
        "sigma": 3.5346999,
        "epsilon": 250.52,
        "k": 0.022674,
        "e": 1028.1,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 19,
      "data": {
        "name": "2,2-dimethyl butane",
        "Mw": 86.178001,
        "Tc": 488.79999,
        "Pc": 30.799999,
        "w": 0.23199999,
        "Tb": 322.79999,
        "m": 2.6008,
        "sigma": 4.0042,
        "epsilon": 243.50999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 20,
      "data": {
        "name": "2,3-dimethyl butane",
        "Mw": 86.178001,
        "Tc": 500,
        "Pc": 31.299999,
        "w": 0.24699999,
        "Tb": 331.10001,
        "m": 2.6853001,
        "sigma": 3.9545,
        "epsilon": 246.07001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 21,
      "data": {
        "name": "2-butanol (polar)",
        "Mw": 74.123001,
        "Tc": 536.09998,
        "Pc": 41.799999,
        "w": 0.57700002,
        "Tb": 372.70001,
        "m": 2.8566,
        "sigma": 3.5543001,
        "epsilon": 239.00999,
        "k": 0.0079800002,
        "e": 2314.8,
        "d": 1.7,
        "x": 0.24505
      }
    },
    {
      "id": 22,
      "data": {
        "name": "2-chloropropane",
        "Mw": 78.542,
        "Tc": 485,
        "Pc": 47.200001,
        "w": 0.23199999,
        "Tb": 308.89999,
        "m": 2.4151001,
        "sigma": 3.6184001,
        "epsilon": 251.47,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 23,
      "data": {
        "name": "2-methyl hexane",
        "Mw": 100.205,
        "Tc": 530.40002,
        "Pc": 27.299999,
        "w": 0.329,
        "Tb": 363.20001,
        "m": 3.3478,
        "sigma": 3.8612001,
        "epsilon": 237.42,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 24,
      "data": {
        "name": "2-methyl pentane",
        "Mw": 86.178001,
        "Tc": 497.5,
        "Pc": 30.1,
        "w": 0.278,
        "Tb": 333.39999,
        "m": 2.9317,
        "sigma": 3.8534999,
        "epsilon": 235.58,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 25,
      "data": {
        "name": "2-methyl-2-butanol",
        "Mw": 88.150002,
        "Tc": 545,
        "Pc": 39.5,
        "w": 9e+30,
        "Tb": 375.5,
        "m": 2.5487001,
        "sigma": 3.9052999,
        "epsilon": 266.01001,
        "k": 0.001863,
        "e": 2618.8,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 26,
      "data": {
        "name": "2-propanol (polar)",
        "Mw": 60.096001,
        "Tc": 508.29999,
        "Pc": 47.599998,
        "w": 0.66500002,
        "Tb": 355.39999,
        "m": 2.6856,
        "sigma": 3.3800001,
        "epsilon": 199.10001,
        "k": 0.022369999,
        "e": 2473.8,
        "d": 1.7,
        "x": 0.26065001
      }
    },
    {
      "id": 27,
      "data": {
        "name": "2-propanol",
        "Mw": 60.096001,
        "Tc": 508.29999,
        "Pc": 47.599998,
        "w": 0.66500002,
        "Tb": 355.39999,
        "m": 3.0929,
        "sigma": 3.2084999,
        "epsilon": 208.42,
        "k": 0.024675,
        "e": 2253.8999,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 28,
      "data": {
        "name": "2-propylamine (ISOPROPYL AMINE)",
        "Mw": 59.112,
        "Tc": 471.79999,
        "Pc": 45.400002,
        "w": 0.29100001,
        "Tb": 305.60001,
        "m": 2.5908,
        "sigma": 3.4777,
        "epsilon": 231.8,
        "k": 0.02134,
        "e": 932.20001,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 29,
      "data": {
        "name": "3-methyl pentane",
        "Mw": 86.178001,
        "Tc": 504.5,
        "Pc": 31.200001,
        "w": 0.27200001,
        "Tb": 336.39999,
        "m": 2.8852,
        "sigma": 3.8605001,
        "epsilon": 240.48,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 30,
      "data": {
        "name": "CO2",
        "Mw": 44.009998,
        "Tc": 304.10001,
        "Pc": 73.800003,
        "w": 0.23899999,
        "Tb": 200,
        "m": 2.0729001,
        "sigma": 2.7852001,
        "epsilon": 169.21001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 31,
      "data": {
        "name": "acetic acid",
        "Mw": 60.051998,
        "Tc": 592.70001,
        "Pc": 57.900002,
        "w": 0.447,
        "Tb": 391.10001,
        "m": 1.3403,
        "sigma": 3.8582001,
        "epsilon": 211.59,
        "k": 0.075549997,
        "e": 3044.3999,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 32,
      "data": {
        "name": "acetone (polar)",
        "Mw": 58.080002,
        "Tc": 508.10001,
        "Pc": 47,
        "w": 0.30399999,
        "Tb": 329.20001,
        "m": 2.221,
        "sigma": 3.607908,
        "epsilon": 259.98999,
        "k": 0,
        "e": 0,
        "d": 2.7,
        "x": 0.22579999
      }
    },
    {
      "id": 33,
      "data": {
        "name": "aniline",
        "Mw": 93.128998,
        "Tc": 699,
        "Pc": 53.099998,
        "w": 0.384,
        "Tb": 457.60001,
        "m": 2.6607001,
        "sigma": 3.7021,
        "epsilon": 335.47,
        "k": 0.074882999,
        "e": 1351.6,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 34,
      "data": {
        "name": "argon",
        "Mw": 39.948002,
        "Tc": 150.8,
        "Pc": 48.700001,
        "w": 0.001,
        "Tb": 87.300003,
        "m": 0.9285,
        "sigma": 3.4784,
        "epsilon": 122.23,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 35,
      "data": {
        "name": "benzene",
        "Mw": 78.113998,
        "Tc": 562.20001,
        "Pc": 48.900002,
        "w": 0.212,
        "Tb": 353.20001,
        "m": 2.4653001,
        "sigma": 3.6478,
        "epsilon": 287.35001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 36,
      "data": {
        "name": "biphenyl",
        "Mw": 154.21201,
        "Tc": 789,
        "Pc": 38.5,
        "w": 0.37200001,
        "Tb": 529.29999,
        "m": 3.8877001,
        "sigma": 3.8151,
        "epsilon": 327.42001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 37,
      "data": {
        "name": "bromobenzene",
        "Mw": 157.00999,
        "Tc": 670,
        "Pc": 45.200001,
        "w": 0.25099999,
        "Tb": 429.20001,
        "m": 2.6456001,
        "sigma": 3.836,
        "epsilon": 334.37,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 38,
      "data": {
        "name": "carbon disulfide",
        "Mw": 76.130997,
        "Tc": 552,
        "Pc": 79,
        "w": 0.109,
        "Tb": 319,
        "m": 1.6919,
        "sigma": 3.6171999,
        "epsilon": 334.82001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 39,
      "data": {
        "name": "carbon monoxide",
        "Mw": 28.01,
        "Tc": 132.89999,
        "Pc": 35,
        "w": 0.066,
        "Tb": 81.699997,
        "m": 1.3097,
        "sigma": 3.2507,
        "epsilon": 92.150002,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 40,
      "data": {
        "name": "chlorine",
        "Mw": 70.905998,
        "Tc": 416.89999,
        "Pc": 79.800003,
        "w": 0.090000004,
        "Tb": 239.2,
        "m": 1.5513999,
        "sigma": 3.3671999,
        "epsilon": 265.67001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 41,
      "data": {
        "name": "chlorobenzene",
        "Mw": 112.559,
        "Tc": 632.40002,
        "Pc": 45.200001,
        "w": 0.249,
        "Tb": 404.89999,
        "m": 2.6485,
        "sigma": 3.7533,
        "epsilon": 315.04001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 42,
      "data": {
        "name": "chloroethane",
        "Mw": 64.514999,
        "Tc": 460.39999,
        "Pc": 52.700001,
        "w": 0.191,
        "Tb": 285.5,
        "m": 2.2637999,
        "sigma": 3.4159999,
        "epsilon": 245.42999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 43,
      "data": {
        "name": "cyclohexane",
        "Mw": 84.162003,
        "Tc": 553.5,
        "Pc": 40.700001,
        "w": 0.212,
        "Tb": 353.79999,
        "m": 2.5302999,
        "sigma": 3.8499,
        "epsilon": 278.10999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 44,
      "data": {
        "name": "cyclopentane",
        "Mw": 70.135002,
        "Tc": 511.70001,
        "Pc": 45.099998,
        "w": 0.19599999,
        "Tb": 322.39999,
        "m": 2.3655,
        "sigma": 3.7114,
        "epsilon": 265.82999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 45,
      "data": {
        "name": "diethyl ether (polar)",
        "Mw": 74.123001,
        "Tc": 466.70001,
        "Pc": 36.400002,
        "w": 0.28099999,
        "Tb": 307.60001,
        "m": 2.8787,
        "sigma": 3.5548999,
        "epsilon": 220.59,
        "k": 0,
        "e": 0,
        "d": 1.2,
        "x": 0.34740001
      }
    },
    {
      "id": 46,
      "data": {
        "name": "diethyl ether",
        "Mw": 74.123001,
        "Tc": 466.70001,
        "Pc": 36.400002,
        "w": 0.28099999,
        "Tb": 307.60001,
        "m": 2.9686,
        "sigma": 3.5146999,
        "epsilon": 220.09,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 47,
      "data": {
        "name": "dimethyl ether (polar)",
        "Mw": 46.069,
        "Tc": 400,
        "Pc": 52.400002,
        "w": 0.2,
        "Tb": 248.3,
        "m": 2.0090001,
        "sigma": 3.4342999,
        "epsilon": 215.98,
        "k": 0,
        "e": 0,
        "d": 1.3,
        "x": 0.49770001
      }
    },
    {
      "id": 48,
      "data": {
        "name": "dimethyl ether",
        "Mw": 46.069,
        "Tc": 400,
        "Pc": 52.400002,
        "w": 0.2,
        "Tb": 248.3,
        "m": 2.3071001,
        "sigma": 3.2528,
        "epsilon": 211.06,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 49,
      "data": {
        "name": "dipropyl ether (polar)",
        "Mw": 102.177,
        "Tc": 530.59998,
        "Pc": 30.299999,
        "w": 0.36899999,
        "Tb": 363.20001,
        "m": 3.3929999,
        "sigma": 3.7425001,
        "epsilon": 236.19,
        "k": 0,
        "e": 0,
        "d": 1.2,
        "x": 0.2947
      }
    },
    {
      "id": 50,
      "data": {
        "name": "ethane",
        "Mw": 30.07,
        "Tc": 305.39999,
        "Pc": 48.799999,
        "w": 0.098999999,
        "Tb": 184.60001,
        "m": 1.6069,
        "sigma": 3.5206001,
        "epsilon": 191.42,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 51,
      "data": {
        "name": "ethanol (polar)",
        "Mw": 46.069,
        "Tc": 513.90002,
        "Pc": 61.400002,
        "w": 0.64399999,
        "Tb": 351.39999,
        "m": 2.2049,
        "sigma": 3.2774,
        "epsilon": 187.24001,
        "k": 0.033629999,
        "e": 2652.7,
        "d": 1.7,
        "x": 0.29466
      }
    },
    {
      "id": 52,
      "data": {
        "name": "ethanol",
        "Mw": 46.069,
        "Tc": 513.90002,
        "Pc": 61.400002,
        "w": 0.64399999,
        "Tb": 351.39999,
        "m": 2.3827,
        "sigma": 3.1770999,
        "epsilon": 198.24001,
        "k": 0.032384001,
        "e": 2653.3999,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 53,
      "data": {
        "name": "ethyl acetate (polar)",
        "Mw": 88.107002,
        "Tc": 523.20001,
        "Pc": 38.299999,
        "w": 0.36199999,
        "Tb": 350.29999,
        "m": 2.7481,
        "sigma": 3.6510999,
        "epsilon": 236.99001,
        "k": 0,
        "e": 0,
        "d": 1.84,
        "x": 0.54579997
      }
    },
    {
      "id": 54,
      "data": {
        "name": "ethyl amine",
        "Mw": 45.084999,
        "Tc": 456.39999,
        "Pc": 56.400002,
        "w": 0.289,
        "Tb": 289.70001,
        "m": 2.7046001,
        "sigma": 3.1343,
        "epsilon": 221.53,
        "k": 0.017275,
        "e": 854.70001,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 55,
      "data": {
        "name": "ethyl ethanoate (ETHYL ACETATE)",
        "Mw": 88.107002,
        "Tc": 523.20001,
        "Pc": 38.299999,
        "w": 0.36199999,
        "Tb": 350.29999,
        "m": 3.5374999,
        "sigma": 3.3079,
        "epsilon": 230.8,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 56,
      "data": {
        "name": "ethyl methanoate (ETHYL FORMATE)",
        "Mw": 74.080002,
        "Tc": 508.5,
        "Pc": 47.400002,
        "w": 0.285,
        "Tb": 327.5,
        "m": 2.8875999,
        "sigma": 3.3109,
        "epsilon": 246.47,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 57,
      "data": {
        "name": "ethyl propanoate (ETHYL PROPIONATE)",
        "Mw": 102.134,
        "Tc": 546,
        "Pc": 33.599998,
        "w": 0.391,
        "Tb": 372.20001,
        "m": 3.8371,
        "sigma": 3.4031,
        "epsilon": 232.78,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 58,
      "data": {
        "name": "ethylbenzene",
        "Mw": 106.168,
        "Tc": 617.20001,
        "Pc": 36,
        "w": 0.30199999,
        "Tb": 409.29999,
        "m": 3.0799,
        "sigma": 3.7974,
        "epsilon": 287.35001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 59,
      "data": {
        "name": "ethylcyclohexane",
        "Mw": 112.216,
        "Tc": 609,
        "Pc": 30,
        "w": 0.243,
        "Tb": 404.89999,
        "m": 2.8255999,
        "sigma": 4.1039,
        "epsilon": 294.04001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 60,
      "data": {
        "name": "ethylcyclopentane",
        "Mw": 98.189003,
        "Tc": 569.5,
        "Pc": 34,
        "w": 0.271,
        "Tb": 376.60001,
        "m": 2.9061999,
        "sigma": 3.8873,
        "epsilon": 270.5,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 61,
      "data": {
        "name": "ethylene",
        "Mw": 28.054001,
        "Tc": 282.39999,
        "Pc": 50.400002,
        "w": 0.089000002,
        "Tb": 169.3,
        "m": 1.5930001,
        "sigma": 3.4449999,
        "epsilon": 176.47,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 62,
      "data": {
        "name": "isobutane",
        "Mw": 58.124001,
        "Tc": 408.20001,
        "Pc": 36.5,
        "w": 0.183,
        "Tb": 261.39999,
        "m": 2.2616,
        "sigma": 3.7574,
        "epsilon": 216.53,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 63,
      "data": {
        "name": "isopentane",
        "Mw": 72.151001,
        "Tc": 460.39999,
        "Pc": 33.900002,
        "w": 0.227,
        "Tb": 301,
        "m": 2.562,
        "sigma": 3.8296001,
        "epsilon": 230.75,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 64,
      "data": {
        "name": "m-xylene",
        "Mw": 106.168,
        "Tc": 617.09998,
        "Pc": 35.400002,
        "w": 0.32499999,
        "Tb": 412.29999,
        "m": 3.1861,
        "sigma": 3.7563,
        "epsilon": 283.98001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 65,
      "data": {
        "name": "methane",
        "Mw": 16.042999,
        "Tc": 190.39999,
        "Pc": 46,
        "w": 0.011,
        "Tb": 111.6,
        "m": 1,
        "sigma": 3.7039001,
        "epsilon": 150.03,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 66,
      "data": {
        "name": "methanol (polar)",
        "Mw": 32.042,
        "Tc": 512.59998,
        "Pc": 80.900002,
        "w": 0.55599999,
        "Tb": 337.70001,
        "m": 1.7266001,
        "sigma": 3.1368999,
        "epsilon": 168.84,
        "k": 0.063110001,
        "e": 2585.8999,
        "d": 1.7,
        "x": 0.35128
      }
    },
    {
      "id": 67,
      "data": {
        "name": "methanol",
        "Mw": 32.042,
        "Tc": 512.59998,
        "Pc": 80.900002,
        "w": 0.55599999,
        "Tb": 337.70001,
        "m": 1.5255001,
        "sigma": 3.23,
        "epsilon": 188.89999,
        "k": 0.035176001,
        "e": 2899.5,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 68,
      "data": {
        "name": "methyl amine",
        "Mw": 31.058001,
        "Tc": 430,
        "Pc": 74.300003,
        "w": 0.292,
        "Tb": 266.79999,
        "m": 2.3966999,
        "sigma": 2.8906,
        "epsilon": 214.94,
        "k": 0.095103003,
        "e": 684.29999,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 69,
      "data": {
        "name": "methyl butanoate (METHYL BUTYRATE)",
        "Mw": 102.134,
        "Tc": 554.40002,
        "Pc": 34.799999,
        "w": 0.38,
        "Tb": 375.89999,
        "m": 3.6758001,
        "sigma": 3.4437001,
        "epsilon": 240.62,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 70,
      "data": {
        "name": "methyl chloride",
        "Mw": 50.487999,
        "Tc": 416.29999,
        "Pc": 67,
        "w": 0.153,
        "Tb": 249.10001,
        "m": 1.9297,
        "sigma": 3.2293,
        "epsilon": 240.56,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 71,
      "data": {
        "name": "methyl methanoate (METHYL FORMATE)",
        "Mw": 60.051998,
        "Tc": 487.20001,
        "Pc": 60,
        "w": 0.257,
        "Tb": 304.89999,
        "m": 2.6784,
        "sigma": 3.0875001,
        "epsilon": 242.63,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 72,
      "data": {
        "name": "methyl propanoate (METHYL PROPIONATE)",
        "Mw": 88.107002,
        "Tc": 530.59998,
        "Pc": 40,
        "w": 0.34999999,
        "Tb": 352.79999,
        "m": 3.4793,
        "sigma": 3.3141999,
        "epsilon": 234.96001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 73,
      "data": {
        "name": "methyl-n-propyl ether",
        "Mw": 74.123001,
        "Tc": 476.29999,
        "Pc": 38,
        "w": 0.271,
        "Tb": 311.70001,
        "m": 3.0086999,
        "sigma": 3.4568999,
        "epsilon": 222.73,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 74,
      "data": {
        "name": "methylcyclohexane",
        "Mw": 98.189003,
        "Tc": 572.20001,
        "Pc": 34.700001,
        "w": 0.236,
        "Tb": 374.10001,
        "m": 2.6637001,
        "sigma": 3.9993,
        "epsilon": 282.32999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 75,
      "data": {
        "name": "methylcyclopentane",
        "Mw": 84.162003,
        "Tc": 532.70001,
        "Pc": 37.799999,
        "w": 0.23100001,
        "Tb": 345,
        "m": 2.6129999,
        "sigma": 3.8253,
        "epsilon": 265.12,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 76,
      "data": {
        "name": "n-butane",
        "Mw": 58.124001,
        "Tc": 425.10001,
        "Pc": 37.959999,
        "w": 0.2,
        "Tb": 272.70001,
        "m": 2.3316,
        "sigma": 3.7086,
        "epsilon": 222.88,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 77,
      "data": {
        "name": "n-butyl acetate (polar)",
        "Mw": 116.16,
        "Tc": 579,
        "Pc": 31.4,
        "w": 0.417,
        "Tb": 399.29999,
        "m": 3.6600001,
        "sigma": 3.6751001,
        "epsilon": 237.42999,
        "k": 0,
        "e": 0,
        "d": 1.86,
        "x": 0.40979999
      }
    },
    {
      "id": 78,
      "data": {
        "name": "n-butyl ethanoate (N-BUTYL ACETATE)",
        "Mw": 116.16,
        "Tc": 579,
        "Pc": 31.4,
        "w": 0.417,
        "Tb": 399.29999,
        "m": 3.9807999,
        "sigma": 3.5427001,
        "epsilon": 242.52,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 79,
      "data": {
        "name": "n-butylbenzene",
        "Mw": 134.222,
        "Tc": 660.5,
        "Pc": 28.9,
        "w": 0.39300001,
        "Tb": 456.5,
        "m": 3.7662001,
        "sigma": 3.8727,
        "epsilon": 283.07001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 80,
      "data": {
        "name": "n-decane",
        "Mw": 142.26801,
        "Tc": 617.70001,
        "Pc": 21.200001,
        "w": 0.48899999,
        "Tb": 447.29999,
        "m": 4.6627002,
        "sigma": 3.8383999,
        "epsilon": 243.87,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 81,
      "data": {
        "name": "n-dodecane",
        "Mw": 170.34,
        "Tc": 658.20001,
        "Pc": 18.200001,
        "w": 0.57499999,
        "Tb": 489.5,
        "m": 5.3060002,
        "sigma": 3.8959,
        "epsilon": 249.21001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 82,
      "data": {
        "name": "n-eicosane",
        "Mw": 282.556,
        "Tc": 767,
        "Pc": 11.1,
        "w": 0.90700001,
        "Tb": 617,
        "m": 7.9849,
        "sigma": 3.9869001,
        "epsilon": 257.75,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 83,
      "data": {
        "name": "n-heptadecane",
        "Mw": 240.47501,
        "Tc": 733,
        "Pc": 13,
        "w": 0.76999998,
        "Tb": 575.20001,
        "m": 6.9808998,
        "sigma": 3.9675,
        "epsilon": 255.64999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 84,
      "data": {
        "name": "n-heptane",
        "Mw": 100.205,
        "Tc": 540.29999,
        "Pc": 27.4,
        "w": 0.34900001,
        "Tb": 371.60001,
        "m": 3.4830999,
        "sigma": 3.8048999,
        "epsilon": 238.39999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 85,
      "data": {
        "name": "n-hexadecane",
        "Mw": 226.448,
        "Tc": 722,
        "Pc": 14.1,
        "w": 0.74199998,
        "Tb": 560,
        "m": 6.6485,
        "sigma": 3.9552,
        "epsilon": 254.7,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 86,
      "data": {
        "name": "n-hexane",
        "Mw": 86.178001,
        "Tc": 507.5,
        "Pc": 30.1,
        "w": 0.29899999,
        "Tb": 341.89999,
        "m": 3.0576,
        "sigma": 3.7983,
        "epsilon": 236.77,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 87,
      "data": {
        "name": "n-nonadecane",
        "Mw": 268.52899,
        "Tc": 756,
        "Pc": 11.1,
        "w": 0.82700002,
        "Tb": 603.09998,
        "m": 7.7175002,
        "sigma": 3.9721,
        "epsilon": 256,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 88,
      "data": {
        "name": "n-nonane",
        "Mw": 128.259,
        "Tc": 594.59998,
        "Pc": 22.9,
        "w": 0.44499999,
        "Tb": 424,
        "m": 4.2079,
        "sigma": 3.8448,
        "epsilon": 244.50999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 89,
      "data": {
        "name": "n-octadecane",
        "Mw": 254.504,
        "Tc": 748,
        "Pc": 12,
        "w": 0.79000002,
        "Tb": 589.5,
        "m": 7.3270998,
        "sigma": 3.9668,
        "epsilon": 256.20001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 90,
      "data": {
        "name": "n-octane",
        "Mw": 114.232,
        "Tc": 568.79999,
        "Pc": 24.9,
        "w": 0.398,
        "Tb": 398.79999,
        "m": 3.8176,
        "sigma": 3.8373001,
        "epsilon": 242.78,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 91,
      "data": {
        "name": "n-pentadecane",
        "Mw": 212.42101,
        "Tc": 707,
        "Pc": 15.2,
        "w": 0.70599997,
        "Tb": 543.79999,
        "m": 6.2855,
        "sigma": 3.9531,
        "epsilon": 254.14,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 92,
      "data": {
        "name": "n-pentane",
        "Mw": 72.151001,
        "Tc": 469.70001,
        "Pc": 33.700001,
        "w": 0.25099999,
        "Tb": 309.20001,
        "m": 2.6896,
        "sigma": 3.7729001,
        "epsilon": 231.2,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 93,
      "data": {
        "name": "n-propyl ethanoate (N-PROPYL ACETATE)",
        "Mw": 102.134,
        "Tc": 549.40002,
        "Pc": 33.299999,
        "w": 0.391,
        "Tb": 374.70001,
        "m": 3.7860999,
        "sigma": 3.4226999,
        "epsilon": 235.75999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 94,
      "data": {
        "name": "n-propyl methanoate (N-PROPYL FORMATE)",
        "Mw": 88.107002,
        "Tc": 538,
        "Pc": 40.599998,
        "w": 0.31400001,
        "Tb": 354.10001,
        "m": 3.2088001,
        "sigma": 3.4168,
        "epsilon": 246.46001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 95,
      "data": {
        "name": "n-propyl propanoate (N-PROPYL PROPIONATE)",
        "Mw": 116.16,
        "Tc": 571,
        "Pc": 30.200001,
        "w": 9e+30,
        "Tb": 395.79999,
        "m": 4.1155,
        "sigma": 3.4875,
        "epsilon": 235.60001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 96,
      "data": {
        "name": "n-propylbenzene",
        "Mw": 120.195,
        "Tc": 638.20001,
        "Pc": 32,
        "w": 0.34400001,
        "Tb": 432.39999,
        "m": 3.3438001,
        "sigma": 3.8438001,
        "epsilon": 288.13,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 97,
      "data": {
        "name": "n-tetradecane",
        "Mw": 198.93401,
        "Tc": 693,
        "Pc": 14.4,
        "w": 0.58099997,
        "Tb": 526.70001,
        "m": 5.9001999,
        "sigma": 3.9396,
        "epsilon": 254.21001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 98,
      "data": {
        "name": "n-tridecane",
        "Mw": 184.367,
        "Tc": 676,
        "Pc": 17.200001,
        "w": 0.61900002,
        "Tb": 508.60001,
        "m": 5.6876998,
        "sigma": 3.9143,
        "epsilon": 249.78,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 99,
      "data": {
        "name": "n-undecane",
        "Mw": 156.313,
        "Tc": 638.79999,
        "Pc": 19.700001,
        "w": 0.53500003,
        "Tb": 469.10001,
        "m": 4.9081998,
        "sigma": 3.8893001,
        "epsilon": 248.82001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 100,
      "data": {
        "name": "neopentane",
        "Mw": 72.151001,
        "Tc": 433.79999,
        "Pc": 32,
        "w": 0.197,
        "Tb": 282.60001,
        "m": 2.3543,
        "sigma": 3.9549999,
        "epsilon": 225.69,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 101,
      "data": {
        "name": "nitrogen",
        "Mw": 28.013,
        "Tc": 126.2,
        "Pc": 33.900002,
        "w": 0.039000001,
        "Tb": 77.400002,
        "m": 1.2053,
        "sigma": 3.313,
        "epsilon": 90.959999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 102,
      "data": {
        "name": "o-xylene",
        "Mw": 106.168,
        "Tc": 630.29999,
        "Pc": 37.299999,
        "w": 0.31,
        "Tb": 417.60001,
        "m": 3.1362,
        "sigma": 3.76,
        "epsilon": 291.04999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 103,
      "data": {
        "name": "p-xylene",
        "Mw": 106.168,
        "Tc": 616.20001,
        "Pc": 35.099998,
        "w": 0.31999999,
        "Tb": 411.5,
        "m": 3.1723001,
        "sigma": 3.7781,
        "epsilon": 283.76999,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 104,
      "data": {
        "name": "propane",
        "Mw": 44.094002,
        "Tc": 369.79999,
        "Pc": 42.5,
        "w": 0.153,
        "Tb": 231.10001,
        "m": 2.0020001,
        "sigma": 3.6184001,
        "epsilon": 208.11,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 105,
      "data": {
        "name": "propylene",
        "Mw": 42.081001,
        "Tc": 364.89999,
        "Pc": 46,
        "w": 0.14399999,
        "Tb": 225.5,
        "m": 1.9597,
        "sigma": 3.5355999,
        "epsilon": 207.19,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 106,
      "data": {
        "name": "sulfur dioxide",
        "Mw": 64.063004,
        "Tc": 430.79999,
        "Pc": 78.800003,
        "w": 0.25600001,
        "Tb": 263.20001,
        "m": 2.8611,
        "sigma": 2.6826,
        "epsilon": 205.35001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 107,
      "data": {
        "name": "tetralin",
        "Mw": 132.20599,
        "Tc": 719,
        "Pc": 35.099998,
        "w": 0.303,
        "Tb": 480.70001,
        "m": 3.3131001,
        "sigma": 3.875,
        "epsilon": 325.07001,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 108,
      "data": {
        "name": "toluene",
        "Mw": 92.140999,
        "Tc": 591.79999,
        "Pc": 41,
        "w": 0.26300001,
        "Tb": 383.79999,
        "m": 2.8148999,
        "sigma": 3.7169001,
        "epsilon": 285.69,
        "k": 0,
        "e": 0,
        "d": 0,
        "x": 0
      }
    },
    {
      "id": 109,
      "data": {
        "name": "water (polar)",
        "Mw": 18.014999,
        "Tc": 647.29999,
        "Pc": 221.2,
        "w": 0.34400001,
        "Tb": 373.14999,
        "m": 1.0405,
        "sigma": 2.9656999,
        "epsilon": 175.14999,
        "k": 0.08924,
        "e": 2706.7,
        "d": 1.85,
        "x": 0.66245002
      }
    },
    {
      "id": 110,
      "data": {
        "name": "water",
        "Mw": 18.014999,
        "Tc": 647.29999,
        "Pc": 221.2,
        "w": 0.34400001,
        "Tb": 373.14999,
        "m": 1.0656,
        "sigma": 3.0007,
        "epsilon": 366.51001,
        "k": 0.034867998,
        "e": 2500.7,
        "d": 0,
        "x": 0
      }
    }
  ]
}`