def equilateral(sides):
    """
    takes array with length of 3 as input and checks if it is a triangle, and furthermore whether it represents an
    equilateral triangle
    :param sides:
    :return:
    """
    return sides[0] == sides[1] and sides[1] == sides[2] and isTriangle(sides)


def isosceles(sides):
    """
    takes array with length of 3 as input and checks if it is a triangle, and furthermore whether it represents an
    isosceles triangle
    :param sides:
    :return:
    """
    return (sides[0] == sides[1] or sides[1] == sides[2] or sides[0] == sides[2]) and isTriangle(sides)


def scalene(sides):
    """
    takes array with length of 3 as input and checks if it is a triangle, and furthermore whether it represents a
    scalene triangle
    :param sides:
    :return:
    """
    return sides[0] != sides[1] and sides[1] != sides[2] and sides[0] != sides[2] and isTriangle(sides)


def isTriangle(sides):
    """
    takes array with length of 3 as input and checks that it represents a triangle
    :param sides:
    :return:
    """
    return sides[0] > 0 and sides[1] > 0 and sides[2] > 0 and sides[0] + sides[1] >= sides[2] and sides[1] + sides[2] >= \
           sides[0] and sides[0] + sides[2] >= sides[1]
