function calculateDiagonalSum(param) {
    var n = param;
    var length = (n * n) + 1;
    var islands = Array.from({ length: n }, () => Array.from({ length: n }).fill(0));
    var doSwitch = false;
    var row = 0, column = n;
    var loopCount = n;
    do {
        if (doSwitch) {
            if (loopCount > 0) {
                // loop row
                for (var i = 0; i < loopCount; i++) {
                    length -= 1;
                    column += 1;
                    islands[row][column] = length;
                }

                loopCount = loopCount - 1;
                // loop column
                for (var i = 0; i < loopCount; i++) {
                    length -= 1;
                    row -= 1;
                    islands[row][column] = length;
                }
            }
        } else {
            if (loopCount > 0) {
                // loop row
                for (var i = 0; i < loopCount; i++) {
                    length -= 1;
                    column -= 1;
                    islands[row][column] = length;
                }

                loopCount = loopCount - 1;
                // loop column
                for (var i = 0; i < loopCount; i++) {
                    length -= 1;
                    row += 1;
                    islands[row][column] = length;
                }
            }
        }
        doSwitch = !doSwitch;
    } while (loopCount > 0);

    //top left - bottom right
    var count = 0;
    for (var i = 0; i < n; i++) {
        count += islands[i][i];
    }

    //top right to botton left
    for (var i = n-1; i >= 0; i--) {
        if (islands[(n-1)-i][i] != 1) count += islands[(n-1)-i][i];
    }

    console.log(count)   
}

calculateDiagonalSum(3)
calculateDiagonalSum(5)
calculateDiagonalSum(7)
