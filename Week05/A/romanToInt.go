func romanToInt(s string) int {
    romanValues := map[byte]int{'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}

    n := len(s)
    var retValue int

    for i:= range s {
        value := romanValues[s[i]]
        if i < n-1 && value < romanValues[s[i+1]] {
            retValue -= value
        } else {
            retValue += value
        }
    }

    return retValue

}