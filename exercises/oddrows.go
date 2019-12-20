package exercises


//RowSumOddNumbers ...
func RowSumOddNumbers(n int) int {
    sum:= 0;
    start := n*(n+1)/2;
    start--;
    for i:=0; i<n; i++ {
      sum+= 1+(2*start)
      start--
    }
    
    return sum;
}
