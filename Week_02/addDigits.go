func addDigits(num int) int {
    if num>=10{
        return addDigits(num/10+num%10)
    }else{
        return num
    }
}
