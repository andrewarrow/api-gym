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
