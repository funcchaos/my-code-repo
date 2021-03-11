package problems

//   char[] cs = s.toCharArray();
//         char[] d = new char[s.length()];
//         int hh = 0, tt = -1;
//         for (char c : cs) {
//             if (hh <= tt && d[tt] == c) {
//                 tt--;
//             } else {
//                 d[++tt] = c;
//             }
//         }
//         return new String(d, 0, tt + 1);
//     }

func removeDuplicates(S string) string {
	b := []byte(S)
	d := make([]byte, len(b))
	ti :=  -1
	// aabaab   d[0]=a d[0]=b d[1]=a   
	for _, v := range b {
		if ti>=0 && d[ti] == v {
			ti--
		} else {
			ti++
			d[ti] = v
		}
	}
	return string(d[:ti+1])
}

func removeDuplicates1(S string) string {
    d:=make([]rune,len(S))
	ti :=  -1
	// aabaab   d[0]=a d[0]=b d[1]=a   
	for _, v := range S {
		if ti>=0 && d[ti] == v {
			ti--
		} else {
			ti++
			d[ti] = v
		}
	}
    return string(d[:ti+1])
}