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

You can also add a nested single model like:

```
{
  "homes": [
    {
      "id": "5b1ee03a-3585-0b0c-fa31-fa2329210c07",
      "address": "7296 Port Tunnel furt, Norfolk, Wisconsin 39744 Heard Island and McDonald Islands",
      "latitude": "4.982696",
      "longitude": "-72.510182",
      "launched_at": "2022-04-18T06:52:29.940Z",
      "name": "which several",
      "bedrooms": 62206,
      "square_feet": 6849,
      "summary": "Aut eaque velit quidem et tempora rerum iste a voluptas quia quasi atque unde et provident inventore molestiae veritatis officia ducimus esse non non perspiciatis soluta repellat qui corrupti rem ut vel at. Qui aliquid optio asperiores est saepe aut omnis sint voluptates deserunt harum veritatis eos tempore voluptas consequuntur ex reprehenderit voluptatem dolor delectus ad quia aut enim est aut tenetur itaque saepe ea ex. Et ea eos et pariatur itaque autem iste quaerat aut odit voluptatem occaecati quaerat unde corrupti magnam labore quibusdam voluptas quo facere aut ipsam aliquam eum esse est cumque ducimus autem sapiente iure.",
      "bathrooms": 28.2,
      "created_by": {
        "id": "2ca93ee3-060a-fe70-3702-81fbf6bfe16a",
        "first_name": "Ewell",
        "last_name": "Batz",
        "pronouns": "she/her",
        "email": "kathleenkoss@hirthe.biz",
        "phone": "779.535.5025",
        "age": 27
      }
    }
  ]
}
```

Or array of models:

```
{
  "homes": [
    {
      "id": "569be8f2-1a0d-e308-0502-a815492671af",
      "address": "779 Lake Locks mouth, El Paso, Nevada 55957 Zambia",
      "latitude": "34.956466",
      "longitude": "-151.279102",
      "launched_at": "2022-04-18T06:52:29.940Z",
      "name": "mine then",
      "bedrooms": 24721,
      "square_feet": 46175,
      "summary": "Aperiam et ut rerum fugit labore sed explicabo rerum in repellendus voluptatem saepe aperiam ea ab aliquid rerum dolorem quam soluta sit ex distinctio perferendis quis fugiat voluptatem officiis in est commodi odit. Corrupti in pariatur voluptatum esse sit voluptatem nesciunt nihil est commodi aliquid et sed optio ut a ut perspiciatis ut mollitia quo et voluptate dolorum sint ut cum quis provident ullam voluptas sit. Aliquid eveniet suscipit et itaque voluptas minima ea omnis ea at velit odio architecto dicta rem id magnam est eaque culpa inventore reiciendis pariatur odio non sed quis minus saepe perspiciatis voluptas ipsam.",
      "bathrooms": 20.5,
      "created_by": {
        "id": "36e3f11c-ae34-8e84-aa90-7c62ca26914a",
        "first_name": "Alysson",
        "last_name": "Morar",
        "pronouns": "they/them",
        "email": "lillianbeier@jaskolski.com",
        "phone": "792-518-2692",
        "age": 11
      },
      "admins": [
        {
          "id": "cd870417-7328-9fe6-ffea-c597aaf3549a",
          "first_name": "Jessie",
          "last_name": "Hagenes",
          "pronouns": "they/them",
          "email": "lydamorar@stracke.com",
          "phone": "241.899.9102",
          "age": 22
        },
        {
          "id": "c29df12c-d636-33f0-05d9-6e90bec01bcd",
          "first_name": "Jammie",
          "last_name": "Hickle",
          "pronouns": "he/him",
          "email": "randallturcotte@king.net",
          "phone": "697-138-9817",
          "age": 60
        }
      ]
    }
  ]
}
```
