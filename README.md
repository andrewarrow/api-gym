# api-gym
A place to workout your api idea. Quickly define models with attributes and see real json.

# example run

```
os/api-gym $ ./api-gym new

New Gym Create with name: 5b6493c4-1732-a95e-5dca-654099e89594
```

After running `new` you get a new gym playground with `user` as your first model.

```
os/api-gym $ ./api-gym json user | jq .
{
  "users": [
    {
      "id": "9258569c-1f3d-6ee1-7d5d-6ac545ccb203",
      "first_name": "Katelin",
      "last_name": "Gaylord",
      "pronouns": "they/them",
      "email": "yesseniakonopelski@bogisich.info",
      "phone": "208-526-0479",
      "age": 28
    }
  ]
}
```

You can add new models:

```
os/api-gym $ ./api-gym add home

┌────────────────────┐┌────────────────────┐ ┌─────────────────────────────────────────────────────────────────────────────────────────────┐
│individual          ││id                  │ │ea8b0f11-440b-776c-a554-7577fa15ecfe                                                         │
│location            ││address             │ │6339 Lake Highway burgh, Oakland, Nevada 98594 Guinea-Bissau                                 │
│phone               ││latitude            │ │73.161558                                                                                    │
│timestamp           ││longitude           │ │52.943533                                                                                    │
│few_words           ││launched_at         │ │2022-04-18T06:52:29.940Z                                                                     │
│int                 ││name                │ │unless but                                                                                   │
│float               ││bedrooms            │ │26253                                                                                        │
│bool                ││square_feet         │ │36307                                                                                        │
│paragraph           ││summary             │ │Doloremque quia optio laudantium rem est maiores laudantium ea nihil doloribus culpa id numq…│
│email               ││bathrooms           │ │9.4                                                                                          │
│uuid                ││                    │ │                                                                                             │
│model               ││                    │ │                                                                                             │
│[]model             ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
└────────────────────┘└────────────────────┘ └─────────────────────────────────────────────────────────────────────────────────────────────┘
```

And then see their json:

```
os/api-gym $ ./api-gym json home | jq .
{
  "homes": [
    {
      "id": "69682aa4-4d04-c5d0-04c1-33baf41a43a3",
      "address": "86676 West Prairie stad, Lincoln, New Hampshire 55393 Aruba",
      "latitude": -88.457909,
      "longitude": 151.268922,
      "launched_at": "2022-04-18T06:52:29.940Z",
      "name": "away front",
      "bedrooms": 15,
      "square_feet": 12777,
      "summary": "Non vel voluptatem sequi sunt sequi voluptatem unde expedita nihil voluptatem sunt laboriosam consectetur perferendis qui est ea quae nihil odit et et ut sapiente qui cupiditate nobis consequatur nostrum culpa aut consequatur. Iusto sapiente voluptatum provident non id modi asperiores et blanditiis autem facilis consequatur explicabo illum nulla saepe saepe quia corporis sit autem et est ipsa est enim fugiat vero explicabo natus quia quisquam. Id qui fuga nobis nesciunt vitae perferendis ea in quis voluptas dolores libero velit veritatis ea modi tempora voluptatem enim officia maxime fugiat sunt deserunt architecto eos quidem excepturi ad repudiandae est laborum.",
      "bathrooms": 7.5
    }
  ]
}
```
