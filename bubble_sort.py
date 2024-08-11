def bubble_sort(l: list):
    for j in range(len(l) - 1):
        for i in range(len(l) - j - 1):
            if l[i] > l[i + 1]:
                l[i + 1], l[i] = l[i], l[i + 1]

if __name__ == '__main__':
    l = list(range(50, 0, -1))
    bubble_sort(l)

    assert l == list(range(1, 51))
