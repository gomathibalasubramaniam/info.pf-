package main
import ("fmt")
func main(){
    var finalFee float64
    var Fees float64
    var fee float64
    var mark float64=70.0
    var courseFee float64=25000.0 
    var extraFee float64=1500.0 
    var scholarship_Percent float64
    scholarship_Percent=(mark/200.0)
    fee=(scholarship_Percent*courseFee)
    Fees=courseFee-fee
    finalFee=Fees+extraFee
    
    //Write your program logic here
    //Populate the variable: finalFee


     //Do not modify the below print statement for verification to work
     fmt.Println(finalFee) 
}
