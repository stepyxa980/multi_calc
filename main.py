def minor(mat, i, j):
    c_mat = [row[:] for row in mat]
    c_mat.pop(i-1)
    for i in range(len(c_mat)):
        c_mat[i].pop(j-1)
    return c_mat


def is_square_mat(mat):
    for row in mat:
        if len(row) != len(mat):
            return False
    return True


def det(mat):
    len_mat = len(mat)
    if not is_square_mat(mat):
        return 0
    
    if len_mat == 1:
        return mat[0][0]
    
    if len_mat == 2:
        return mat[0][0] * mat[1][1] - mat[0][1] * mat[1][0]

    determinant = 0

    for j in range(len_mat):
        minor_mat = minor(mat, 1, j+1)

        determinant += ((-1) ** j) * mat[0][j] * det(minor_mat)
    
    return determinant


'''def rang(mat):
    len_mat = len(mat)
    if is_square_mat(mat):
        return 0
    
    rank = 0
    for i in range(len(mat)):

        for row in mat[i]:


        rank += 1'''


def main():
    mat = [[1, 2, 3],
           [4, 8, 6],
           [1, 1, 3]]
    
    mat_minor = minor(mat, 2, 3)

    print("Матрица А: ")
    for row in mat:
        print(row)
    
    print("Минор 2,3 матрицы А: ")
    for row in mat_minor:
        print(row)

    print("Детерминант матрицы А: ")
    print(det(mat))


if __name__ == "__main__":
    main()