# api-gym
A place to workout your api idea. We use `examples/future_society.json` as a way to demo this tool.

# Future Society
Picture 12 to 30 people living together in a big house or some sort of living space. Could be a very large aparment, or farm, etc.

The people in these shared living spaces are either:

1. Just an individual with no kids.
2. In a partnership with another indivdual but raising no kids together.
3. An individual with N kids.
4. A partnership with two people, raising N kids.
5. A partnership with three people, raising N kids.

You can chose to start your own `living_space` (LS) and be the founding member or members.

Or you can chose to apply to an existing one.

Each LS can explain it's desire combination of people, parterships, and kids. Some might be 100% kid free. Some might be all families. Some might be in between with young partners thinking about having or adopting their first child.

Similar to [https://www.pacaso.com/learn](https://www.pacaso.com/learn) but not a 2nd home, your primary home, and you own a % of it in an LLC.

Existing apartment/condo buildings could be converted into having very large (think 10 bedrooms, 4 bathrooms) units by knocking down walls and combining units. Someone with the cash can start a conversion like this, and then post the option to buy into it with other individuals that meet the set criteria.

# The structs and routes

```
 0.  GET /humans/                  []Human
 1.  GET /living_spaces/           []LivingSpace
 2. POST /humans/
 3. POST /living_spaces/
 4. POST /living_spaces/apply

 0. LivingSpace
   0. Id                             string                         uuid
   1. Name                           string                         color
   2. Address                        string                         address
   3. Latitude                       float64                        latitude
   4. Longitude                      float64                        longitude
   5. HumanCount                     int                            small_int_range
   6. Humans                         []Human
   7. Familes                        []Family
   8. Partners                       []Partner
 1. Human
   0. Id                             string                         uuid
   1. Name                           string                         name
   2. Pronouns                       string                         pronouns
 2. Family
   0. Id                             string                         uuid
   1. Name                           string                         color
```


