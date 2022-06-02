package types

//Money  представляет собой денежную единицу в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

//Category представляет собой категорию, в каторый был совершён платёж (авто, аптеки, рестораны и.т)
type Category string

//Payment представляет информацию о платёже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
