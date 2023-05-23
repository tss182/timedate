# Install
    go get github.com/tss182/timedate

# How to using

    package main
    
    import (
    	"fmt"
    	"github.com/tss182/timedate"
    )
    
    func main()  {
        startDate,endDate := timedate.GetStartEndDate(time.Now(),date.Month)
        fmt.Println(startDate,endDate)
    }
