func isAnagram(s string, t string) bool {
    if len(s) != len(t){
        return false
    }

    sMap := make(map[rune]int, len(s))
    tMap := make(map[rune]int, len(s))
    for _, v := range s {
        // tempCount := sMap[v]
        sMap[v]++

        // tempCount2 := tMap[t[idx]]
        // tMap[t[idx]] = tempCount2
        // tMap[t[idx]]++
    }

     for _, v := range t {
        tMap[v]++
     }

    // for _, v2 := range t {
    //     _, isExist := sMap[v2]
    //     if !isExist {
    //         return false
    //     }
    // }

    for k, v := range sMap {
        if v != tMap[k] {
            return false
        } 
    }

    return true
}
