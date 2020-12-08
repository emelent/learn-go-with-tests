package reflection

type WalkedFunc func(string)

func Walk(x interface{}, fn WalkedFunc) {
}
