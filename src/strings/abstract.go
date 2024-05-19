package strings

type StirngSearch interface {
    GetName() string
    FindPattern(s string) []int
}

