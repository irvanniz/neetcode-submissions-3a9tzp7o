func isValidSudoku(board [][]byte) bool {
    columnMapList := make([]map[string]bool, len(board))
    subBoxMapList := make(map[string]map[string]bool, len(board))

    for i, row := range board {
        rowMap := make(map[string]bool)
        
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

            val := string(v)
            isExist := rowMap[val]
            if isExist {
                return false
            }
            rowMap[val] = true

            colMap := columnMapList[j]
            isExist = colMap[val]
            if isExist {
                return false
            }
            colMap[val] = true
            columnMapList[j] = colMap

            sbMap := subBoxMapList[sbNum]
            isExist = sbMap[val]
            if isExist {
                return false
            }
            sbMap[val] = true
            subBoxMapList[sbNum] = sbMap
        }
    }

    return true
}

func subBoxNum(row, column int) string {
    return fmt.Sprintf("%d-%d", row/3, column/3)
}
