func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    sMap := make(map[rune]int, len(s))
    
    for _, v := range s {
        sMap[v]++
    }

    for _, v := range t {
        sMap[v]--
    }

    for _, v := range sMap {
        if v != 0 {
            return false
        } 
    }

    return true
}
