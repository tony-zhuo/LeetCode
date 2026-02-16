import pytest
from string_builder import StringBuilder


class TestAppend:
    def test_append(self):
        sb = StringBuilder("hello")
        sb.append(" world")
        assert str(sb) == "hello world"

    def test_append_empty(self):
        sb = StringBuilder()
        sb.append("hi")
        assert str(sb) == "hi"


class TestInsertAt:
    def test_insert_middle(self):
        sb = StringBuilder("hllo")
        sb.insert_at(1, "e")
        assert str(sb) == "hello"

    def test_insert_start(self):
        sb = StringBuilder("world")
        sb.insert_at(0, "hello ")
        assert str(sb) == "hello world"

    def test_insert_out_of_range(self):
        sb = StringBuilder("hi")
        with pytest.raises(IndexError):
            sb.insert_at(-1, "x")
        with pytest.raises(IndexError):
            sb.insert_at(3, "x")


class TestDeleteRange:
    def test_delete(self):
        sb = StringBuilder("hello world")
        sb.delete_range(5, 11)
        assert str(sb) == "hello"

    def test_delete_out_of_range(self):
        sb = StringBuilder("hi")
        with pytest.raises(IndexError):
            sb.delete_range(-1, 1)
        with pytest.raises(IndexError):
            sb.delete_range(0, 3)


class TestCharAt:
    def test_char_at(self):
        sb = StringBuilder("abc")
        assert sb.char_at(0) == "a"
        assert sb.char_at(2) == "c"

    def test_out_of_range(self):
        sb = StringBuilder("abc")
        with pytest.raises(IndexError):
            sb.char_at(3)


class TestSubstring:
    def test_substring(self):
        sb = StringBuilder("hello world")
        assert sb.substring(0, 5) == "hello"
        assert sb.substring(6, 11) == "world"

    def test_out_of_range(self):
        sb = StringBuilder("hi")
        with pytest.raises(IndexError):
            sb.substring(0, 3)


class TestIndexOfAndContains:
    def test_index_of(self):
        sb = StringBuilder("hello world")
        assert sb.index_of("world") == 6
        assert sb.index_of("xyz") == -1

    def test_contains(self):
        sb = StringBuilder("hello world")
        assert sb.contains("hello") is True
        assert sb.contains("xyz") is False


class TestReplace:
    def test_replace_first(self):
        sb = StringBuilder("aabaa")
        assert sb.replace("a", "x") is True
        assert str(sb) == "xabaa"

    def test_replace_not_found(self):
        sb = StringBuilder("abc")
        assert sb.replace("x", "y") is False


class TestReplaceAll:
    def test_replace_all(self):
        sb = StringBuilder("aabaa")
        count = sb.replace_all("a", "x")
        assert str(sb) == "xxbxx"
        assert count == 4

    def test_replace_all_longer(self):
        sb = StringBuilder("abab")
        count = sb.replace_all("ab", "xyz")
        assert str(sb) == "xyzxyz"
        assert count == 2


class TestReverse:
    def test_reverse(self):
        sb = StringBuilder("hello")
        sb.reverse()
        assert str(sb) == "olleh"

    def test_palindrome(self):
        sb = StringBuilder("racecar")
        assert sb.is_palindrome() is True

    def test_not_palindrome(self):
        sb = StringBuilder("hello")
        assert sb.is_palindrome() is False

    def test_empty_palindrome(self):
        assert StringBuilder().is_palindrome() is True


class TestLenAndMisc:
    def test_len(self):
        assert len(StringBuilder("abc")) == 3
        assert len(StringBuilder()) == 0

    def test_count_char(self):
        sb = StringBuilder("aabba")
        assert sb.count_char("a") == 3
        assert sb.count_char("b") == 2
        assert sb.count_char("c") == 0

    def test_to_char_array(self):
        sb = StringBuilder("abc")
        arr = sb.to_char_array()
        assert arr == ["a", "b", "c"]
        arr[0] = "x"  # modifying copy shouldn't affect original
        assert str(sb) == "abc"
