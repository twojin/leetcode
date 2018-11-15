package string

func reverseInt(x int) int {
	result := 0

	for x != 0 {
		if !isInt32(x, x%10) {
			return 0
		}

		result = result*10 + x%10
		x /= 10
	}

	return result
}

/*
C：
int reverse(int x) {
    int result;

    while (x != 0){
       int temp;
        temp = result*10 + x%10;
        if (temp/10 != result){
            return 0;
        }
        result = temp;
        x/=10;
    }

    return result;
}
*/
