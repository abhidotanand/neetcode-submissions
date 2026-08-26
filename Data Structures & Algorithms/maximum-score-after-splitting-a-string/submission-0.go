func maxScore(s string) int {
    zeros, ones := 0, 0
    res := -1 << 31

    if s[0] == '0' {
        zeros++
    } else {
        ones++
    }

    for i := 1; i < len(s); i++ {
        if zeros-ones > res {
            res = zeros - ones
        }
        if s[i] == '0' {
            zeros++
        } else {
            ones++
        }
    }

    return res + ones
}