func isValid(s string) bool {
    var n int = len(s)
    if n == 1 {
        return false
    }
    var ans []string
	
    for i := 0; i < n; i++ {
        if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
            ans = append(ans, string(s[i]))
        } else if string(s[i]) == ")" {
            if ans[len(ans)-1] == "(" {
                ans = ans[:len(ans)-1]
            } else {
                return false
            }
        } else if string(s[i]) == "}" {
            if ans[len(ans)-1] == "{" {
                ans = ans[:len(ans)-1]
            } else {
                return false
            }
        } else if string(s[i]) == "]" {
            if ans[len(ans)-1] == "[" {
                ans = ans[:len(ans)-1]
            } else {
                return false
            }
        } else [
            return false
        ]
    }
    if len(ans) == 0 {
        return true
    }
    return false
}
