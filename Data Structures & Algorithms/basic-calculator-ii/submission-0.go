func calculate(s string) int {
    total, prev, num := 0, 0, 0
    op := byte('+')
    n := len(s)
    i := 0

    for i <= n {
        var ch byte
        if i < n {
            ch = s[i]
        } else {
            ch = '+'
        }
        if ch == ' ' {
            i++
            continue
        }
        if ch >= '0' && ch <= '9' {
            num = num*10 + int(ch-'0')
        } else {
            switch op {
            case '+':
                total += prev
                prev = num
            case '-':
                total += prev
                prev = -num
            case '*':
                prev = prev * num
            default:
                prev = prev / num
            }
            op = ch
            num = 0
        }
        i++
    }
    total += prev
    return total
}