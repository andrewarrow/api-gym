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
      "age": "8"
    }
  ]
}
```

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
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
│                    ││                    │ │                                                                                             │
└────────────────────┘└────────────────────┘ └─────────────────────────────────────────────────────────────────────────────────────────────┘
```
