func isValidSudoku(board [][]byte) bool {
    columnMapList := make([]map[string]bool, len(board))
    subBoxMapList := make(map[string]map[string]bool, len(board))
    // lastRowIdx := len(board) - 1

    for i, row := range board {
        rowMap := make(map[string]bool)
        fmt.Printf("\n\n [%d] row: %s \n", i, row)
        
        for j, v := range row {
            sbNum := subBoxNum(i, j)
            if columnMapList[j] == nil {
                columnMapList[j] = make(map[string]bool, len(board))
            }
            if subBoxMapList[sbNum] == nil {
                subBoxMapList[sbNum] = make(map[string]bool, len(board))
            }
            
            if v == '.' {
                continue
            }
            fmt.Printf("-- i:%d | j:%d | num:%d - v:%s \n", i, j, sbNum, string(v))

            isExist := rowMap[string(v)]
            if isExist {
                return false
            }
            rowMap[string(v)] = true

            colMap := columnMapList[j]
            isExist = colMap[string(v)]
            if isExist {
                return false
            }
            colMap[string(v)] = true
            columnMapList[j] = colMap

            sbMap := subBoxMapList[sbNum]
            isExist = sbMap[string(v)]
            if isExist {
                return false
            }
            sbMap[string(v)] = true
            subBoxMapList[sbNum] = sbMap

            fmt.Printf("[%d] row_item: %s(%T) \n row_map: %+v \n col_map: %+v \n sub_box_map[%d]: %+v\n", 
                j, string(v), v, rowMap, colMap, sbNum ,sbMap)

            // if i == lastRowIdx {
            //     // since `Each column must contain the digits 1-9`
            //     if len(columnMapList[j]) == 0 {
            //         return false
            //     }
            // }
        }

        // since `Each row must contain the digits 1-9`
        // if len(rowMap) == 0{
        //     return false
        // }
    }

    fmt.Printf("\n\n --------------------- \n")
    for i, v := range columnMapList {
        fmt.Printf("[%d] colM: %+v\n", i, v)
    }
    fmt.Printf("\n\n --------------------- \n")
    for i, v := range subBoxMapList {
        fmt.Printf("[%s] subBox: %+v\n", i, v)
    }
    return true
}

// subBoxNum to identify its number based on row & columnd index
// subBox1: row 1-3 | col 1-3
// subBox2: row 1-3 | col 4-6
// subBox3: row 1-3 | col 7-9
// subBox4: row 4-6 | col 1-3
func subBoxNum(row, column int) string {
    return fmt.Sprintf("%d-%d", row/3, column/3)
}
