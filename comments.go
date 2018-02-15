package mapast

// LookupComments fills EnderSepar with comment location information from file.
// This information is necessary to recognize comments of various types, like
// comments that span end of line only, or comments that follow an empty lines.
func LookupComments(file []byte, EnderSepar [2]map[int]struct{}) {
	var whitespace bool
	var cleanline bool
	var sawender bool
	for i := 0; i+1 < len(file); i++ {
		var c, d = file[i], file[i+1]
		if c == '\n' {
			if whitespace {
				cleanline = true
			}
			whitespace = true
			sawender = false
			continue
		}
		if c == '/' && d == '/' && !whitespace && !sawender {
			EnderSepar[0][(i+1)/2] = struct{}{}
			sawender = true
		}
		if c == '/' && d == '*' && !whitespace && !sawender {
			EnderSepar[0][(i+1)/2] = struct{}{}
			sawender = true
		}
		if c == '/' && d == '/' && cleanline {
			EnderSepar[1][(i+1)/2] = struct{}{}
			cleanline = false
		}
		if c == '/' && d == '*' && cleanline {
			EnderSepar[1][(i+1)/2] = struct{}{}
			cleanline = false
		}
		if c == 'p' && d == 'a' && cleanline {
			EnderSepar[1][(i+1)/2] = struct{}{}
			cleanline = false
		}
		if c > ' ' {
			whitespace = false
			cleanline = false
		}
	}
}
