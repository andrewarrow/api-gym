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

Each LS can explain its desired combination of people, parterships, and kids. Some might be 100% kid free. Some might be all families. Some might be in between with young partners thinking about having or adopting their first child.

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
   5. BedroomCount                   int
   5. BathCount                      float64
   5. SquareFeet                     int
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

Let's create the JSON for a sample existing LS in Culver City, CA. The name of the LS will be "Culver Families for Love." It's a condo conversion from a few years ago, eight 2 bedrooms, 1 baths were merged. The total was 16 bedrooms, 8 baths but a few baths were removed to make more bedrooms. The final count was 20 bedrooms and 4.5 baths in 11,200 square feet. Picture a very large communal kitchen, play area for kids, and a large outdoor deck. All 11,200 square feet was repurposed to make sense for 30 people that will be living and working in this space.

The desired 30 people breaks down like:

1. Two families of six.
2. Two families of three.
3. Four partnerships with no kids.
4. Three individuals with no partners no kids.
5. Two individuals with one kid each.

12+6+4+3+4 = 29 with one empty spot for a future kid.



