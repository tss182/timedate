# Install
    go get github.com/tss182/date

# How to using

    package main
    
    import (
    	"fmt"
    	"github.com/tss182/date"
    )
    
    func main()  {
    	startDate,endDate := date.GetStartEndDate(time.Now(),"month")
        fmt.Println(startDate,endDate)
    }
