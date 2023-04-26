def square(number):
    """

    :param number:
    :return:
    """
    if number > 64 or number < 1:
        raise ValueError("square must be between 1 and 64")
    else:
        return 2 ** (number-1)


def total():
    counter = 0
    for x in range(1, 65):
        counter += square(x)
    return counter
