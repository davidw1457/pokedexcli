package pokeapi

type Pokemon struct {
    Name string `json:"name"`
    BaseExperience int `json:"base_experience"`
    Height int `json:"height"`
    Types []struct{
        Type struct{
            Name string `json:"name"`
        } `json:"type"`
    } `json:"types"`
    Weight int `json:"weight"`
    Stats []struct{
        BaseStat int `json:"base_stat"`
        Stat struct{
            Name string `json:"name"`
        } `json:"stat"`
    } `json:"stats"`
}
