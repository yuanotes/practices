import random
import copy


def insertion_sort(array):
    """
    Complexity is O(n*n)
    """
    for j in range(1, len(array)):
        key = array[j]  # the key is import
        i = j - 1
        while i >= 0 and array[i] > key:
            array[i + 1] = array[i]
            i -= 1
        array[i + 1] = key
    return array


if __name__ == "__main__":
    array_1 = range(10)
    array_2 = range(-10, 10)
    array_3 = range(-10, 0)
    for array in [array_1, array_2, array_3]:
        original = copy.copy(array)
        random.shuffle(array)
        assert insertion_sort(array) == original
