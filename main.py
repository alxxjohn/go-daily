import typing


class Movie(typing.TypedDict):
    name: str
    year: int


class ActionMovies(typing.TypedDict):
    pass


die_hard = Movie(Name="diehard", year=1998)
print(die_hard)