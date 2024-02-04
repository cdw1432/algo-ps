### useful functions
```python
#숫자 진법 변환
def conversion(number, base):
    base_string = "0123456789ABCDEF"
    result = ""
    while number:
        result += base_string[number % base]
        number //= base
    return result[::-1] or "0"
```

### useful modules
```python
from itertools import combinations
from itertools import permutations
```

### notes
```python
#특정 문자열의 인덱스의 문자 바꾸기
str = str[:index] + new_character + str[index+1:]
```