package stats

import (
	"fmt"

	"github.com/ShavqatKavrakov/Lesson10_bank_1/pkg/bank/types"
)

func ExampleAvg() {
	fmt.Println(Avg([]types.Payment{
     {
      Amount: 1_000_00,
	 },		
	 {
		 Amount: 2_000_00,
	 },
	 {
		 Amount: 12_000_00,
	 },
	}))
	fmt.Println(Avg([]types.Payment{	
		{
			Amount: 12_000_00,
		},
	   }))
	   
	   //Output:
	   //500000
	   //1200000
}