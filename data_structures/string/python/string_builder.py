class StringBuilder:
    def __init__(self, s: str = ""):
        self._data = list(s)

    def append(self, s: str) -> None:
        self._data.extend(s)

    def insert_at(self, index: int, s: str) -> None:
        if index < 0 or index > len(self._data):
            raise IndexError("string index out of range")
        self._data[index:index] = list(s)

    def delete_range(self, start: int, end: int) -> None:
        if start < 0 or end > len(self._data) or start > end:
            raise IndexError("string index out of range")
        del self._data[start:end]

    def char_at(self, index: int) -> str:
        if index < 0 or index >= len(self._data):
            raise IndexError("string index out of range")
        return self._data[index]

    def substring(self, start: int, end: int) -> str:
        if start < 0 or end > len(self._data) or start > end:
            raise IndexError("string index out of range")
        return "".join(self._data[start:end])

    def index_of(self, target: str) -> int:
        s = "".join(self._data)
        idx = s.find(target)
        return idx

    def contains(self, target: str) -> bool:
        return self.index_of(target) >= 0

    def replace(self, old: str, new: str) -> bool:
        idx = self.index_of(old)
        if idx == -1:
            return False
        self._data[idx:idx + len(old)] = list(new)
        return True

    def replace_all(self, old: str, new: str) -> int:
        count = 0
        result: list[str] = []
        s = "".join(self._data)
        i = 0
        while i <= len(s) - len(old):
            if s[i:i + len(old)] == old:
                result.extend(new)
                i += len(old)
                count += 1
            else:
                result.append(s[i])
                i += 1
        result.extend(s[i:])
        self._data = result
        return count

    def reverse(self) -> None:
        self._data.reverse()

    def is_palindrome(self) -> bool:
        i, j = 0, len(self._data) - 1
        while i < j:
            if self._data[i] != self._data[j]:
                return False
            i += 1
            j -= 1
        return True

    def __len__(self) -> int:
        return len(self._data)

    def __str__(self) -> str:
        return "".join(self._data)

    def count_char(self, ch: str) -> int:
        return self._data.count(ch)

    def to_char_array(self) -> list[str]:
        return self._data.copy()
