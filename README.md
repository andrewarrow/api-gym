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
    Id                             string                         uuid
    Name                           string                         color
    Mission                        string                         
    Address                        string                         address
    Latitude                       float64                        latitude
    Longitude                      float64                        longitude
    BedroomCount                   int
    BathCount                      float64
    SquareFeet                     int
    HumanCount                     int                            small_int_range
    Humans                         []Human
    Familes                        []Family
    Partners                       []Partner
 1. Human
    Id                             string                         uuid
    Name                           string                         name
    Pronouns                       string                         pronouns
 2. Family
    Id                             string                         uuid
    Name                           string                         color
```

Let's create the JSON for two sample existing LS's in Culver City, CA. The name of the first LS will be "Culver Families for Love." It's a condo conversion from a few years ago, eight 2 bedrooms, 1 baths were merged. The total was 16 bedrooms, 8 baths but a few baths were removed to make more bedrooms. The final count was 20 bedrooms and 4.5 baths in 11,200 square feet. Picture a very large communal kitchen, play area for kids, and a large outdoor deck. All 11,200 square feet was repurposed to make sense for 30 people that will be living and working in this space. This was the entire first floor of the building.

There was another 11,200 of square feet in the same building (2nd floor) also converted and will be the second example LS. The two of them together share the one pool, garage, mail room, etc.

The desired 30 people for "Culver Families for Love" breaks down like:

1. Two families of six.
2. Two families of three.
3. Four partnerships with no kids.
4. Three individuals with no partners no kids.
5. Two individuals with one kid each.

12+6+4+3+4 = 29 with one empty spot for a future kid.

The `mission` statement of this LS is: "Live together with love for everyone. We allow all types of food in the kichen but have many vegans. We allow for all types of gender identification. We believe in strong values for our kids and all the adults help all the kids. Our adults have a least one member working in the tech industry. Areas of the home are designed for their work with network access and no kids areas during working hours."

The second example LS's name is "Culver Looking" and it breaks down like:

1. Eight partnerships with no kids.
2. Eight individuals with no partners no kids.

16+8 = 24 with six open spots for future kids or individuals.

The `mission` statement of this LS is: "No kids yet but looking for a family oriented life someday. Working professional tend to be in the film and television business but all professional are welcome. Dating within our LS is allowed but we value respecting existing partnerships and not disrupting them with bad intent. We welcome the first of our community to have children and spots are open for them."

```
{
  "data": [
    {
      "id": "c40fde42-a950-50c5-5e33-68f61188e31f",
      "name": "Culver Families for Love",
      "mission": "Live together with love for everyone. We allow all types of food in the kichen but have many vegans. We allow for all types of gender identification. We believe in strong values for our kids and all the adults help all the kids. Our adults have a least one member working in the tech industry. Areas of the home are designed for their work with network access and no kids areas during working hours.",
      "address": "6410 Culver Blvd Floor #1, Culver City, CA 90230 USA",
      "latitude": 23.531596,
      "longitude": -54.374458,
      "human_count": 29,
      "bedroom_count": 20,
      "bath_count": 4.5,
      "square_feet": 11200
    },
    {
      "id": "aa603ada-c537-6641-37a8-a419163c9d87",
      "name": "Culver Looking",
      "mission": "No kids yet but looking for a family oriented life someday. Working professional tend to be in the film and television business but all professional are welcome. Dating within our LS is allowed but we value respecting existing partnerships and not disrupting them with bad intent. We welcome the first of our community to have children and spots are open for them.",
      "address": "6410 Culver Blvd Floor #2, Culver City, CA 90230 USA",
      "latitude": 23.531596,
      "longitude": -54.374458,
      "human_count": 24,
      "bedroom_count": 20,
      "bath_count": 4.5,
      "square_feet": 11200
    }
  ]
}

```
